package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

/////////////////////////MODEL CRUD///////////////////////////

func getEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	broadcast <- "STARTS getEntities"
	if err != nil {
		return response, err
	}
	response, err = Filter(request)
	broadcast <- "END getEntities"
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
	response = UpdateEntity(request)
	return response, err
}

func savePageSequence(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {
	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	updatePageSequence(request)
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
		var latestURI string = getLatestURI(w, r)
		if latestURI != "" {
			// w.WriteHeader(302)
			log.Println(" getLatestURI(w, r) : ", latestURI)
			w.Header().Add("location", latestURI)
		}
	}
	return response, err

}
