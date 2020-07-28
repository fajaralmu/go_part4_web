package app

import (
	"net/http"
	"text/template"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getWebFiles() []string {
	return []string{
		"./templates/BASE_PAGE.html",
		"./templates/error/notFound.html",
		"./templates/include/head.html",
		"./templates/include/foot.html",
	}
}

func getProfile() entities.Profile {
	profiles := repository.FilterByKey(&[]entities.Profile{}, "AppCode", "123")
	return profiles[0].(entities.Profile)
}

func homeRoute(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(getWebFiles()...)

	if err == nil {
		pageData := pageData{
			PageCode: "about",
			Title:    "Welcome",
			Message:  "Hello World",
		}
		pageData.prepareWebData()
		pageData.setStylePath("about")

		tmpl.ExecuteTemplate(w, "layout", pageData)

	} else {
		writeResponseHeaders(w)
		writeErrorMsgBadRequest(w, err.Error())
	}
}
