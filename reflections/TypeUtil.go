package reflections

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/entities"
)

//GetStructType Gets Struct Type
func GetStructType(object interface{}) reflect.Type {
	return reflect.TypeOf(object)
}

//GetStructTableName returns snake cased of struct name
func GetStructTableName(object interface{}) string {
	typeName := GetStructType(object)

	return camelCaseToSnakeCase(typeName.Name())
}

func isUpperCase(str string) bool {
	return strings.ToUpper(str) == str
}

func GetFieldValue(fieldName string, model entities.InterfaceEntity) interface{} {

	r := reflect.ValueOf(model)
	return reflect.Indirect(r).FieldByName(fieldName)
}

func GetJoinColumnFields(model entities.InterfaceEntity) []reflect.StructField {

	var result []reflect.StructField

	println("====ValidateJoinColumn====")
	t := reflect.TypeOf(model)
	// r := reflect.ValueOf(model)
	// fmt.Println(" t.Kind(): ", t.Kind(), "r:", r.Kind())
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			fieldValue := GetFieldValue(structField.Name, model)
			isStructType := isStruct(structField)

			if isStructType {
				result = append(result, structField)
			}
			fmt.Println("type:", structField.Type, "name: ", structField.Name, "value:", fieldValue, "isStructType: ", isStructType)
		}
	} else {
		fmt.Println("not a struct")
	}
	return result
}

func isStruct(field reflect.StructField) bool {
	fmt.Println("field.Type: ", field.Type)
	return field.Type.Kind() == reflect.Struct
}

// func convertBytes(b []byte) string {
// 	s := make([]string, len(b))
// 	for i := range b {
// 		s[i] = strconv.Itoa(int(b[i]))
// 	}
// 	return strings.Join(s, ",")
// }

func camelCaseToSnakeCase(camelCased string) string {

	var result string

	for i, char := range camelCased {

		_char := string(char)
		if i > 0 && isUpperCase(_char) {
			result += "_"

		}
		_charStr := strings.ToLower(_char)
		if 0 == i {
			result += strings.ToLower(_char)
		} else {
			result += (_charStr)
		}

	}

	return result
}
