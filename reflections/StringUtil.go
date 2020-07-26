package reflections

import (
	"fmt"
	"strings"
)

//CreateLikeQueryString generate filter like clause
func CreateLikeQueryString(filter map[string]interface{}) []interface{} {

	var result string
	var likeStrs []string
	var args []string
	for key, value := range filter {
		strItem := "`" + key + "` like ?"
		likeStrs = append(likeStrs, strItem)
		valueAsString := fmt.Sprintf("%v", value)
		args = append(args, "%"+(valueAsString)+"%")
	}
	fmt.Println("likeStrs: ", likeStrs)
	result = strings.Join(likeStrs, " AND ")

	whereClauses := []interface{}{
		result,
	}
	for _, item := range args {
		whereClauses = append(whereClauses, item)
	}

	return whereClauses

}

func isUpperCase(str string) bool {
	return strings.ToUpper(str) == str
}

//ToSnakeCase converts camelCased word to snake_cased
func ToSnakeCase(camelCased string) string {

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