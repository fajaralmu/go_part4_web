package entities

import (
	"fmt"

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
	println("will init DB")
	var err error
	databaseConnection, err = gorm.Open("mysql", "root@(localhost:3306)/base_app_go?charset=utf8&parseTime=True&loc=Local")
	if nil != err {
		fmt.Println("Error Opening DB:", err)
	} else {
		defer databaseConnection.Close()
		println("success init DB")
		operation()

	}
}

//CreateNew adds new db record
func CreateNew(model interface{}) interface{} {

	dbOperation(func() {
		println("will create model")
		res := databaseConnection.AutoMigrate(model)
		fmt.Println("PK is blank :", res)
		databaseConnection.Create(&model)
		println("model created")
		res2 := databaseConnection.NewRecord(&model)
		fmt.Println("PK is blank :", res2)
	})

	return model
}
