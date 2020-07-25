package main

import (
	"fmt"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
)

func main() {
	println("____start____")

	dataaccess.InitDatabase()

	// var userRole *entities.UserRole
	// userRole := entities.UserRole{Code: "02", Name: "Regular 2"}
	// entities.CreateNew(&userRole)
	var userRole2 entities.UserRole
	dataaccess.FindByID(&userRole2, 17)
	fmt.Println("userRole2: ", userRole2)

	user := entities.User{Username: "Fajar2", DisplayName: "El Fajr2", Password: "12345"}
	reflections.GetJoinColumnFields(user)
	dataaccess.CreateNew(&user)

}
