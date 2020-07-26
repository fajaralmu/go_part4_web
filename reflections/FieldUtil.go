package reflections

import (
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/entities"
)

func GetFieldValue(fieldName string, model entities.InterfaceEntity) interface{} {

	r := reflect.ValueOf(model)
	value := reflect.Indirect(r).FieldByName(fieldName)

	// fmt.Println("value: ", value, "value interface: ", value.Interface())

	return value.Interface()
}

func GetMapOfTag(field reflect.StructField, tagName string) (map[string]string, bool) {

	result := map[string]string{}
	value, ok := field.Tag.Lookup(tagName)

	if !ok {
		return result, false
	}

	tagValues := strings.Split(value, ";")
	for _, item := range tagValues {
		keyVal := strings.Split(item, ":")
		result[keyVal[0]] = keyVal[1]
	}
	return result, true
}
