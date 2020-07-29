package app

import (
	"fmt"
	"net/http"
	"reflect"
	"text/template"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/repository"
	"github.com/gorilla/mux"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getWebFiles() []string {
	return []string{
		"./templates/BASE_PAGE.html",
		"./templates/error/notFound.html",
		"./templates/include/head.html",
		"./templates/include/foot.html",
	}
}

func getProfile() entities.Profile {
	profiles := repository.FilterByKey(&[]entities.Profile{}, "AppCode", "123")
	return profiles[0].(entities.Profile)
}

func getMuxParam(r *http.Request, param string) string {
	params := mux.Vars(r)
	paramValue := params[param]
	return paramValue
}

func getPageByCode(code string) entities.Page {
	fmt.Println("___________________________getPageByCode:", code)
	pageList := repository.FilterByKey(&[]entities.Page{}, "Code", code)

	page := pageList[0].(entities.Page)

	menuList := repository.FilterByKey(&[]entities.Menu{}, "PageID", page.ID)
	fmt.Println("____________________________menuList size: ", len(menuList))
	page.Menus = toSliceOfMenu(menuList)
	return page

}

func toSliceOfMenu(menus []interface{}) []entities.Menu {
	result := []entities.Menu{}
	for _, item := range menus {
		result = append(result, item.(entities.Menu))
	}
	return result
}

func managementRoute(w http.ResponseWriter, r *http.Request) {
	managementFiles := getWebFiles()
	managementFiles = append(managementFiles)
	tmpl, err := template.ParseFiles(managementFiles...)
	entityProperty := reflections.CreateEntityProperty(reflect.TypeOf(entities.Page{}), nil)
	if err == nil {
		pageData := pageData{
			PageCode:       "entityManagementPage",
			Title:          "Management Page",
			Message:        "Manage Models",
			EntityProperty: entityProperty,
			AdditionalPages: []string{
				"./templates/entity-management-component/detail-element.html", "./templates/entity-management-component/form-element.html",
			},
		}
		pageData.prepareWebData()

		tmpl.ExecuteTemplate(w, "layout", pageData)

	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, err.Error())
	}
}

func commonPageRoute(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(getWebFiles()...)
	pageCode := getMuxParam(r, "code")

	selectedPage := getPageByCode(pageCode)

	if err == nil {
		pageData := pageData{
			PageCode: "commonPage",
			Title:    "Common Page",
			Message:  "Hello World",
			Page:     selectedPage,
		}
		pageData.prepareWebData()

		tmpl.ExecuteTemplate(w, "layout", pageData)

	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, err.Error())
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(getWebFiles()...)

	if err == nil {
		pageData := pageData{
			PageCode: "about",
			Title:    "Welcome",
			Message:  "Hello World",
		}
		pageData.prepareWebData()
		pageData.setStylePath("about")

		tmpl.ExecuteTemplate(w, "layout", pageData)

	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, err.Error())
	}
}
