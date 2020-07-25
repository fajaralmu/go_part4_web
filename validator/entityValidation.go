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
			println("NOT Custom Tag")
			continue
		}

		processCustomTag(customTag, field, model)

	}

	return true
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity) {

	println("__________-processCustomTag____________")
	foreignKey := customTag["foreignKey"]
	fieldValue := reflections.GetFieldValue(field.Name, model)

	foreignKeyValue := reflections.GetFieldValue(foreignKey, model)
	entity := fieldValue.(entities.InterfaceEntity)
	fmt.Println("fieldValue: ", entity)
	result := dataaccess.FindByID(entity, foreignKeyValue)
	fmt.Println("result FIND BY ID: ", result)
}
