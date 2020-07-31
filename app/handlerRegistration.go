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

	handleAPI("/api/entities", getEntities, "POST")
	handleAPI("/api/entities/add", addEntities, "POST")
	handleAPI("/api/entities/update", updateEntities, "POST")
	handleAPI("/api/entities/delete", deleteEntities, "POST")
	// router.HandleFunc("/api/books", createBook).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Println("END registerAPIs")
}

func handleAPI(path string, handler func(w http.ResponseWriter, r *http.Request), method string) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("api-START///////////URI: ", r.RequestURI)
			if apiPreHandle(w, r) == false {
				log.Println("mvc-END////////////URI: ", path)
				return
			}
			handler(w, r)
			log.Println("api-END////////////URI: ", r.RequestURI)
		},
	}

	router.HandleFunc(path, h.handler).Methods(method)

}

func registerWebPages() {

	log.Println("START Register Web Pages")
	handleMvc("/home", homeRoute, "GET")
	handleMvc("/page/{code}", commonPageRoute, "GET")
	handleMvc("/management/{code}", managementRoute, "GET")

	handleMvc("/account/login", loginRoute, "GET")
	handleMvc("/account/register", registerRoute, "GET")

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/static/").Handler(fs)

	log.Println("END Register Web Pages")
}

func handleMvc(path string, handler func(w http.ResponseWriter, r *http.Request), method string) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("mvc-START///////////URI: ", r.RequestURI)
			if mvcPreHandle(w, r) == false {
				log.Println("mvc-END////////////URI: ", path)
				return
			}
			handler(w, r)
			log.Println("mvc-END////////////URI: ", r.RequestURI)
		},
	}

	router.HandleFunc(path, h.handler).Methods(method)

}

type appHandler struct {
	handler func(w http.ResponseWriter, r *http.Request)
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
	writeResponseHeaders(w)
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
