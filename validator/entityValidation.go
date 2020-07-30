package validator

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/files"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
)

//ValidateEntity validates entity field before Persisting to DB
func ValidateEntity(model entities.InterfaceEntity, currentRecord entities.InterfaceEntity) bool {
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
		var customTagResult bool
		if currentRecord != nil && reflect.ValueOf(currentRecord).IsValid() && reflect.ValueOf(currentRecord).IsZero() == false {
			currentFieldRecord, _ := reflections.GetFieldValue(field.Name, currentRecord)
			customTagResult = processCustomTag(customTag, field, model, currentFieldRecord)
		} else {
			customTagResult = processCustomTag(customTag, field, model, nil)
		}
		if !customTagResult {
			valid = false
			break loop
		}

	}

	println("_________ END VALIDATION ___________")

	return valid
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity, currentFieldRecord interface{}) bool {

	println("__________-processCustomTag____________ for ", field.Name)

	foreignKey := customTag["foreignKey"]
	foreignKeyOk := processForeignKey(foreignKey, field, model)

	fieldType := customTag["type"]
	fieldOK := processFieldValue(fieldType, field, model, currentFieldRecord)

	println("__________END processCustomTag (", foreignKeyOk, fieldOK, ")____________")

	return foreignKeyOk && fieldOK
}

func processFieldValue(fieldType string, field reflect.StructField, model entities.InterfaceEntity, currentFieldRecord interface{}) bool {
	log.Println("processFieldValue: ", field.Name, "currentFieldRecord: ", currentFieldRecord)
	fieldValue, _ := reflections.GetFieldValue(field.Name, model)

	switch fieldType {
	case "FIELD_TYPE_IMAGE":
		if fieldValue != nil {
			code := strings.Replace(reflect.TypeOf(model).String(), ".", "", -1)
			code = strings.Replace(code, "*", "", -1)
			fieldValue = processImg(fieldValue.(string), code, currentFieldRecord)
			reflections.SetFieldValue(field.Name, fieldValue, model)

		} else {
			log.Println("IMG base64data is Empty")
		}

	}

	return true
}

func processImg(imgData string, code string, currentFieldRecord interface{}) string {
	if (imgData == "") && currentFieldRecord != "" && currentFieldRecord != nil {
		log.Println("imgData is BLANK ... returns currentFieldRecord: ", currentFieldRecord)
		return currentFieldRecord.(string)
	}
	log.Println("Process image dat, code: ", code)
	return files.WriteBase64Img(imgData, code)
}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) bool {
	println("begin process foreign key: ", foreignKey)
	if "" == foreignKey {
		return true
	}
	entity := reflections.StructFieldToEntity(field, model)
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
