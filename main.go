package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/repository"
)

func main2() {
	println("____start____")

	dataaccess.InitDatabase()

	// var userRole *entities.UserRole
	// userRole := entities.UserRole{Code: "02", Name: "Regular 2"}
	// entities.CreateNew(&userRole)
	// var userRole2 entities.UserRole
	// dataaccess.FindByID(&userRole2, 17)
	// fmt.Println("userRole2: ", userRole2)
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
	println("USER ROLE ID: ", user.RoleID)

}

func main() {
	println("____start____")
	testFilter()
}

func testFilter() {
	user := []entities.User{}

	repository.Filter(&user, entities.Filter{
		Page:  0,
		Limit: 3,
		FieldsFilter: map[string]interface{}{
			"Username":    "Fajar",
			"DisplayName": "Fajr2",
		},
	}, true)
}

func testFindById() {
	res := repository.FindByID(&entities.User{}, 99)

	fmt.Println("res: ", res)
}

func testUpdate() {
	user := entities.User{
		Model:       gorm.Model{ID: 45},
		Username:    "Fajar_55000",
		DisplayName: "El Fajr Part550000",
		Password:    "12345",
		RoleID:      18,
	}
	repository.CreateNew(&user)

}

func testDelete() {
	user := entities.User{
		Model: gorm.Model{ID: 4225},
	}
	repository.Delete(&user)
}
