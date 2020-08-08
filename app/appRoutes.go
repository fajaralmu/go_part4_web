package app

import (
	"net/http"
	"reflect"

	"github.com/fajaralmu/go_part4_web/reflections"
)

type appRoutes struct {
	HomeRoute           func(w http.ResponseWriter, r *http.Request) error `custom:"path:/home;authenticated:true"`
	AboutRoute          func(w http.ResponseWriter, r *http.Request) error `custom:"path:/public/about;authenticated:false"`
	CommonPageRoute     func(w http.ResponseWriter, r *http.Request) error `custom:"path:/page/{code};authenticated:true"`
	ManagementRoute     func(w http.ResponseWriter, r *http.Request) error `custom:"path:/management/{code};authenticated:true"`
	AdminDashboardRoute func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/home;authenticated:true"`
	PageSettingRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/pagesettings;authenticated:true"`
	ResetMenusRoute     func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/resetmenus;authenticated:true"`

	LoginRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/login;authenticated:false"`
	RegisterRoute func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/register;authenticated:false"`
	LogoutRoute   func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/logout;authenticated:false"`
}

func registerRoutes() {
	appRoute := appRoutes{}
	appRoute.HomeRoute = func(w http.ResponseWriter, r *http.Request) error {
		return homeRoute(w, r)
	}
	appRoute.ManagementRoute = func(w http.ResponseWriter, r *http.Request) error {
		return managementRoute(w, r)
	}

	appRoute.AboutRoute = appRoute.HomeRoute
	appRoute.AdminDashboardRoute = appRoute.HomeRoute

	appRoute.CommonPageRoute = func(w http.ResponseWriter, r *http.Request) error {
		return commonPageRoute(w, r)
	}
	appRoute.LoginRoute = func(w http.ResponseWriter, r *http.Request) error {
		return loginRoute(w, r)
	}
	appRoute.LogoutRoute = func(w http.ResponseWriter, r *http.Request) error {
		return logoutRoute(w, r)
	}
	appRoute.PageSettingRoute = func(w http.ResponseWriter, r *http.Request) error {
		return pageSettingRoute(w, r)
	}
	appRoute.ResetMenusRoute = func(w http.ResponseWriter, r *http.Request) error {
		return resetMenus(w, r)
	}
	appRoute.RegisterRoute = func(w http.ResponseWriter, r *http.Request) error {
		return registerRoute(w, r)
	}
	registerHandlers(appRoute)
}

func registerHandlers(appRoute appRoutes) {
	fields := reflections.GetDeclaredFields(reflect.TypeOf(appRoute))
	for _, field := range fields {
		if field.Type.Kind() == reflect.Func {
			customTag, _ := reflections.GetMapOfTag(field, "custom")

			path := customTag["path"]
			authenticated := customTag["authenticated"] == "true"
			fieldValue, _ := reflections.GetFieldValue(field.Name, appRoute)
			handleMvc(path, fieldValue.(func(w http.ResponseWriter, r *http.Request) error), "GET", authenticated)
		}
	}

}
