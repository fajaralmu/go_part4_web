package app

import (
	"errors"
	"fmt"
	"log"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/entities"
)

func updateEntity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("UPDATE entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	repository.SaveAndValidate(fieldValue.(entities.InterfaceEntity))
	fmt.Println("SAVED Entity: ", fieldValue)

	res := entities.WebResponse{
		Result: fieldValue,
	}
	return res
}

func deleteEntity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("DEL entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	deleted := repository.Delete(fieldValue.(entities.InterfaceEntity), true)
	fmt.Println("Deleted Entity: ", deleted)

	res := entities.WebResponse{
		Result: deleted,
	}
	return res
}

func addEntity(request entities.WebRequest) entities.WebResponse {
	entityName := request.Filter.EntityName
	log.Println("ADD entityName: ", entityName)
	fieldValue, _ := reflections.GetFieldValue(entityName, &request)

	repository.CreateNew(fieldValue.(entities.InterfaceEntity))
	fmt.Println("created Entity: ", fieldValue)

	res := entities.WebResponse{
		Result: fieldValue,
	}
	return res
}

func filterEntity(request entities.WebRequest) (res entities.WebResponse, err error) {

	filter := request.Filter
	entityConf := appConfig.GetEntityConf(filter.EntityName)

	if nil == entityConf {
		return res, errors.New("Invalid entityName")
	}

	createdSlice := reflections.CreateNewType(entityConf.ListType)
	fmt.Println("--createdSlice--: ", createdSlice)

	list, totalData := repository.Filter(createdSlice, filter)

	res = entities.WebResponse{
		ResultList: list,
		TotalData:  totalData,
		Filter:     filter,
		//AdditionalData: appConfig.CreateEntityProperty(reflect.TypeOf(entities.User{})),
	}
	// fmt.Println("RESPONSE: ", response)
	return res, nil
}
