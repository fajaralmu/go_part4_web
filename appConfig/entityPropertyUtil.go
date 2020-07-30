package appConfig

import (
	"log"
	"reflect"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/repository"
)

//CreateEntityProperty creates entity field properties
func CreateEntityProperty(clazz reflect.Type) EntityProperty {
	additionalObjectList := map[string][]interface{}{}
	log.Println("~~~~~~~~~~CreateEntityProperty~~~~~~~~~~~")
	// if (clazz == null || getClassAnnotation(clazz, Dto.class) == null) {
	// 	return null
	// }

	// Dto dto = (Dto) getClassAnnotation(clazz, Dto.class)
	// final boolean ignoreBaseField = dto.ignoreBaseField()
	ignoreBaseField := true
	// final boolean isQuestionare = dto.quistionare()
	isQuestionare := false

	var entityProperty EntityProperty = EntityProperty{
		IDField:         "ID",
		IgnoreBaseField: ignoreBaseField,
		EntityName:      clazz.Name(),
		IsQuestionare:   isQuestionare,
	}

	// // try {
	// obj := CreateNewTypeNotPointer(clazz)
	// r := reflect.ValueOf(obj)
	// v := reflect.Indirect(r).FieldByName("ID").Elem

	var fieldList []reflect.StructField = reflections.GetDeclaredFields(clazz)
	log.Printf("field LIST size: %v \n", len(fieldList))
	// if (isQuestionare) {
	// 	Map<String, List<Field>> groupedFields = sortListByQuestionareSection(fieldList)
	// 	fieldList = CollectionUtil.mapOfListToList(groupedFields)
	// 	Set<String> groupKeys = groupedFields.keySet()
	// 	String[] keyNames = CollectionUtil.toArrayOfString(groupKeys.toArray())

	// 	entityProperty.setGroupNames(keyNames)
	// }
	entityElements := []EntityElement{}
	fieldNames := []string{}
	fieldToShowDetail := ""

	for _, field := range fieldList {

		customTag, ok := reflections.GetMapOfTag(field, "custom")

		if ok {
			if customTag["foreignKey"] != "" {
				entityConf := GetEntityConf(reflections.ToSnakeCase(field.Name, false))
				if nil == entityConf {
					println(field.Name, " is not registered in entity cinfig")
					continue
				}

				newList := reflections.CreateNewType(entityConf.ListType)
				list, _ := repository.Filter(newList, entities.Filter{})
				additionalObjectList[entityConf.Name] = list
			}
		}

		entityElement := EntityElement{
			Field:           field,
			IgnoreBaseField: entityProperty.IgnoreBaseField,
			entityProperty:  &entityProperty,
			AdditionalMap:   additionalObjectList,
			IsGrouped:       entityProperty.IsQuestionare,
		}
		entityElement.init()
		elementBuilt := entityElement.Build()

		log.Printf("elementBuilt [%v]: %v \n", field.Name, elementBuilt)

		if false == elementBuilt {
			continue
		}
		if entityElement.DetailField {
			fieldToShowDetail = entityElement.ID
		}

		fieldNames = append(fieldNames, entityElement.ID)
		entityElements = append(entityElements, entityElement)
	}

	log.Printf("entityElements size: %v \n", len(entityElements))
	log.Printf("fieldNames: %v ", fieldNames)
	entityProperty.Alias = clazz.Name() //(dto.value().isEmpty() ? clazz.getSimpleName() : dto.value())
	entityProperty.Editable = true      //(dto.editable())
	entityProperty.setElementJsonList()
	entityProperty.Elements = entityElements
	entityProperty.DetailFieldName = (fieldToShowDetail)
	entityProperty.DateElementsJSON = (reflections.ToJSONString(&entityProperty.DateElements))
	entityProperty.FieldNames = (reflections.ToJSONString(&fieldNames))
	entityProperty.FieldNameList = (fieldNames)
	entityProperty.FormInputColumn = 1 //dto.formInputColumn().value)
	entityProperty.determineIdField()

	// log.Println("============ENTITY PROPERTY: {} ", entityProperty)

	return entityProperty
	// } catch (Exception e) {
	// 	e.printStackTrace()
	// 	throw e
	// }

}
