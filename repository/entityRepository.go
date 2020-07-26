package repository

import (
	"fmt"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/fajaralmu/go_part4_web/validator"
)

//CreateNew insert new record to entity, Will REMOVE ID Field
func CreateNew(model entities.InterfaceEntity) {
	validator.RemoveID(model)

	ok := validator.ValidateEntity(model)
	if ok {
		println("Creating Model")
		dataaccess.CreateNew(model)
	} else {
		println("Entity Invalid!")
	}
}

//Save updates entity
func Save(model entities.InterfaceEntity) {

	existInDB := isExistInDB(model)
	fmt.Println("existInDB: ", existInDB)

	if existInDB {
		CreateNew(model)

	} else {

		ok := validator.ValidateEntity(model)
		if ok {
			println("saving model")
			dataaccess.Save(model)
		} else {
			println("Entity Invalid!")
		}
	}
}

func isExistInDB(model entities.InterfaceEntity) bool {
	ID := reflections.GetIDValue(model)
	_, ok := dataaccess.FindByID(model, ID)
	return ok
}
