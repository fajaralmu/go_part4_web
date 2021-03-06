package appConfig

import (
	"log"
	"reflect"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/fajaralmu/go_part4_web/repository"
)

//CreateEntityProperty creates entity field properties
func CreateEntityProperty(modelType reflect.Type, inputColumn int) EntityProperty {
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
		EntityName:      reflections.ToSnakeCase(modelType.Name(), false),
		IsQuestionare:   isQuestionare,
		FormInputColumn: inputColumn,
	}

	// // try {
	// obj := CreateNewTypeNotPointer(clazz)
	// r := reflect.ValueOf(obj)
	// v := reflect.Indirect(r).FieldByName("ID").Elem

	var fieldList []reflect.StructField = reflections.GetDeclaredFields(modelType)
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
				configKey := reflections.GetWordsAfterLastChar(field.Type.String(), ".")
				entityConf := GetEntityConf(reflections.ToSnakeCase(configKey, false))
				if nil == entityConf {
					println(configKey, " is not registered in entity config")
					continue
				}

				newList := reflections.CreateNewType(entityConf.ListType)
				list, _ := repository.Filter(newList, entities.Filter{})
				additionalMapKey := field.Name
				additionalObjectList[additionalMapKey] = list

				log.Println("additionalMapKey: ", additionalMapKey, "list size: ", len(list))
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
	entityProperty.Alias = modelType.Name() //(dto.value().isEmpty() ? clazz.getSimpleName() : dto.value())
	entityProperty.Editable = true          //(dto.editable())
	entityProperty.setElementJSONList()
	entityProperty.Elements = entityElements
	entityProperty.DetailFieldName = (fieldToShowDetail)
	entityProperty.FieldNames = sliceOfStringToJSONString(fieldNames)
	entityProperty.FieldNameList = (fieldNames)
	entityProperty.determineIdField()
	entityProperty.setGridTemplateColumns()
	// log.Println("============ENTITY PROPERTY: {} ", entityProperty)

	return entityProperty

}
