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
var entityConfigMap map[string]reflect.Type = map[string]reflect.Type{}

func Init() {
	router = mux.NewRouter()
	initiated = true
	putConfig([]entities.User{},
		[]entities.UserRole{},
		[]entities.RegisteredRequest{},
		[]entities.Menu{},
		[]entities.Page{},
		[]entities.Profile{})
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

func putConfig(t ...interface{}) {

	for _, item := range t {
		_type := reflect.TypeOf(item)

		key := reflections.ToSnakeCase(_type.Elem().Name())
		// fmt.Println("KEY:", key, "_type: ", _type)
		entityConfigMap[key] = _type
	}

}
