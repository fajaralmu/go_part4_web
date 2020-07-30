package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getEntities(w http.ResponseWriter, r *http.Request) {
	log.Println("///////////////////////START API getEntities////////////////")

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())

	} else {
		response, err := Filter(request)
		if nil == err {
			writeWebResponse(w, response)
		} else {
			writeErrorMsgBadRequest(w, err.Error())
		}
	}

	log.Println("///////////////////////END API getEntities////////////////")

}

func deleteEntities(w http.ResponseWriter, r *http.Request) {
	log.Println("///////////////////////START API deleteEntities////////////////")

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
	} else {
		response := Delete(request)
		writeWebResponse(w, response)
	}

	log.Println("///////////////////////END API deleteEntities////////////////")
}

func addEntities(w http.ResponseWriter, r *http.Request) {
	log.Println("///////////////////////START API addEntities////////////////")

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
	} else {
		response := AddEntity(request)
		writeWebResponse(w, response)
	}

	log.Println("///////////////////////END API addEntities////////////////")
}

func updateEntities(w http.ResponseWriter, r *http.Request) {
	log.Println("///////////////////////START API updateEntities////////////////")

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
	} else {
		response := UpdateEnity(request)
		writeWebResponse(w, response)
	}
	log.Println("///////////////////////END API updateEntities////////////////")
}
