package appConfig

import (
	"log"
	"reflect"
	"strings"
)

var entityConfigMap map[string]*EntityConfig = map[string]*EntityConfig{}

//EntityConfig holds CRUD related information of the model
type EntityConfig struct {
	Name            string
	ListType        reflect.Type
	SingleType      reflect.Type
	FormInputColumn int
}

//GetEntityConf returns *entityConfig
func GetEntityConf(key string) *EntityConfig {
	return entityConfigMap[strings.ToLower(key)]
}

//GetEntitiesTypes returns slice of entity types from entityConfigMap
func GetEntitiesTypes() (types []reflect.Type) {
	for _, val := range entityConfigMap {
		types = append(types, val.SingleType)
	}
	return types
}

//PutConfig add config to the entityConfigMap
func PutConfig(t ...*EntityConfig) {

	for _, item := range t {

		lowerCased := strings.ToLower(item.Name)
		log.Println("put entity Config: ", lowerCased)
		entityConfigMap[lowerCased] = item
	}

}
