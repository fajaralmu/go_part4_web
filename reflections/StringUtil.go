package reflections

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

//ToSnakeCase converts camelCased word to snake_cased (ALL LOWERCASE)
func ToSnakeCase(camelCased string) string {

	var result string
	var currentUpperCase bool = false

	for i, char := range camelCased {

		_char := string(char)
		if i > 0 && isUpperCase(_char) && currentUpperCase == false {
			currentUpperCase = true
			if i > 1 {
				result += "_"
			}

		} else {
			currentUpperCase = false
		}
		_charStr := strings.ToLower(_char)
		if 0 == i {
			result += strings.ToLower(_char)
		} else {
			result += (_charStr)
		}
	}

	return strings.ToLower(result)
}

func ToJSONString(i interface{}) string {
	jsonStr, _ := json.Marshal(i)
	return string(jsonStr)
}

func IsNumericValue(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//RandomNum generates random Int string with specified length
func RandomNum(length int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	res := strconv.Itoa(r1.Intn(length))
	return res
}

func extractCamelCase(camelCased string) string {

	var result string = ""

	for i, char := range camelCased {
		_char := string(char)
		if isUpperCase(_char) {
			result += " "
		}
		if 0 == i {
			result += strings.ToUpper(_char)

		} else {
			result += (_char)
		}

	}

	return result
}
