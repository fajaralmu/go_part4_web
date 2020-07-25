package validator

import (
	"fmt"
	"reflect"

	"github.com/fajaralmu/go_part4_web/dataaccess"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
)

func ValidateEntity(model entities.InterfaceEntity) bool {
	println("***ValidateEntity***")

	structFields := reflections.GetJoinColumnFields(model)
	fmt.Println("structFields size: ", len(structFields))
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
	fieldValue := reflections.GetFieldValue(field.Name, model)
	entity := fieldValue.(entities.InterfaceEntity)
	return entity
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity) bool {

	println("__________-processCustomTag____________")

	foreignKey := customTag["foreignKey"]
	foreignKeyOk := processForeignKey(foreignKey, field, model)

	println("__________END processCustomTag (", foreignKeyOk, ")____________")
	return foreignKeyOk
}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) bool {
	println("begin process foreign key")
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
