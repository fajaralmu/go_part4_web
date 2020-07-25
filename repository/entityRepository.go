package repository

import (
	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/validator"
)

func CreateNew(model entities.InterfaceEntity) {

	ok := validator.ValidateEntity(model)
	if ok {
		println("creating model")
		dataaccess.CreateNew(model)
	} else {
		println("Entity Invalid!")
	}
}

func Save(model entities.InterfaceEntity) {

	ok := validator.ValidateEntity(model)
	if ok {
		println("saving model")
		dataaccess.Save(model)
	} else {
		println("Entity Invalid!")
	}
}
