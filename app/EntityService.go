package app

import (
	"fmt"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/entities"
)

func AddEntity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	fieldValue, _ := reflections.GetFieldValue(entityName, request)
	repository.CreateNew(fieldValue.(entities.InterfaceEntity))
	fmt.Println("created Entity: ", fieldValue)
	response := entities.WebResponse{
		Result: fieldValue,
	}
	return response
}

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
