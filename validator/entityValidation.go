package validator

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/fajaralmu/go_part4_web/dataaccess"
	"github.com/fajaralmu/go_part4_web/files"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
)

//ValidateEntity validates entity field before Persisting to DB
func ValidateEntity(model entities.InterfaceEntity, currentRecord entities.InterfaceEntity) bool {
	//println("***ValidateEntity***", reflect.TypeOf(model))

	structFields := reflections.GetDeclaredFields(reflect.TypeOf(reflections.Dereference(model).Interface()))
	fmt.Println("GetDeclaredFields size: ", len(structFields))
	valid := true
loop:
	for _, field := range structFields {

		customTag, ok := reflections.GetMapOfTag(field, "custom")

		if !ok {
			//println("NO Custom Tag")
			continue
		}
		var customTagResult bool
		if currentRecord != nil && reflect.ValueOf(currentRecord).IsValid() && reflect.ValueOf(currentRecord).IsZero() == false {
			currentFieldRecord, _ := reflections.GetFieldValue(field.Name, currentRecord)
			customTagResult = processCustomTag(customTag, field, model, currentFieldRecord)
		} else {
			customTagResult = processCustomTag(customTag, field, model, nil)
		}
		if !customTagResult {
			valid = false
			break loop
		}
	}

	//println("_________ END VALIDATION ___________")

	return valid
}

func processCustomTag(customTag map[string]string, field reflect.StructField, model entities.InterfaceEntity, currentFieldRecord interface{}) bool {
	foreignKey := customTag["foreignKey"]
	foreignKeyOk := processForeignKey(foreignKey, field, model)

	fieldType := customTag["type"]
	fieldOK := processFieldValue(fieldType, field, model, currentFieldRecord)

	return foreignKeyOk && fieldOK
}

func processFieldValue(fieldType string, field reflect.StructField, model entities.InterfaceEntity, currentFieldRecord interface{}) bool {
	fieldValue, _ := reflections.GetFieldValue(field.Name, model)

	switch fieldType {
	case "FIELD_TYPE_IMAGE":

		fieldTags, _ := reflections.GetMapOfTag(field, "custom")
		multipleImg := fieldTags["multiple"] == "true"

		if fieldValue != nil {
			typeName := strings.Replace(reflect.TypeOf(model).String(), ".", "", -1)
			typeName = strings.Replace(typeName, "*", "", -1)
			fieldValue = processImg(fieldValue.(string), typeName, multipleImg, currentFieldRecord)
			reflections.SetFieldValue(field.Name, fieldValue, model)

		} else {
			log.Println("processing FIELD_TYPE_IMAGE but IMG base64data is Empty")
		}

	}

	return true
}

func processImg(imgData string, code string, multipleImg bool, currentFieldRecord interface{}) string {
	if (imgData == "") && currentFieldRecord != "" && currentFieldRecord != nil {
		log.Println("imgData is BLANK ... returns currentFieldRecord: ")
		return currentFieldRecord.(string)
	}
	if multipleImg {
		return processMultipleImageData(imgData, code)
	}
	return files.WriteBase64Img(imgData, code)
}

const originalPreffix = "{ORIGINAL>>"

func processMultipleImageData(imageData string, code string) string {
	if imageData == "NULL" {
		return ""
	}
	//log.Println("processMultipleImageData code: ", code)
	base64Images := strings.Split(imageData, "~")
	finalImgURL := ""
	if base64Images != nil && len(base64Images) > 0 {

		imageUrls := []string{}
		for i, base64Image := range base64Images {
			reflections.RandomCounter++
			if base64Image == ("") {
				continue
			}
			var needWriting bool = true
			var imageName string
			if strings.HasPrefix(base64Image, originalPreffix) {

				raw := strings.Split(base64Image, "}")

				if len(raw) > 1 && raw[1] != "" {
					base64Image = raw[1]

				} else {
					imageName = strings.Replace(raw[0], originalPreffix, "", -1)
					needWriting = false
				}
			}
			if needWriting {
				imageName = files.WriteBase64Img(base64Image, code+"_"+strconv.Itoa(i))
			}

			if "" != imageName {
				//log.Println("append imageName: ", imageName)
				imageUrls = append(imageUrls, imageName)
			}
		}
		finalImgURL = strings.Join(imageUrls, "~")

	}

	return finalImgURL

}

func processForeignKey(foreignKey string, field reflect.StructField, model entities.InterfaceEntity) bool {
	//println("begin process foreign key: ", foreignKey)
	if "" == foreignKey {
		return true
	}
	entity := reflections.StructFieldToEntity(field, model)
	entityID := reflections.GetIDValue(entity)

	setUintValue(foreignKey, entityID.(uint), model)

	result, ok := dataaccess.FindByID(entity, entityID)
	fmt.Println("result FIND BY ID: ", result)
	//println("end process foreign key")

	return ok
}

func setUintValue(fieldName string, value uint, model entities.InterfaceEntity) {
	uintVal := uint64(value)
	obj := reflect.Indirect(reflect.ValueOf(model))
	obj.FieldByName(fieldName).SetUint(uintVal)
}

//RemoveID removes `ID` value
func RemoveID(model entities.InterfaceEntity) {
	setUintValue("ID", 0, model)
}

//SetModelID set ID field of the model
func SetModelID(model entities.InterfaceEntity, id uint) {
	setUintValue("ID", id, model)
}
