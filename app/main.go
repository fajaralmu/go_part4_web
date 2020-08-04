package app

import (
	"log"
	"net/http"
	"reflect"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/gorilla/mux"
)

var initiated bool = false

func newEConf(single interface{}, list interface{}, FormInputColumn int) *appConfig.EntityConfig {
	singleType := reflect.TypeOf(single)
	log.Println("create CONFIG: ", singleType.Name())
	return &appConfig.EntityConfig{
		Name:            reflections.ToSnakeCase(singleType.Name(), true),
		ListType:        reflect.TypeOf(list),
		SingleType:      reflect.TypeOf(single),
		FormInputColumn: FormInputColumn,
	}
}

func Init() {
	router = mux.NewRouter()
	initiated = true
	appConfig.PutConfig(newEConf(entities.User{}, []entities.User{}, 1),
		newEConf(entities.UserRole{}, []entities.UserRole{}, 1),
		newEConf(entities.RegisteredRequest{}, []entities.RegisteredRequest{}, 2),
		newEConf(entities.Menu{}, []entities.Menu{}, 2),
		newEConf(entities.Page{}, []entities.Page{}, 2),
		newEConf(entities.Profile{}, []entities.Profile{}, 1))
	registerSessions()
	initMenus()
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
