package reflections

import "fmt"

//EvaluateFilterMap convert map keys to snake_case
func EvaluateFilterMap(filter map[string]interface{}) {

	for key, value := range filter {
		fmt.Println("key: ", key, "value: ", value)
		newKey := ToSnakeCase(key)
		delete(filter, key)

		filter[newKey] = value

	}
	fmt.Println("filter evaluated: ", filter)
}
