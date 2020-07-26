package repository

import (
	"fmt"
	"reflect"

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

//Delete removes from record
func Delete(model entities.InterfaceEntity) {
	existInDB := isExistInDB(model)
	fmt.Println("existInDB: ", existInDB)
	if existInDB {
		dataaccess.Delete(model)
	} else {
		println("Record does not exist!")
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

//FindByID return model from DB with given ID
func FindByID(model entities.InterfaceEntity, ID uint) entities.InterfaceEntity {
	validator.RemoveID(model)
	validator.SetID(model, ID)
	result, ok := dataaccess.FindByID(model, ID)
	if ok {
		return result
	}
	return nil
}

func Filter(models interface{}, filter entities.Filter) {
	//	models := toSliceOfInterfaceEntity(sliceOfModel)
	list, count := dataaccess.FilterLike(models, filter.FieldsFilter, filter.Page, filter.Limit)
	fmt.Println("List size: ", reflect.TypeOf(list), " count result: ", count)
}
