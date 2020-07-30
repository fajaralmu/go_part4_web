package reflections

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

var arrayOfNums []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "0"}
var arrayOfStrings []string = []string{"a", "b", "c", "e", "d", "f", "g", "h", "i", "j", "k",
	"l", "m", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

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
func ToSnakeCase(camelCased string, lowerCaseResult bool) string {

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
		if 0 == i && lowerCaseResult {
			result += strings.ToLower(_char)
		} else {
			if 0 == i {
				result += _char
			} else {
				result += (_charStr)
			}

		}
	}

	if lowerCaseResult {
		result = strings.ToLower(result)
	}

	log.Println("SNAKE CASER camelCased: ", camelCased, "-->", result, ";lowerCaseResult(", lowerCaseResult, ")")

	return result

}

func GetWordsAfterLastChar(str string, lastChar string) string {
	res := str
	lastCharIdx := 0
	for i, char := range str {
		if string(char) == lastChar {
			lastCharIdx = i
		}
	}
	if lastCharIdx > 0 && len(str) > lastCharIdx+1 {
		res = str[(lastCharIdx + 1):]
	}

	return res
}

func ToJSONString(i interface{}) string {
	jsonStr, _ := json.Marshal(i)
	return string(jsonStr)
}

func IsNumericValue(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

var randomCounter int = 0

//RandomNum generates random Int string with specified length
func RandomNum(length int) string {
	res := ""
	numLength := len(arrayOfNums)
	for i := 0; i < length; i++ {
		s1 := rand.NewSource(int64(i * randomCounter))
		r1 := rand.New(s1)
		res += arrayOfNums[r1.Intn(numLength)]

		randomCounter++
	}
	randomCounter += randomCounter * 2
	return res
}

func ExtractCamelCase(camelCased string) string {

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
