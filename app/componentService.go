package app

import (
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/repository"
)

func updatePageSequence(req entities.WebRequest) bool {
	pages := req.Pages

	for _, page := range pages {
		repository.Save(page)
	}
	return true

}
