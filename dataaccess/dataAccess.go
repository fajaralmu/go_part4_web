package dataaccess

import (
	"fmt"
	"log"
	"reflect"

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

//FilterLike queries by like clause,  results must be a slice
func FilterLike(result interface{}, filter map[string]interface{}, page int, limit int, orderBy string, orderType string) ([]interface{}, int) { //[]interface{}, int{
	resType := extractPointerType(result)
	isSlice := resType.Kind() == reflect.Slice
	log.Println("FilterLike, ", resType, " isSlice: ", isSlice)

	if !isSlice {
		fmt.Errorf("given result %v is not a slice ", resType.Name())
	}

	count := 0
	sliceDeref := reflect.ValueOf(result).Elem().Interface()
	underlyingType := reflections.GetUnderlyingSliceType(sliceDeref)
	reflections.EvaluateFilterMap(filter, underlyingType)
	offset := page * limit

	dbOperation(func() {

		whereClauses := reflections.CreateLikeQueryString(filter)

		//process count
		databaseConnection.Where(whereClauses[0], whereClauses[1:]...).Find(result).Count(&count)
		//end count
		if count == 0 {
			return
		}

		if limit > 0 {
			if orderBy != "" {
				if orderType == "" {
					orderType = "asc"
				}
				orderClause := reflections.ToSnakeCase(orderBy, true) + " " + orderType
				databaseConnection.Offset(offset).Limit(limit).Order(orderClause).Find(result, whereClauses...)
			} else {
				databaseConnection.Offset(offset).Limit(limit).Find(result, whereClauses...)
			}

		} else {
			if orderBy != "" {
				if orderType == "" {
					orderType = "asc"
				}
				orderClause := reflections.ToSnakeCase(orderBy, true) + " " + orderType
				databaseConnection.Offset(offset).Order(orderClause).Find(result, whereClauses...)
			} else {
				databaseConnection.Offset(offset).Find(result, whereClauses...)
			}

		}

	})
	result = reflections.ToInterfaceSlice(result)
	fmt.Println("Result list size: ", len(result.([]interface{})), "total data: ", count)
	return result.([]interface{}), count

}

func extractPointerType(pointer interface{}) reflect.Type {
	return reflect.TypeOf(reflect.ValueOf(pointer).Elem().Interface())
}

//FilterMatch queries by equals clause, results must be a slice
func FilterMatch(result interface{}, filter map[string]interface{}, page int, limit int, orderBy string, orderType string) ([]interface{}, int) { //[]interface{}, int{
	resType := extractPointerType(result)
	isSlice := resType.Kind() == reflect.Slice
	log.Println("FilterMatch, ", resType, " isSlice: ", isSlice)
	if !isSlice {
		fmt.Errorf("\n Given result %v is not a slice! \n", resType.Name())
	}
	tableName := reflections.GetStructTableNameFromType(resType)
	count := 0
	sliceDeref := reflect.ValueOf(result).Elem().Interface()
	underlyingType := reflections.GetUnderlyingSliceType(sliceDeref)
	reflections.EvaluateFilterMap(filter, underlyingType)
	offset := page * limit

	dbOperation(func() {
		//process count
		// res := &result
		databaseConnection.Where(filter).Table(tableName).Count(&count)
		log.Println("//end count: ", count)
		if count == 0 {
			return
		}

		if limit > 0 {
			if orderBy != "" {
				if orderType == "" {
					orderType = "asc"
				}
				orderClause := reflections.ToSnakeCase(orderBy, true) + " " + orderType
				databaseConnection.Offset(offset).Limit(limit).Order(orderClause).Find(result, filter)

			} else {
				databaseConnection.Offset(offset).Limit(limit).Find(result, filter)
			}

		} else {
			if orderBy != "" {
				if orderType == "" {
					orderType = "asc"
				}
				orderClause := reflections.ToSnakeCase(orderBy, true) + " " + orderType
				databaseConnection.Offset(offset).Order(orderClause).Find(result, filter)
			} else {
				databaseConnection.Offset(offset).Find(result, filter)
			}
		}

	})
	fmt.Println("result: ", result)
	result = reflections.ToInterfaceSlice(result)
	fmt.Println("Result list size: ", len(result.([]interface{})), "total data: ", count)
	return result.([]interface{}), count

}

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
