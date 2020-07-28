package app

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

type footer struct {
	Year    int
	Profile entities.Profile
}

type header struct {
	Profile entities.Profile
	Pages   []entities.Page
}

type pageData struct {
	PageCode              string
	Title                 string
	Message               string
	Content               interface{}
	Profile               entities.Profile
	Footer                footer
	Header                header
	AdditionalStylePaths  []string
	AdditionalScriptPaths []string
}

func (pageData *pageData) setStylePath(paths ...string) {
	pageData.AdditionalStylePaths = paths
}

func (pageData *pageData) setScriptPath(paths ...string) {
	pageData.AdditionalScriptPaths = paths
}

func (pageData *pageData) setHeaderFooter() {
	pageData.Header = header{
		Profile: pageData.Profile,
		Pages:   getPages(),
	}
	pageData.Footer = footer{
		Year:    getCurrentYr(),
		Profile: pageData.Profile,
	}
}

func (pageData *pageData) prepareWebData() {
	pageData.Profile = getProfile()

	pageData.setHeaderFooter()
	pageData.parseContent()
}

func getPages() []entities.Page {
	filter := entities.Filter{}
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

func (pageData *pageData) parseContent() {

	tmpl2, _ := template.ParseFiles("./templates/pages/" + pageData.PageCode + ".html")
	var tpl bytes.Buffer

	e := tmpl2.ExecuteTemplate(&tpl, pageData.PageCode, pageData.Header)
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
