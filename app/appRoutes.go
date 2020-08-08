package app

import "net/http"

type appRoutes struct {
	homeRoute          func(w http.ResponseWriter, r *http.Request) error `custom:"path:/home;authenticated:true"`
	aboutRoute         func(w http.ResponseWriter, r *http.Request) error `custom:"path:/public/about;authenticated:false"`
	commonPageRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/page/{code};authenticated:true"`
	managementRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/management/{code};authenticated:true"`
	adminDasboardRoute func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/home;authenticated:true"`
	pageSettingRoute   func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/pagesettings;authenticated:true"`
	resetMenusRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/admin/resetmenus;authenticated:true"`

	loginRoute    func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/login;authenticated:false"`
	registerRoute func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/register;authenticated:false"`
	logoutRoute   func(w http.ResponseWriter, r *http.Request) error `custom:"path:/account/logout;authenticated:false"`
}
