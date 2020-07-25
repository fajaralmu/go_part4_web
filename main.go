package main

import (
	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/validator"
)

func main() {
	println("____start____")

	dataaccess.InitDatabase()

	// var userRole *entities.UserRole
	// userRole := entities.UserRole{Code: "02", Name: "Regular 2"}
	// entities.CreateNew(&userRole)
	// var userRole2 entities.UserRole
	// dataaccess.FindByID(&userRole2, 17)
	// fmt.Println("userRole2: ", userRole2)
	userRole := &entities.UserRole{}
	userRole.ID = 18
	user := entities.User{

		Username:    "Fajar_5",
		DisplayName: "El Fajr Part5",
		Password:    "12345",
		// RoleID:      18,
		Role: userRole,
	}

	validator.ValidateEntity(&user)
	println("USER ROLE ID: ", user.RoleID)
	dataaccess.CreateNew(&user)

}
