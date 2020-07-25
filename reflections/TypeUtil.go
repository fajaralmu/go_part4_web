package reflections

import (
	"reflect"
	"strings"
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
