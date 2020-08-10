package reflections

import (
	"log"
	"reflect"
)

//SliceOfFieldToMap convert slice of structFields to map[fieldName][fields]
func SliceOfFieldToMap(fields []reflect.StructField) map[string]reflect.StructField {
	result := map[string]reflect.StructField{}
	// switch reflect.TypeOf(slice).Kind() {
	// case reflect.Slice:
	// 	s := reflect.ValueOf(slice)

	// 	for i := 0; i < s.Len(); i++ {
	// 		fmt.Println(s.Index(i))
	// 	}
	// }
	for _, field := range fields {
		result[field.Name] = field
	}
	return result
}

//ConvertInterfaceToSlice converts interface to []interface
func ConvertInterfaceToSlice(rawSlice interface{}) []interface{} {
	result := []interface{}{}
	switch reflect.TypeOf(rawSlice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(rawSlice)

		for i := 0; i < s.Len(); i++ {
			result = append(result, s.Index(i).Interface())
		}
	}
	log.Println("result size: ", len(result))
	return result
}
