package app

import (
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

type webAPIRoute struct {
	GetEntities       func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities;authenticated:true"`
	AddEntities       func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/add;authenticated:true"`
	UpdateEntities    func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/update;authenticated:true"`
	DeleteEntities    func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/entities/delete;authenticated:true"`
	SavePageSequence  func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/admin/savepagesequence;authenticated:true"`
	PrintModelsReport func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/report/entities;authenticated:true;jsonResponse:false"`

	Login         func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/login;authenticated:false"`
	CheckUserName func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/checkusername;authenticated:false"`
	Register      func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) `custom:"path:/api/account/register;authenticated:false"`
}

func registerWebAPIRoutes() {
	appRoute := webAPIRoute{

		GetEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return getEntitiesREST(w, r)
		},
		AddEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return addEntitiesREST(w, r)
		},
		UpdateEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return updateEntitiesREST(w, r)
		},
		DeleteEntities: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return deleteEntitiesREST(w, r)
		},
		SavePageSequence: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return savePageSequenceREST(w, r)
		},
		Login: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return loginREST(w, r)
		},
		CheckUserName: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return checkUserNameREST(w, r)
		},
		Register: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return registerREST(w, r)
		},
		PrintModelsReport: func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error) {
			return printModelReportREST(w, r)
		},
	}
	registerHandlers(appRoute, "api")

}
