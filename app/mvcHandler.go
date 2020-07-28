package app

import (
	"net/http"
	"text/template"
	"time"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getCurrentTime() (int, time.Month, int) {
	return time.Now().Date()
}

func getCurrentYr() int {
	yr, _, _ := getCurrentTime()
	return yr
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/BASE_PAGE.html", "./templates/include/head.html", "./templates/include/foot.html")

	if err == nil {
		pageData := pageData{
			Title:   "Welcome",
			Message: "Hello World",
			Profile: entities.Profile{
				Name:             "CRUD WEB APP",
				Color:            "orange",
				ShortDescription: "HELLOW WORLD........",
			},
		}

		pageData.Footer = footer{
			Year:    getCurrentYr(),
			Profile: pageData.Profile,
		}

		tmpl.ExecuteTemplate(w, "layout", pageData)
	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, "Requested Web Page Not Found")
	}
}
