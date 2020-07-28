package app

import "github.com/fajaralmu/go_part4_web/entities"

type footer struct {
	Year    int
	Profile entities.Profile
}

type pageData struct {
	Title   string
	Message string
	Content interface{}
	Profile entities.Profile
	Footer  footer
}
