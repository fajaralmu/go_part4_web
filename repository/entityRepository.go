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

	ok := validator.ValidateEntity(model, nil)
	if ok {
		println("Creating Model")
		dataaccess.CreateNew(model)
	} else {
		println("Entity Invalid!")
	}
}

//Delete removes from record
func Delete(model entities.InterfaceEntity) bool {
	_, existInDB := isExistInDB(model)
	fmt.Println("existInDB: ", existInDB)
	if existInDB {
		dataaccess.Delete(model)
	} else {
		println("Record does not exist!")
		return existInDB
	}

	_, stillExist := isExistInDB(model)
	return stillExist == false
}

//Save updates entity
func doSave(model entities.InterfaceEntity, validate bool) {

	result, existInDB := isExistInDB(model)
	fmt.Println("existInDB: ", existInDB)

	if !existInDB {
		CreateNew(model)

	} else {
		ok := true
		if validate {
			ok = validator.ValidateEntity(model, result)
		}

		if ok {
			fmt.Println("saving model: ", model)
			dataaccess.Save(model)
		} else {
			println("Entity Invalid!")
		}
	}
}

func Save(model entities.InterfaceEntity) {

	doSave(model, true)
}
func SaveWihoutValidation(model entities.InterfaceEntity) {

	doSave(model, false)
}

func isExistInDB(model entities.InterfaceEntity) (entities.InterfaceEntity, bool) {
	ID := reflections.GetIDValue(model)
	duplicate := reflections.Dereference(model).Interface()
	duplicatePtr := reflections.CreateNewType(reflect.TypeOf(duplicate))
	obj, ok := dataaccess.FindByID(duplicatePtr, ID)
	return obj, ok
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

//FilterByKey get list of obj by key match given value
func FilterByKey(models interface{}, key string, val interface{}) []interface{} {
	var list []interface{}
	list, _ = dataaccess.FilterMatch(models, map[string]interface{}{
		key: val,
	}, 0, 0, "", "")
	return list
}

//Filter searches in DB by given parameters
func Filter(models interface{}, filter entities.Filter) ([]interface{}, int) {
	//	models := toSliceOfInterfaceEntity(sliceOfModel)
	fmt.Println("model type: ", reflect.TypeOf(models))

	var list []interface{}
	var validatedList []interface{}
	totalData := 0
	if filter.Exacts {
		list, totalData = dataaccess.FilterMatch(models, filter.FieldsFilter, filter.Page, filter.Limit, filter.OrderBy, filter.OrderType)

	} else {
		list, totalData = dataaccess.FilterLike(models, filter.FieldsFilter, filter.Page, filter.Limit, filter.OrderBy, filter.OrderType)

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
	log.Println("START_VALIDATE_RESULT_OBJECT")
	structFields := reflections.GetJoinColumnFields(model, false)
	fmt.Println("JOIN COLUMN FIELD size: ", len(structFields))

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

		} else {
			fmt.Println("processForeignKey of ", field.Name, "returned false")
		}
	}

	log.Println("END_VALIDATE_RESULT_OBJECT")

	return model
}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) (interface{}, bool) {
	log.Println("begin process foreign key: ", foreignKey)

	if "" == foreignKey {
		return nil, false
	}
	fmt.Println("model: ", reflect.TypeOf(model))

	entity := reflections.StructFieldToEntity(field, model)
	foreignKeyID, zero := reflections.GetFieldValue(foreignKey, model)

	if zero {
		return nil, false
	}
	result, ok := dataaccess.FindByID(entity, foreignKeyID)
	fmt.Println("result FIND BY ID: ", result)
	println("end process foreign key")

	return result, ok
}
