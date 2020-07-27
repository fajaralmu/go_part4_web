package repository

import (
	"fmt"
	"log"
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

//Filter searches in DB by given parameters
func Filter(models interface{}, filter entities.Filter) ([]interface{}, int) {
	//	models := toSliceOfInterfaceEntity(sliceOfModel)
	fmt.Println("model type: ", reflect.TypeOf(models))

	var list []interface{}
	var validatedList []interface{}
	totalData := 0
	if filter.Exact {
		list, totalData = dataaccess.FilterMatch(models, filter.FieldsFilter, filter.Page, filter.Limit)

	} else {
		list, totalData = dataaccess.FilterLike(models, filter.FieldsFilter, filter.Page, filter.Limit)

	}

	fmt.Println("List size: ", reflect.TypeOf(list), " count result: ", totalData)

	for _, item := range list {
		validatedItem := item.(entities.InterfaceEntity).Validate()
		validated := validateResultObject(validatedItem.(entities.InterfaceEntity))
		validatedList = append(validatedList, validated)
	}

	return validatedList, totalData
}

func validateResultObject(model entities.InterfaceEntity) interface{} {
	println("__validateResultObject__")
	structFields := reflections.GetJoinColumnFields(model, false)
	fmt.Println("structFields size: ", len(structFields))

	for _, field := range structFields {

		customTag, ok := reflections.GetMapOfTag(field, "custom")

		if !ok {
			println("NO Custom Tag")
			continue
		}

		foreignEntity, exist := processForeignKey(customTag["foreignKey"], field, model)
		if exist {
			fmt.Println("reflect.TypeOf(model)", reflect.TypeOf(model))
			modelPtr := reflections.CreateNewType(reflect.TypeOf(model))
			reflections.SetFieldValue(field.Name, foreignEntity, modelPtr)

		}
	}
	return model
}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) (interface{}, bool) {
	log.Println("begin process foreign key: ", foreignKey)

	if "" == foreignKey {
		return nil, false
	}
	fmt.Println("model: ", reflect.TypeOf(model))

	entity := structFieldToEntity(field, model)
	foreignKeyID, zero := reflections.GetFieldValue(foreignKey, model)

	if zero {
		return nil, false
	}

	result, ok := dataaccess.FindByID(entity, foreignKeyID)
	fmt.Println("result FIND BY ID: ", result)
	println("end process foreign key")

	return result, ok
}

func structFieldToEntity(field reflect.StructField, model interface{}) entities.InterfaceEntity {
	fieldValue, _ := reflections.GetFieldValue(field.Name, model)
	entity := fieldValue.(entities.InterfaceEntity)
	return entity
}
