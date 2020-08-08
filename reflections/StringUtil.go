package reflections

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

		strItem := "`"
		//check join columns
		if strings.Contains(key, ".") {
			stringRaw := strings.Split(key, ".")
			strItem += stringRaw[0] + "`.`" + stringRaw[1]
		} else {
			strItem += ToSnakeCase(key, true)
		}

		if strings.HasSuffix(strItem, "[exacts]") {
			strItem = strings.Replace(strItem, "[exacts]", "", 1)
			strItem += "` = ?"
			likeStrs = append(likeStrs, strItem)
			valueAsString := fmt.Sprintf("%v", value)
			args = append(args, (valueAsString))
		} else {
			strItem += "` like ?"
			likeStrs = append(likeStrs, strItem)
			valueAsString := fmt.Sprintf("%v", value)
			valueAsString = "%" + (valueAsString) + "%"
			args = append(args, (valueAsString))
		}

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

func adjustCamelCase(raw string) string {
	res := ""
	currentChar := ""
	for i, char := range raw {
		_char := string(char)
		toLower := false
		if i > 0 && isUpperCase(_char) {
			if isUpperCase(currentChar) {
				toLower = true
			}
		}
		currentChar = _char
		if toLower {
			res += strings.ToLower(_char)
		} else {
			res += _char
		}

	}
	return res
}

//ToSnakeCase converts camelCased word to snake_cased (ALL LOWERCASE)
func ToSnakeCase(camelCased string, lowerCaseResult bool) string {

	var result string
	var currentUpperCase bool = false
	camelCased = adjustCamelCase(camelCased)
	currentChar := ""
loop:
	for i, char := range camelCased {
		_char := string(char)
		if _char == "_" {
			result += _char
			currentChar = _char
			continue loop
		}

		if i > 0 && isUpperCase(_char) && currentUpperCase == false {
			currentUpperCase = true
			if i > 1 && currentChar != "_" {
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
		currentChar = _char

	}

	if lowerCaseResult {
		result = strings.ToLower(result)
	}

	log.Println("SNAKE CASED from: ", camelCased, "-->", result, ";lowerCaseResult(", lowerCaseResult, ")")

	return result

}

//GetFileExtension returns file extension
func GetFileExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	ext = strings.Replace(ext, ".", "", 1)
	return ext
}

//RemoveCharsAfter excludes chars in the string str after given _char
func RemoveCharsAfter(str string, _char string) string {
	res := ""

	for _, char := range str {
		if string(char) == _char {
			break
		}
		res += string(char)
	}
	return res
}

//GetWordsAfterLastChar removes all character before specified character
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

//LimitStr removes indexes of string after the limit
func LimitStr(raw interface{}, limit int) string {
	var fieldValueStr string = ""
	str := fmt.Sprintf("%v", raw)
	if len(str) > limit {
		fieldValueStr = str[:(limit - 1)]
	}
	return fieldValueStr
}

//ToJSONString create JSON string
func ToJSONString(i interface{}) string {
	jsonStr, _ := json.Marshal(i)
	return string(jsonStr)
}

//IsNumericValue checks if the given string is numeric
func IsNumericValue(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//RandomCounter is random generation counter
var RandomCounter int = time.Now().Second()

//RandomNum generates random Int string with specified length
func RandomNum(length int) string {
	res := ""
	numLength := len(arrayOfNums)
	for i := 0; i < length; i++ {
		s1 := rand.NewSource(int64(i * RandomCounter))
		r1 := rand.New(s1)
		res += arrayOfNums[r1.Intn(numLength)]

		RandomCounter++
	}
	RandomCounter += RandomCounter * 2
	return res
}

//ExtractCamelCase returns word with space delimiter
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
