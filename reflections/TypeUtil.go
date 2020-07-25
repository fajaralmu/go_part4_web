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
	value := reflect.Indirect(r).FieldByName(fieldName)

	// fmt.Println("value: ", value, "value interface: ", value.Interface())

	return value.Interface()
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
			isStructType := isPointerToStruct(structField, model)

			if isStructType {
				result = append(result, structField)
			}
			fmt.Println("type:", structField.Type.Kind(), structField.Type.PkgPath(), "name: ", structField.Name, "value:", fieldValue, "\nisPointerToStruct: ", isStructType)
			fmt.Println("__________________")
		}
	} else {
		fmt.Println("not a struct")
	}
	return result
}

func isStruct(field reflect.StructField) bool {
	return strings.Contains(field.Type.PkgPath(), "entities") && field.Type.Kind() == reflect.Struct
}

func isPointerToStruct(field reflect.StructField, model entities.InterfaceEntity) bool {

	fieldValue := GetFieldValue(field.Name, model)
	fieldVal := reflect.ValueOf(fieldValue)
	var fieldValDereference reflect.Value

	if fieldVal.Kind() == reflect.Ptr {
		// fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
		fieldValDereference = fieldVal.Elem()
		// fmt.Println("fieldValDeref: ", reflect.TypeOf(fieldValDereference.Interface()))
		return reflect.TypeOf(fieldValDereference.Interface()).Kind() == reflect.Struct
	} else {
		fieldValDereference = fieldVal
	}

	return false
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
