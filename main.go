package main

import (
	"github.com/fajaralmu/go_part4_web/entities"
)

func main() {
	println("____start____")

	entities.InitDatabase()

	userRole := entities.UserRole{Code: "01", Name: "Regular"}
	entities.CreateNew(userRole)
}
