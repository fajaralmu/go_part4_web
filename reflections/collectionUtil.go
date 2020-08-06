package reflections

import "reflect"

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
