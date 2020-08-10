package app

import (
	"log"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/report"

	"github.com/fajaralmu/go_part4_web/entities"
)

func getEntitiesReport(request entities.WebRequest) []byte {
	log.Println("generate Entity Report")
	//		request.getFilter().setLimit(0);

	filtered, _ := filterEntity(request)
	entityProp := getEntityProperty(request)

	list := reflections.ConvertInterfaceToSlice(filtered.ResultList)
	file := report.GetEntityReport(list, entityProp)
	if nil != file {
		buffer, err := file.WriteToBuffer()
		if err == nil {

			return buffer.Bytes()
		}
	}
	// return file;
	return nil
}

func getEntityProperty(request entities.WebRequest) appConfig.EntityProperty {
	entityName := request.Filter.EntityName
	entitytConf := appConfig.GetEntityConf(entityName)
	return appConfig.CreateEntityProperty(entitytConf.SingleType, entitytConf.FormInputColumn)
}
