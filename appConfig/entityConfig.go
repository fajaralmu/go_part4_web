package appConfig

import (
	"log"
	"reflect"
)

var entityConfigMap map[string]*EntityConfig = map[string]*EntityConfig{}

type EntityConfig struct {
	Name       string
	ListType   reflect.Type
	SingleType reflect.Type
}

//GetEntityConf returns *entityConfig
func GetEntityConf(key string) *EntityConfig {
	return entityConfigMap[key]
}

func PutConfig(t ...*EntityConfig) {

	for _, item := range t {
		log.Println("put entity Config: ", item.Name)
		entityConfigMap[item.Name] = item
	}

}
