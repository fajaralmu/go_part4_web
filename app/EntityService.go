package app

import (
	"fmt"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/entities"
)

func Filter(request entities.WebRequest) entities.WebResponse {

	filter := request.Filter
	entityType := entityConfigMap[filter.EntityName]

	createdSlice := reflections.CreateNewType(entityType)
	fmt.Println("--createdSlice--: ", createdSlice)

	list, totalData := repository.Filter(createdSlice, filter)

	response := entities.WebResponse{
		ResultList: list,
		TotalData:  totalData,
	}
	fmt.Println("RESPONSE: ", response)
	return response
}
