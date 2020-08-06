package dataaccess

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var databaseConnection *gorm.DB

// InitDatabase initialize db connection
func InitDatabase() {
	dbOperation(func() {
		println("Test DB")
	})
}

func dbOperation(operation func()) {
	println("_______________________________ will init DB")
	var err error
	databaseConnection, err = gorm.Open("mysql", "root@(localhost:3306)/base_app_go?charset=utf8&parseTime=True&loc=Local")
	if nil != err {
		fmt.Println("Error Opening DB:", err)
	} else {
		defer databaseConnection.Close()
		databaseConnection.SingularTable(true)
		databaseConnection.LogMode(true)
		println("success init DB, Operation BEGINS*****")
		operation()
		println("*****operation ENDS*****")

	}
	println("_______________________________")
}

func autoMigrate(model interface{}) {
	println("will AutoMigrate ", reflections.GetStructType(model).Name())
	databaseConnection.AutoMigrate(model)
	println("AutoMigrated")
}

func addNewRecord(model entities.InterfaceEntity) {
	databaseConnection.NewRecord(model)

	databaseConnection.Create(model)
	println("model created")
}
func updateRecord(model entities.InterfaceEntity) {

	databaseConnection.Save(model)
	println("model saved")
}

//FindByID find model by ID, model must have ID
func FindByID(model interface{}, id interface{}) (entities.InterfaceEntity, bool) {
	fmt.Println("FindByID type: ", reflect.TypeOf(model), "ID: ", id)
	fmt.Println("model: ", model)
	count := 0
	dbOperation(func() {
		databaseConnection.Find(model, id).Count(&count)

	})
	println("count: ", count)
	return model.(entities.InterfaceEntity), count > 0
}

func createOrderClause(orderBy, orderType string) string {
	if orderType == "" {
		orderType = "asc"
	}
	orderClause := reflections.ToSnakeCase(orderBy, true) + " " + orderType
	return orderClause
}

//FilterLike queries by like clause,  results must be a slice
func FilterLike(resultModels interface{}, filter map[string]interface{}, page int, limit int, orderBy string, orderType string) ([]interface{}, int) { //[]interface{}, int{
	resType := extractPointerType(resultModels)
	isSlice := resType.Kind() == reflect.Slice
	log.Println("FilterLike, ", resType, " isSlice: ", isSlice)

	if !isSlice {
		fmt.Errorf("given result %v is not a slice ", resType.Name())
	}

	count := 0
	sliceDeref := reflect.ValueOf(resultModels).Elem().Interface()
	underlyingType := reflections.GetUnderlyingSliceType(sliceDeref)
	reflections.EvaluateFilterMap(filter, underlyingType)
	joinColumns := getJoinQueries(filter, underlyingType)

	dbOperation(func() {

		offset := page * limit
		whereClauses := reflections.CreateLikeQueryString(filter)

		//process count
		dbCount := databaseConnection.Where(whereClauses[0], whereClauses[1:]...)
		if len(joinColumns) > 0 {

			for _, s := range joinColumns {
				dbCount = dbCount.Joins(s)
			}
		}
		dbCount.Find(resultModels).Count(&count)
		//end count
		if count == 0 {
			return
		}
		db := databaseConnection.Offset(offset)
		if limit > 0 {
			db = db.Limit(limit)
		}
		if orderBy != "" {
			orderClause := createOrderClause(orderBy, orderType)
			db = db.Order(orderClause)
		}
		if len(joinColumns) > 0 {
			for _, s := range joinColumns {
				db = db.Joins(s)
			}
		}

		db.Find(resultModels, whereClauses...)
	})
	resultModels = reflections.ToInterfaceSlice(resultModels)
	fmt.Println("Result list size: ", len(resultModels.([]interface{})), "total data: ", count)
	return resultModels.([]interface{}), count

}

//FilterMatch queries by equals clause, results must be a slice
func FilterMatch(resultModels interface{}, filter map[string]interface{}, page int, limit int, orderBy string, orderType string) ([]interface{}, int) { //[]interface{}, int{
	resType := extractPointerType(resultModels)
	isSlice := resType.Kind() == reflect.Slice
	log.Println("FilterMatch, ", resType, " isSlice: ", isSlice)

	if !isSlice {
		fmt.Errorf("Given result %v is not a slice!", resType.Name())
	}

	tableName := reflections.GetStructTableNameFromType(resType)
	count := 0
	sliceDeref := reflect.ValueOf(resultModels).Elem().Interface()
	underlyingType := reflections.GetUnderlyingSliceType(sliceDeref)
	reflections.EvaluateFilterMap(filter, underlyingType)

	dbOperation(func() {

		offset := page * limit

		//process count
		databaseConnection.Where(filter).Table(tableName).Count(&count)
		log.Println("//end count: ", count)
		if count == 0 {
			return
		}
		db := databaseConnection.Offset(offset)
		if limit > 0 {
			db = db.Limit(limit)
		}
		if orderBy != "" {
			orderClause := createOrderClause(orderBy, orderType)
			db = db.Order(orderClause)
		}
		//Finally
		db.Find(resultModels, filter)
	})
	fmt.Println("result: ", resultModels)
	resultModels = reflections.ToInterfaceSlice(resultModels)
	fmt.Println("Result list size: ", len(resultModels.([]interface{})), "total data: ", count)
	return resultModels.([]interface{}), count

}

func extractPointerType(pointer interface{}) reflect.Type {
	return reflect.TypeOf(reflect.ValueOf(pointer).Elem().Interface())
}

func getJoinQueries(filter map[string]interface{}, t reflect.Type) []string {
	result := []string{}
	fields := reflections.GetDeclaredFields(t)
	fieldsMap := reflections.SliceOfFieldToMap(fields)

	for key, field := range fieldsMap {
		newKey := reflections.ToSnakeCase(key, true)
		fieldsMap[newKey] = field
	}
	for key, val := range filter {
		joinItem := ""
		if strings.Contains(key, ".") {
			splitString := strings.Split(key, ".")
			currentField := fieldsMap[splitString[0]]
			foreignKeyName := reflections.GetCustomTagKey(currentField, "foreignKey")

			if foreignKeyName == "" {
				continue
			}

			foreignKeyName = reflections.ToSnakeCase(foreignKeyName, true)

			tableName := reflections.GetStructTableNameFromType(currentField.Type)
			joinItem = "left join " + tableName + " on " + tableName + ".id = " + foreignKeyName

			result = append(result, joinItem)
			filter[tableName+"."+splitString[1]] = val
			delete(filter, key)
		}
	}

	return result
}

////////////////////////////////////

//CreateNew adds new db record
func CreateNew(model entities.InterfaceEntity) interface{} {

	dbOperation(func() {
		autoMigrate(model)
		addNewRecord(model)
		// res2 := databaseConnection.NewRecord(model)
		// fmt.Println("PK is blank :", res2)

	})

	return model
}

//Save update db record
func Save(model entities.InterfaceEntity) interface{} {

	dbOperation(func() {
		autoMigrate(model)
		updateRecord(model)

	})

	return model
}

//Delete removes from record, if has DeletedAt field ti deletes softly
func Delete(model entities.InterfaceEntity, softDelete bool) {

	dbOperation(func() {
		// autoMigrate(model)
		if softDelete {
			log.Println("SOFT DELETE")
			deleteModel(model)
		} else {
			log.Println("/!\\ HARD DELETE /!\\")
			deleteModelPermanently(model)
		}

	})

}

func deleteModel(model entities.InterfaceEntity) {
	databaseConnection.Delete(model)
	println("model deleted")
}

func deleteModelPermanently(model entities.InterfaceEntity) {
	databaseConnection.Unscoped().Delete(model)
	println("model deleted permanently")
}
