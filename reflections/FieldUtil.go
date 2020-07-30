package reflections

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/entities"
)

func StructFieldToEntity(field reflect.StructField, model entities.InterfaceEntity) entities.InterfaceEntity {
	fieldValue, _ := GetFieldValue(field.Name, model)
	fmt.Println("structFieldToEntity fieldValue: ", fieldValue)
	entity := fieldValue.(entities.InterfaceEntity)
	return entity
}

func GetFieldValue(fieldName string, model interface{}) (interface{}, bool) {
	log.Printf("GetFieldValue [%v] FROM MODEL: %v \n", fieldName, reflect.TypeOf(model))
	//fmt.Println("MODEL : ", model)
	r := reflect.ValueOf(model)
	value := reflect.Indirect(r).FieldByName(fieldName)
	isZero := !value.IsValid() || value.IsZero()

	fmt.Println("value interface: ", value.Interface(), " isZero: ", isZero)

	return value.Interface(), isZero
}

func SetFieldValue(fieldName string, fieldValue interface{}, model interface{}) {
	fmt.Println("SET", fieldName, "value: ", fieldValue, reflect.TypeOf(model))

	r := reflect.ValueOf(model)
	value := reflect.Indirect(r).FieldByName(fieldName)
	//val := reflect.ValueOf(fieldValue)
	value.Set(reflect.ValueOf(fieldValue))
}

func GetMapOfTag(field reflect.StructField, tagName string) (map[string]string, bool) {

	result := map[string]string{}
	value, ok := field.Tag.Lookup(tagName)

	log.Printf("Lookup field %v tagName %v, ok: %v \n", field.Name, tagName, ok)

	if !ok {
		return result, false
	}

	tagValues := strings.Split(value, ";")
	for _, item := range tagValues {
		keyVal := strings.Split(item, ":")
		result[strings.Trim(keyVal[0], " ")] = strings.Trim(keyVal[1], " ")
	}
	return result, true
}
