package app

import (
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/repository"
)

func updatePageSequence(req entities.WebRequest) bool {
	pages := req.Pages

	for i, page := range pages {
		page.Sequence = i
		repository.SaveWihoutValidation(&page)
	}
	return true

}
