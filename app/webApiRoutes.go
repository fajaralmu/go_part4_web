package app

import (
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

type webAPIRoute struct {
	GetEntities      func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities;authenticated:true"`
	AddEntities      func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/add;authenticated:true"`
	UpdateEntities   func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/update;authenticated:true"`
	DeleteEntities   func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/delete;authenticated:true"`
	SavePageSequence func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/admin/savepagesequence;authenticated:true"`

	Login         func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/login;authenticated:false"`
	CheckUserName func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/checkusername;authenticated:false"`
	Register      func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/register;authenticated:false"`
}

func registerWebAPIRoutes() {
	appRoute := webAPIRoute{
		GetEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return getEntities(w, r)
		},
		AddEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return addEntities(w, r)
		},
		UpdateEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return updateEntities(w, r)
		},
		DeleteEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return deleteEntities(w, r)
		},
		SavePageSequence: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return savePageSequence(w, r)
		},

		Login: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return login(w, r)
		},
		CheckUserName: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return checkUserName(w, r)
		},
		Register: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return register(w, r)
		},
	}
	registerHandlers(appRoute, "api")

}
