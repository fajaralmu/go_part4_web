package app

import (
	"encoding/json"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getEntities(w http.ResponseWriter, r *http.Request) {

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}
	response, err := Filter(request)
	if nil == err {
		writeWebResponse(w, response)
	} else {
		writeErrorMsgBadRequest(w, err.Error())
	}

}

func deleteEntities(w http.ResponseWriter, r *http.Request) {

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}
	response := Delete(request)
	writeWebResponse(w, response)
}

func addEntities(w http.ResponseWriter, r *http.Request) {

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}
	response := AddEntity(request)
	writeWebResponse(w, response)
}

func updateEntities(w http.ResponseWriter, r *http.Request) {

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}
	response := UpdateEnity(request)
	writeWebResponse(w, response)
}
