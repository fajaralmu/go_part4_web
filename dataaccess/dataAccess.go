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
	// tableName := model.TableName()
	// // modelMap := make(map[string]interface{}) //, 5)
	// // modelMap = map[string]interface{}{

	// // 	"Code":   "dddd",
	// // 	"Access": "33333",
	// // 	"Name":   "TEST",
	// // }
	// println("tableName: ", tableName)

	databaseConnection.Create(model)
	println("model created")
}

//FindByID find model by ID
func FindByID(model entities.InterfaceEntity, id interface{}) entities.InterfaceEntity {
	fmt.Println("FindByID type: ", reflect.TypeOf(model), "ID: ", id)
	dbOperation(func() {
		databaseConnection.Find(model, id)
	})
	return model
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
