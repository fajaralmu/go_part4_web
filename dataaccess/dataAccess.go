package dataaccess

import (
	"fmt"
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
	println("_______________________________")
	println("will init DB")
	var err error
	databaseConnection, err = gorm.Open("mysql", "root@(localhost:3306)/base_app_go?charset=utf8&parseTime=True&loc=Local")
	if nil != err {
		fmt.Println("Error Opening DB:", err)
	} else {
		defer databaseConnection.Close()
		databaseConnection.SingularTable(true)
		databaseConnection.LogMode(true)
		println("success init DB")

		println("*****operation BEGINS*****")
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
	count := 0
	dbOperation(func() {
		databaseConnection.Find(model, id).Count(&count)

	})
	println("count: ", count)
	return model.(entities.InterfaceEntity), count > 0
}

//FilterLike queries by like clause
func FilterLike(result interface{}, filter map[string]interface{}, page int, limit int) ([]interface{}, int) { //[]interface{}, int{
	count := 0
	reflections.EvaluateFilterMap(filter)
	offset := page * limit

	dbOperation(func() {

		whereClauses := reflections.CreateLikeQueryString(filter)
		if limit > 0 {
			databaseConnection.Offset(offset).Limit(limit).Find(result, whereClauses...)
		} else {
			databaseConnection.Offset(offset).Find(result, whereClauses...)
		}

		databaseConnection.Where(whereClauses[0], whereClauses[1:]...).Find(result).Count(&count)

	})
	result = reflections.ToInterfaceSlice(result)
	fmt.Println("Result list size: ", len(result.([]interface{})), "total data: ", count)
	return result.([]interface{}), count

}

//FilterMatch queries by equals clause
func FilterMatch(result interface{}, filter map[string]interface{}, page int, limit int) ([]interface{}, int) { //[]interface{}, int{
	count := 0
	reflections.EvaluateFilterMap(filter)
	offset := page * limit

	dbOperation(func() {

		if limit > 0 {
			databaseConnection.Offset(offset).Limit(limit).Find(result, filter)
		} else {
			databaseConnection.Offset(offset).Find(result, filter)
		}

		databaseConnection.Where(filter).Find(result).Count(&count)

	})
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
func Delete(model entities.InterfaceEntity) {

	dbOperation(func() {
		// autoMigrate(model)
		deleteModel(model)
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

/**
	//long and bored code
t := reflect.TypeOf(*&model)
if t.Kind() == reflect.Struct {
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		fmt.Println(structField.Type, structField.Name)
	}
} else {
	fmt.Println("not a stuct")
}

*/
