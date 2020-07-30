package app

import (
	"log"
	"net/http"
	"reflect"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/gorilla/mux"
)

var initiated bool = false
var entityConfigMap map[string]*entityConfig = map[string]*entityConfig{}

type entityConfig struct {
	name       string
	listType   reflect.Type
	singleType reflect.Type
}

func getEConf(single interface{}, list interface{}) *entityConfig {
	singleType := reflect.TypeOf(single)
	log.Println("create CONFIG: ", singleType.Name())
	return &entityConfig{
		name:       reflections.ToSnakeCase(singleType.Name()),
		listType:   reflect.TypeOf(list),
		singleType: reflect.TypeOf(single),
	}
}

func Init() {
	router = mux.NewRouter()
	initiated = true
	putConfig(getEConf(entities.User{}, []entities.User{}),
		getEConf(entities.UserRole{}, []entities.UserRole{}),
		getEConf(entities.RegisteredRequest{}, []entities.RegisteredRequest{}),
		getEConf(entities.Menu{}, []entities.Menu{}),
		getEConf(entities.Page{}, []entities.Page{}),
		getEConf(entities.Profile{}, []entities.Profile{}))
}

func Run() {
	initWebApp()

	// webReq := entities.WebRequest{
	// 	Filter: entities.Filter{
	// 		EntityName: "user",
	// 		Page:       0,
	// 		Limit:      3,
	// 		FieldsFilter: map[string]interface{}{
	// 			"Username":    "Fajar",
	// 			"DisplayName": "Fajr2",
	// 		},
	// 	},
	// }
	// Filter(webReq)

}

func initWebApp() {
	registerAPIs()
	registerWebPages()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func putConfig(t ...*entityConfig) {

	for _, item := range t {
		log.Println("put entity Config: ", item.name)
		entityConfigMap[item.name] = item
	}

}
