package app

import (
	"encoding/json"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

/////////////////////////MODEL CRUD///////////////////////////

func getEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	response, err = Filter(request)
	return response, err

}

func deleteEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	response = Delete(request)
	writeWebResponse(w, response)
	return response, nil

}

func addEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	response = AddEntity(request)
	return response, err
}

func updateEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	response = UpdateEnity(request)
	return response, err
}

/////////////////ACCOUNT////////////////
func login(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	response, err = Login(request, w, r)
	if err == nil {
		w.Header().Add("location", "/admin/home")
	}
	return response, err

}
