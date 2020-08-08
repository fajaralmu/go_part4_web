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
	doCreate(model, true)
}

//CreateNewWithoutValidation insert new record to entity (NO VALIDTION), Will REMOVE ID Field
func CreateNewWithoutValidation(model entities.InterfaceEntity) {
	doCreate(model, false)
}

func doCreate(model entities.InterfaceEntity, withValidation bool) {
	validator.RemoveID(model)
	ok := true
	if withValidation {
		ok = validator.ValidateEntity(model, nil)
	}

	if ok {
		fmt.Println("Creating Model")
		dataaccess.CreateNew(model)
	} else {
		fmt.Println("Entity Invalid!")
	}
}

//Delete removes from record
func Delete(model entities.InterfaceEntity, softDelete bool) bool {
	_, existInDB := isExistInDB(model)
	fmt.Println("existInDB: ", existInDB)
	if existInDB {
		dataaccess.Delete(model, softDelete)
	} else {
		fmt.Println("Record does not exist!")
		return existInDB
	}

	_, stillExist := isExistInDB(model)
	return stillExist == false
}

func doSave(model entities.InterfaceEntity, validate bool) interface{} {

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
			result := dataaccess.Save(model)
			return result
		} else {
			fmt.Println("Entity Invalid!")
		}
	}
	return model
}

//SaveAndValidate validates entity value and save record
func SaveAndValidate(model entities.InterfaceEntity) interface{} {

	return doSave(model, true)
}

//SaveWihoutValidation save record without validating
func SaveWihoutValidation(model entities.InterfaceEntity) interface{} {

	return doSave(model, false)
}

func isExistInDB(model entities.InterfaceEntity) (entities.InterfaceEntity, bool) {
	ID := reflections.GetIDValue(model)
	modelDereferenced := reflections.Dereference(model).Interface()
	duplicatedPtr := reflections.CreateNewType(reflect.TypeOf(modelDereferenced))
	obj, ok := dataaccess.FindByID(duplicatedPtr, ID)
	return obj, ok
}

//FindByID return model from DB with given ID
func FindByID(model entities.InterfaceEntity, ID uint) entities.InterfaceEntity {
	validator.RemoveID(model)
	validator.SetModelID(model, ID)
	result, ok := dataaccess.FindByID(model, ID)
	if ok {
		return result
	}
	return nil
}

//FilterByKey get list of obj by key match given value, models must be a slice
func FilterByKey(models interface{}, key string, val interface{}) []interface{} {
	log.Println("FilterByKey: ", key, " val: ", val)

	var list []interface{}
	list, count := dataaccess.FilterMatch(models, map[string]interface{}{
		key: val,
	}, 0, 0, "", "")

	log.Println("res size: ", len(list), count)
	return list
}

//Filter searches in DB by given parameters, models must be a slice
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

	structFields := reflections.GetJoinColumnFields(model, false)
	fmt.Println("START_VALIDATE_RESULT_OBJECT\nJOIN COLUMN FIELD size: ", len(structFields))

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
	fmt.Println("result FIND BY ID: ", result, "\n end process foreign key")

	return result, ok
}
