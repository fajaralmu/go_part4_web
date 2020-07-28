package app

import (
	"fmt"
	"log"
	"reflect"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/entities"
)

func UpdateEnity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	repository.Save(fieldValue.(entities.InterfaceEntity))
	fmt.Println("SAVED Entity: ", fieldValue)

	response := entities.WebResponse{
		Result: fieldValue,
	}
	return response
}

func Delete(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	deleted := repository.Delete(fieldValue.(entities.InterfaceEntity))
	fmt.Println("Deleted Entity: ", deleted)

	response := entities.WebResponse{
		Result: deleted,
	}
	return response
}

func AddEntity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	repository.CreateNew(fieldValue.(entities.InterfaceEntity))
	fmt.Println("created Entity: ", fieldValue)

	response := entities.WebResponse{
		Result: fieldValue,
	}
	return response
}

//Filter returns entities by given keywords
func Filter(request entities.WebRequest) entities.WebResponse {

	filter := request.Filter
	entityType := entityConfigMap[filter.EntityName]

	createdSlice := reflections.CreateNewType(entityType)
	fmt.Println("--createdSlice--: ", createdSlice)

	list, totalData := repository.Filter(createdSlice, filter)

	response := entities.WebResponse{
		ResultList:     list,
		TotalData:      totalData,
		AdditionalData: reflections.CreateEntityProperty(reflect.TypeOf(entities.User{}), map[string][]interface{}{}),
	}
	fmt.Println("RESPONSE: ", response)
	return response
}
