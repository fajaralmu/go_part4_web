package app

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Title   string
	Message string
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/home.html")

	if err == nil {
		pageData := PageData{
			Title:   "Welcome",
			Message: "Hello World",
		}
		tmpl.Execute(w, pageData)
	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, "Requested Web Page Not Found")
	}
}
