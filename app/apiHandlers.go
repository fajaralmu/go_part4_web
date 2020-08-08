package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

/////////////////////////MODEL CRUD///////////////////////////

func getEntitiesREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {

	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	sendBroadcastMessage("start GET entities")
	if err != nil {
		return res, err
	}
	res, err = filterEntity(req)
	sendBroadcastMessage("end GET entities")
	return res, err

}

func deleteEntitiesREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {

	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return res, err
	}

	res = deleteEntity(req)
	writeWebResponse(w, res)
	return res, nil

}

func addEntitiesREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {

	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return res, err
	}
	res = addEntity(req)
	return res, err
}

func updateEntitiesREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {

	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return res, err
	}
	res = updateEntity(req)
	return res, err
}

func savePageSequenceREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {
	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return res, err
	}
	updatePageSequence(req)
	return res, err

}

func checkUserNameREST(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {
	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return response, err
	}
	available := isUsernameAvailable(req.User.Username)
	log.Println("isUsernameAvailable: ", available)
	if available {
		return webResponse("00", "username available"), nil
	}
	return webResponse("01", "username unavailable"), nil
}
func registerREST(w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {
	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return response, err
	}
	user := req.User
	repository.CreateNewWithoutValidation(user)

	return webResponse("00", "success"), nil

}

/////////////////ACCOUNT////////////////
func loginREST(w http.ResponseWriter, r *http.Request) (res entities.WebResponse, err error) {

	var req entities.WebRequest
	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return res, err
	}
	res, err = Login(req, w, r)
	if err == nil {
		var latestURI string = getLatestURI(w, r)
		if latestURI != "" {
			// w.WriteHeader(302)
			log.Println(" getLatestURI(w, r) : ", latestURI)
			w.Header().Add("location", latestURI)
		}
	}
	return res, err

}
