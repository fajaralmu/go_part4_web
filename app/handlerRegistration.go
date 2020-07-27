package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/gorilla/mux"
)

var router *mux.Router

func registerAPIs() {

	log.Println("START registerAPIs")

	router.HandleFunc("/api/entities", getEntities).Methods("POST")
	router.HandleFunc("/api/entities/add", addEntities).Methods("POST")
	// router.HandleFunc("/api/books", createBook).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("END registerAPIs")
}

func registerWebPages() {

	log.Println("START Register Web Pages")
	// router.HandleFunc("/home", app.homeRoute).Methods("GET")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/static/").Handler(fs)

	log.Println("END Register Web Pages")
}

func getEntities(w http.ResponseWriter, r *http.Request) {

	var request entities.WebRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}
	response := Filter(request)
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

/////////////////////////////////////////

func webResponse(code string, message string) entities.WebResponse {
	return entities.WebResponse{Code: code, Message: message, Date: time.Now()}
}

func writeWebResponse(w http.ResponseWriter, webResponse entities.WebResponse) {
	webResponse.Date = time.Now()
	if (webResponse.Code) == "" {
		webResponse.Code = "00"
	}
	if webResponse.Message == "" {
		webResponse.Message = "success"
	}
	writeJSONResponse(w, webResponse)
}

func writeJSONResponse(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func writeResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

}

func writeErrorMsgBadRequest(w http.ResponseWriter, msg string) {
	writeResponseHeaders(w)
	w.WriteHeader(400)
	writeJSONResponse(w, webResponse("400", msg))
}
