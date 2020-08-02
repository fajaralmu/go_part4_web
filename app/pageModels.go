package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

type footer struct {
	Year    int
	Profile entities.Profile
}

type header struct {
	Profile       entities.Profile
	Pages         []entities.Page
	User          *entities.User
	Authenticated bool
	Greeting      string
}

type pageData struct {
	PageCode              string
	RequestID             string
	Title                 string
	Message               string
	Content               interface{}
	Profile               entities.Profile
	Footer                footer
	Header                header
	AdditionalStylePaths  []string
	AdditionalScriptPaths []string
	AdditionalPages       []string
	Page                  entities.Page
	EntityProperty        appConfig.EntityProperty

	//private
	r *http.Request
	w http.ResponseWriter
}

func (p *pageData) setStylePath(paths ...string) {
	p.AdditionalStylePaths = paths
}

func (p *pageData) setScriptPath(paths ...string) {
	p.AdditionalScriptPaths = paths
}

func (p *pageData) setHeaderFooter() {
	p.RequestID = reflections.RandomNum(15)
	p.Header = header{
		Profile:  p.Profile,
		Pages:    getPages(p.w, p.r),
		Greeting: reflections.GetTimeGreeting(),
	}
	p.Footer = footer{
		Year:    getCurrentYr(),
		Profile: p.Profile,
	}
	var loggedUser *entities.User

	if p.r != nil {
		loggedUser = getUserFromSession(p.w, p.r)
		if nil != loggedUser {
			p.Header.User = loggedUser
		}
	}

	p.Header.Authenticated = loggedUser != nil

}

func (pageData *pageData) prepareWebData() {
	pageData.Profile = getProfile()

	pageData.setHeaderFooter()
	pageData.parseContent(pageData.AdditionalPages...)
}

func getPages(w http.ResponseWriter, r *http.Request) []entities.Page {
	sessionValid := false
	if r != nil {
		sessionValid = validateSessionn(w, r)
	}

	filter := entities.Filter{}
	if !sessionValid {
		filter.Exacts = true
		filter.FieldsFilter = map[string]interface{}{
			"Authorized": 0,
		}

	} else {

	}

	filter.OrderBy = "Sequence"

	list, count := repository.Filter(&[]entities.Page{}, filter)
	fmt.Println("Total Pages: ", count)
	return toSliceOfPage(list)
}

func toSliceOfPage(list []interface{}) []entities.Page {

	pages := []entities.Page{}

	for _, item := range list {
		pages = append(pages, item.(entities.Page))
	}
	return pages
}

func (pageData *pageData) parseContent(additionalPage ...string) {

	webPages := []string{
		"./templates/pages/" + pageData.PageCode + ".html",
	}
	webPages = append(webPages, additionalPage...)

	tmpl2, err := template.ParseFiles(webPages...)
	var tpl bytes.Buffer

	if err != nil {
		log.Println("tmpl2 ERR: ", err.Error())
	}
	log.Println("********pageData.PageCode: ", pageData.PageCode)

	e := tmpl2.ExecuteTemplate(&tpl, pageData.PageCode, pageData)
	if e == nil {

		fmt.Println("Success parsing ", pageData.PageCode)
		pageData.Content = tpl.String()
	} else {
		fmt.Println("Error parsing web content: ", e.Error())
		pageData.PageCode = ""
	}
}

func getCurrentTime() (int, time.Month, int) {
	fmt.Println("CURRENT TIME: ", time.Now())
	return time.Now().Date()
}

func getCurrentYr() int {
	yr, _, _ := getCurrentTime()
	fmt.Println("year: ", yr)
	return yr
}
