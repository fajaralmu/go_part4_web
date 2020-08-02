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

	handleAPI("/api/entities", getEntities, "POST", true)
	handleAPI("/api/entities/add", addEntities, "POST", true)
	handleAPI("/api/entities/update", updateEntities, "POST", true)
	handleAPI("/api/entities/delete", deleteEntities, "POST", true)
	handleAPI("/api/admin/savepagesequence", savePageSequence, "POST", true)
	// router.HandleFunc("/api/books", createBook).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	handleAPI("/api/account/login", login, "POST", false)

	log.Println("END registerAPIs")
}

func handleAPI(path string, handler func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error), method string, authenticated bool) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("api-START///////////URI: ", r.RequestURI)
			preHandleResult := apiPreHandle(w, r, authenticated)
			if preHandleResult == false {
				if authenticated {

					writeErrorMsgBadRequest(w, "Invalid Request 01")
					return
				}

				log.Println("api-END////////////URI: ", path)
				return
			}
			response, err := handler(w, r)
			if nil != err {
				writeErrorMsgBadRequest(w, err.Error())
			} else {
				writeWebResponse(w, response)
			}
			log.Println("API-END////////////URI: ", r.RequestURI)
		},
	}

	router.HandleFunc(path, h.handler).Methods(method)

}

func registerWebPages() {

	log.Println("START Register Web Pages")
	handleMvc("/home", homeRoute, "GET", false)
	handleMvc("/page/{code}", commonPageRoute, "GET", true)
	handleMvc("/management/{code}", managementRoute, "GET", true)
	handleMvc("/admin/home", homeRoute, "GET", false)
	handleMvc("/admin/pagesettings", pageSettingRoute, "GET", true)

	handleMvc("/account/login", loginRoute, "GET", false)
	handleMvc("/account/register", registerRoute, "GET", false)
	handleMvc("/account/logout", logoutRoute, "GET", false)
	/////static resources/////
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/static/").Handler(fs)

	log.Println("END Register Web Pages")
}

func handleMvc(path string, handler func(w http.ResponseWriter, r *http.Request) error, method string, authenticated bool) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("mvc-START///////////URI: ", r.RequestURI)

			preHandleResult := mvcPreHandle(w, r, authenticated)
			if preHandleResult == false {
				if authenticated {
					setLatestURI(w, r, r.RequestURI)
					sendRedirect(w, r, "/account/login")
					return
				}

				log.Println("mvc-END////////////URI: ", path)
				return
			}
			err := handler(w, r)
			if err != nil {
				writeErrorMsgBadRequest(w, err.Error())
			}
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
