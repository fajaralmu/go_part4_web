package repository

import (
	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/validator"
)

func CreateNew(model entities.InterfaceEntity) {

	ok := validator.ValidateEntity(model)
	if ok {
		dataaccess.CreateNew(model)
	} else {
		println("Entity Invalid!")
	}
}
