package validator

import (
	"fmt"
	"log"
	"reflect"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/files"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
)

//ValidateEntity validates entity field before persisting to DB
func ValidateEntity(model entities.InterfaceEntity) bool {
	println("***ValidateEntity***", reflect.TypeOf(model))

	structFields := reflections.GetDeclaredFields(reflect.TypeOf(reflections.Dereference(model).Interface()))
	fmt.Println("GetDeclaredFields size: ", len(structFields))
	valid := true
loop:
	for _, field := range structFields {

		customTag, ok := reflections.GetMapOfTag(field, "custom")

		if !ok {
			println("NO Custom Tag")
			continue
		}

		customTagResult := processCustomTag(customTag, field, model)
		if !customTagResult {
			valid = false
			break loop
		}

	}

	println("_________ END VALIDATION ___________")

	return valid
}

func structFieldToEntity(field reflect.StructField, model entities.InterfaceEntity) entities.InterfaceEntity {
	fieldValue, _ := reflections.GetFieldValue(field.Name, model)
	entity := fieldValue.(entities.InterfaceEntity)
	return entity
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity) bool {

	println("__________-processCustomTag____________ for ", field.Name)

	foreignKey := customTag["foreignKey"]
	foreignKeyOk := processForeignKey(foreignKey, field, model)

	fieldType := customTag["type"]
	fieldOK := processFieldValue(fieldType, field, model)

	println("__________END processCustomTag (", foreignKeyOk, fieldOK, ")____________")

	return foreignKeyOk && fieldOK
}

func processFieldValue(fieldType string, field reflect.StructField, model entities.InterfaceEntity) bool {
	log.Println("processFieldValue: ", field.Name)
	fieldValue, _ := reflections.GetFieldValue(field.Name, model)

	// if !ok {
	// 	log.Println("Error getting field [", field.Name, "] Value")
	// 	return false
	// }

	switch fieldType {
	case "FIELD_TYPE_IMAGE":
		if fieldValue != nil {
			fieldValue = processImg(fieldValue.(string), reflect.TypeOf(model).Name())
			reflections.SetFieldValue(field.Name, fieldValue, model)
		} else {
			log.Println("IMG VAL is NIL")
		}

	}

	return true
}

func processImg(imgData string, code string) string {
	log.Println("Process image dat, code: ", code)
	return files.WriteBase64Img(imgData, code)
}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) bool {
	println("begin process foreign key: ", foreignKey)
	if "" == foreignKey {
		return true
	}
	entity := structFieldToEntity(field, model)
	entityID := reflections.GetIDValue(entity)

	setUintValue(foreignKey, entityID.(uint), model)

	result, ok := dataaccess.FindByID(entity, entityID)
	fmt.Println("result FIND BY ID: ", result)
	println("end process foreign key")

	return ok
}

func setUintValue(fieldName string, value uint, model entities.InterfaceEntity) {
	uintVal := uint64(value)
	obj := reflect.Indirect(reflect.ValueOf(model))
	obj.FieldByName(fieldName).SetUint(uintVal)
}

//RemoveID removes `ID` value
func RemoveID(model entities.InterfaceEntity) {
	setUintValue("ID", 0, model)
}

//SetID set ID field of the model
func SetID(model entities.InterfaceEntity, id uint) {
	setUintValue("ID", id, model)
}
