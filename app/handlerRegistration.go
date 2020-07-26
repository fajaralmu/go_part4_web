package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/gorilla/mux"
)

var router *mux.Router

func registerAPIs() {

	log.Println("START registerAPIs")

	router.HandleFunc("/api/entities", getEntities).Methods("GET")
	// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
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
	_ = json.NewDecoder(r.Body).Decode(&request)
	response := Filter(request)
	writeJSONResponse(w, response)
}

/////////////////////////////////////////

func writeJSONResponse(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func writeResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

}

func writeErrorMsg(w http.ResponseWriter, msg string) {
	w.WriteHeader(404)
	writeJSONResponse(w, entities.WebResponse{Code: "404", Message: msg})
}
