package app

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"text/template"

	"github.com/fajaralmu/go_part4_web/appConfig"

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
	if len(profiles) > 0 {
		return profiles[0].(entities.Profile)
	}
	return entities.Profile{
		Name: "Undefined",
	}
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

func managementRoute(w http.ResponseWriter, r *http.Request) error {

	entityCode := getMuxParam(r, "code")

	if "" == entityCode {

		return errors.New("Invalid Request, entityCode is EMPTY")
	}
	entityConf := appConfig.GetEntityConf(entityCode)
	if nil == entityConf {

		return errors.New("Invalid Request, entityConf Not Found")
	}
	entityProperty := appConfig.CreateEntityProperty(entityConf.SingleType)

	pageData := pageData{
		PageCode:       "entityManagementPage",
		Title:          "Management Page",
		Message:        "Manage Models",
		EntityProperty: entityProperty,
		AdditionalPages: []string{
			"./templates/entity-management-component/detail-element.html", "./templates/entity-management-component/form-element.html",
		},
	}
	return executeWebContents(pageData, w, r)
}

func resetMenus(w http.ResponseWriter, r *http.Request) (err error) {

	resetAllMenus()
	sendRedirect(w, r, "/admin/home")
	return nil

}

func executeWebContents(p pageData, w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles(getWebFiles()...)
	t := (reflect.ValueOf(p))

	if err == nil && t.IsValid() {
		p.w = w
		p.r = r
		p.prepareWebData()
		tmpl.ExecuteTemplate(w, "layout", p)
		return nil
	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, err.Error())
		return err
	}
}

func commonPageRoute(w http.ResponseWriter, r *http.Request) error {

	pageCode := getMuxParam(r, "code")
	selectedPage := getPageByCode(pageCode)

	pageData := pageData{
		PageCode: "commonPage",
		Title:    "Common Page",
		Message:  "Hello World",
		Page:     selectedPage,
	}
	return executeWebContents(pageData, w, r)

}

func loginRoute(w http.ResponseWriter, r *http.Request) error {

	pageData := pageData{
		PageCode: "login",
		Title:    "Login Page",
	}
	return executeWebContents(pageData, w, r)
}
func logoutRoute(w http.ResponseWriter, r *http.Request) error {

	setUserToSession(w, r, nil)
	sendRedirect(w, r, "/account/login")
	return nil
}
func registerRoute(w http.ResponseWriter, r *http.Request) error {

	pageData := pageData{
		PageCode: "register",
		Title:    "Register Page",
	}
	return executeWebContents(pageData, w, r)
}

func pageSettingRoute(w http.ResponseWriter, r *http.Request) error {

	pageData := pageData{
		PageCode: "pageSequence",
		Title:    "Page Setting",
	}
	pageData.setStylePath("pagesequence")
	return executeWebContents(pageData, w, r)
}

func homeRoute(w http.ResponseWriter, r *http.Request) error {

	pageData := pageData{
		PageCode: "about",
		Title:    "About Us",
		Message:  "Hello World",
	}
	pageData.setStylePath("about")
	return executeWebContents(pageData, w, r)
}
