package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/fajaralmu/go_part4_web/app"
	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/repository"
)

func main2() {
	println("____start____")

	dataaccess.InitDatabase()

	userRole := &entities.UserRole{}
	userRole.ID = 1811
	user := entities.User{

		Username:    "Fajar_5",
		DisplayName: "El Fajr Part5",
		Password:    "12345",
		// RoleID:      18,
		Role: userRole,
	}

	repository.CreateNew(&user)
	println("USER ROLE ID: ", user.UserRoleID)

}

func main() {
	log.Println("____start____")
	app.Init()
	app.Run()
	// testFilter()
}

func testFilter() {
	var user *[]entities.User = &[]entities.User{}

	resulstList, count := repository.Filter(user, entities.Filter{
		Page:  0,
		Limit: 3,
		FieldsFilter: map[string]interface{}{
			"Username":    "Fajar",
			"DisplayName": "Fajr2",
		},
	})

	fmt.Println("list size: ", len(resulstList), "count: ", count)
}

func testFindByID() {
	res := repository.FindByID(&entities.User{}, 99)

	fmt.Println("res: ", res)
}

func testUpdate() {
	user := entities.User{
		Model:       gorm.Model{ID: 45},
		Username:    "Fajar_55000",
		DisplayName: "El Fajr Part550000",
		Password:    "12345",
		UserRoleID:  18,
	}
	repository.CreateNew(&user)

}

func testDelete() {
	user := entities.User{
		Model: gorm.Model{ID: 4225},
	}
	repository.Delete(&user, true)
}
