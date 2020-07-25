package main

import (
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
	updateTest()
}

func updateTest() {
	user := entities.User{
		Model:       gorm.Model{ID: 99},
		Username:    "Fajar_0000",
		DisplayName: "El Fajr Part00000",
		Password:    "12345",
		RoleID:      18,
	}
	repository.Save(&user)

}
