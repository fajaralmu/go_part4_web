package reflections

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

const exactsKeyword string = "[EXACTS]"

//EvaluateFilterMap convert map keys to snake_case
func EvaluateFilterMap(filter map[string]interface{}, t reflect.Type) {
	fields := GetDeclaredFields(t)
	fieldsMap := SliceOfFieldToMap(fields)
	for key, value := range filter {
		newKey := key
		//check if joinColumn
		var fieldKey string
		if strings.Contains(key, exactsKeyword) {
			fieldKey = strings.Replace(key, exactsKeyword, "", 1)
		} else {
			fieldKey = key
		}

		field, ok := fieldsMap[fieldKey]
		if ok {
			customTag, tagOk := GetMapOfTag(field, "custom")
			if tagOk {
				if customTag["foreignKey"] != "" {
					itemName := customTag["optionItemName"]

					log.Println("JOIN key: ", key, ".", itemName)
					newKey = ToSnakeCase(key, true) + "." + ToSnakeCase(itemName, true)
				} else {
					newKey = ToSnakeCase(key, true)
				}
			} else {
				newKey = ToSnakeCase(key, true)
			}

		} else {
			newKey = ToSnakeCase(key, true)
		}

		delete(filter, key)

		filter[newKey] = value
		fmt.Println("key: ", key, "value: ", value, "newKey: ", newKey)

	}
	fmt.Println("filter evaluated: ", filter)
}
