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

	return ToSnakeCase(typeName.Name())
}

//ToInterfaceSlice converts war slice to slice of interface
func ToInterfaceSlice(object interface{}) []interface{} {

	rawSlice := Dereference(object).Interface()
	result := []interface{}{}
	s := reflect.ValueOf(rawSlice)
	// rawSlice = Dereference()
	switch reflect.TypeOf(rawSlice).Kind() {
	case reflect.Slice:

		for i := 0; i < s.Len(); i++ {
			item := s.Index(i).Interface()
			result = append(result, item)
		}
	}
	return result
}

//GetJoinColumnFields return fields having tag "custom" and tagKey: "foreign key"
func GetJoinColumnFields(_model entities.InterfaceEntity, skipNull bool) []reflect.StructField {

	var result []reflect.StructField

	println("====ValidateJoinColumn====")
	model := Dereference(_model)
	entity := model.Interface().(entities.InterfaceEntity)
	t := reflect.TypeOf(entity)

	if t.Kind() == reflect.Struct {
	loop:
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)

			fieldValue, _ := GetFieldValue(structField.Name, entity)

			if skipNull && isNil(fieldValue) {
				println(structField.Name, "is nil, will continue")
				continue loop
			}
			isStructType := isJoinColumn(structField, entity)

			if isStructType {
				result = append(result, structField)
			}
			fmt.Println("type:", structField.Type.Kind(), structField.Type.PkgPath(), "name: ", structField.Name, "value:", fieldValue, "is join column: ", isStructType)
			fmt.Println("__________________")
		}
	} else {
		fmt.Println("not a struct")
	}
	return result
}

func isNil(val interface{}) bool {
	return val == nil || (reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil())
}

//GetIDValue return `ID` field value
func GetIDValue(model entities.InterfaceEntity) interface{} {

	res, _ := GetFieldValue("ID", model)
	return res

}

func isStruct(field reflect.StructField) bool {
	return strings.Contains(field.Type.PkgPath(), "entities") && field.Type.Kind() == reflect.Struct
}

//Dereference from ptr to pointedTo
func Dereference(model interface{}) reflect.Value {
	fieldVal := reflect.ValueOf(model)

	if fieldVal.Kind() == reflect.Ptr {
		// fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
		return fieldVal.Elem()
	}
	return fieldVal
}

func isJoinColumn(field reflect.StructField, model entities.InterfaceEntity) bool {
	customTag, ok := GetMapOfTag(field, "custom")

	if !ok {
		println("NO Custom Tag")
		return false
	}

	return customTag["foreignKey"] != ""
}

//CreateNewTypeNotPointer generate new type
func CreateNewTypeNotPointer(t reflect.Type) interface{} {
	return reflect.Indirect(reflect.New(t)).Interface()
}

//CreateNewType generate new Type pointer pointing TO given type
func CreateNewType(t reflect.Type) interface{} {
	fmt.Println("Initialize new pointer pointing to type : ", t)

	return reflect.New(t).Interface()
}
