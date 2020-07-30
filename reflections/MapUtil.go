package reflections

import "fmt"

//EvaluateFilterMap convert map keys to snake_case
func EvaluateFilterMap(filter map[string]interface{}) {

	for key, value := range filter {
		newKey := ToSnakeCase(key, true)
		delete(filter, key)

		filter[newKey] = value
		fmt.Println("key: ", key, "value: ", value, "newKey: ", newKey)

	}
	fmt.Println("filter evaluated: ", filter)
}
