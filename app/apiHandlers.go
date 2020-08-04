package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

/////////////////////////MODEL CRUD///////////////////////////

func getEntities(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	sendBroadcastMessage("start GET entities")
	if err != nil {
		return response, err
	}
	response, err = Filter(request)
	sendBroadcastMessage("end GET entities")
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

func checkUserName(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {
	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	available := isUsernameAvailable(request.User.Username)
	log.Println("isUsernameAvailable: ", available)
	if available {
		return webResponse("00", "username available"), nil
	}
	return webResponse("01", "username unavailable"), nil
}
func register(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {
	var request entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return response, err
	}
	user := request.User
	repository.CreateNewWithoutValidation(user)

	return webResponse("00", "success"), nil

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
