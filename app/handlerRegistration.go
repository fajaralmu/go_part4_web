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
	registerWebAPIRoutes()
	log.Println("END registerAPIs")
}

func handleAPI(path string, handlerMethod func(w http.ResponseWriter, r *http.Request) (entities.WebResponse, error), method string, authenticated bool) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("api-START///////////URI: ", r.RequestURI)
			preHandleResult := apiPreHandle(w, r, authenticated)

			if !preHandleResult {
				if authenticated {
					writeErrorMsgBadRequest(w, "Invalid Request 01")
					return
				}

				log.Println("API-END////////////URI: ", path)
				return
			}
			response, err := handlerMethod(w, r)
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

	///////////// WEB PAGES /////////////
	log.Println("START Register Web Pages")
	registerWebPageRoutes()

	///////////// Web Socket /////////////
	handleWebsocket("/chat", wsRoute)

	/////STATIC RESOURCES/////

	// fs := http.StripPrefix("/static/", fileServer())
	var cs *customStaticHandler = &customStaticHandler{root: http.Dir("./public/")}
	router.PathPrefix("/static/").Handler(cs)

	log.Println("END Register Web Pages")
}

// func fileServer() http.Handler {
// 	return http.FileServer(http.Dir("./public/"))
// }

func handleWebsocket(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, handler)
}

func handleMvc(path string, handlerMethod func(w http.ResponseWriter, r *http.Request) error, method string, authenticated bool) {

	h := appHandler{
		handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("MVC-START///////////URI: ", r.RequestURI)

			preHandleResult := mvcPreHandle(w, r, authenticated)
			if !preHandleResult {
				if authenticated {
					setLatestURI(w, r, r.RequestURI)
					sendRedirect(w, r, "/account/login")
					return
				}

				log.Println("MVC-END////////////URI: ", path)
				return
			}
			err := handlerMethod(w, r)
			if err != nil {
				writeErrorMsgBadRequest(w, err.Error())
			}
			log.Println("MVC-END////////////URI: ", r.RequestURI)
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
