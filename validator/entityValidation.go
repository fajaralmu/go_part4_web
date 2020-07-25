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
	for _, field := range structFields {

		customTag, ok := reflections.GetMapOfTag(field, "custom")

		if !ok {
			println("NO Custom Tag")
			continue
		}

		processCustomTag(customTag, field, model)

	}

	println("_________ END VALIDATION ___________")

	return true
}

func structFieldToEntity(field reflect.StructField, model entities.InterfaceEntity) entities.InterfaceEntity {
	fieldValue := reflections.GetFieldValue(field.Name, model)
	entity := fieldValue.(entities.InterfaceEntity)
	return entity
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity) {

	println("__________-processCustomTag____________")

	foreignKey := customTag["foreignKey"]
	entity := structFieldToEntity(field, model)
	entityID := reflections.GetIDValue(entity)

	setUintValue(foreignKey, entityID.(uint), model)

	result := dataaccess.FindByID(entity, entityID)
	fmt.Println("result FIND BY ID: ", result)

	println("__________END processCustomTag____________")
}

func setUintValue(fieldName string, value uint, model entities.InterfaceEntity) {
	uintVal := uint64(value)
	obj := reflect.Indirect(reflect.ValueOf(model))
	obj.FieldByName(fieldName).SetUint(uintVal)
}
