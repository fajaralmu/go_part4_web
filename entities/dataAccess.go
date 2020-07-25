package entities

import (
	"fmt"

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
		println("operation BDGINS")
		operation()
		println("operation ENDS")

	}
	println("_______________________________")
}

func (u BaseEntity) getTableName() string {

	res := reflections.GetStructTableName(u)
	println("result: ", res)
	return res
}

func autoMigrate(model interface{}) {
	println("will AutoMigrate ", reflections.GetStructType(model).Name())
	databaseConnection.AutoMigrate(model)
	println("AutoMigrated")
}

func addNewRecord(model InterfaceEntity) {
	databaseConnection.NewRecord(model)
	tableName := reflections.GetStructTableName(model)
	// modelMap := make(map[string]interface{}) //, 5)
	// modelMap = map[string]interface{}{

	// 	"code":   "dddd",
	// 	"access": "33333",
	// 	"name":   "TEST",
	// }

	databaseConnection.Table(tableName).Create(&model)
	println("model created")
}

//CreateNew adds new db record
func CreateNew(model InterfaceEntity) interface{} {

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
