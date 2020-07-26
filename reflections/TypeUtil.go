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
func GetJoinColumnFields(_model entities.InterfaceEntity) []reflect.StructField {

	var result []reflect.StructField

	println("====ValidateJoinColumn====")
	model := Dereference(_model)
	entity := model.Interface().(entities.InterfaceEntity)
	t := reflect.TypeOf(entity)
	// r := reflect.ValueOf(model)
	fmt.Println(" t.Kind(): ", t.Kind(), "entity: ", reflect.TypeOf(entity)) //, "r:", r.Kind())
	if t.Kind() == reflect.Struct {
	loop:
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)

			fieldValue := GetFieldValue(structField.Name, entity)

			if isNil(fieldValue) {
				println(structField.Name, "is nil")
				continue loop
			}
			isStructType := isPointerToStruct(structField, entity)

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

func isNil(val interface{}) bool {
	return val == nil || (reflect.ValueOf(val).Kind() == reflect.Ptr && reflect.ValueOf(val).IsNil())
}

//GetIDValue return `ID` field value
func GetIDValue(model entities.InterfaceEntity) interface{} {

	return GetFieldValue("ID", model)

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

//CreateNewType generate new Type of given type
func CreateNewType(t reflect.Type) interface{} {
	fmt.Println("CreateNewType: ", t)

	return reflect.New(t).Interface()
}
