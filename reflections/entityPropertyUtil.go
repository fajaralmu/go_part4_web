package reflections

import (
	"log"
	"reflect"
)

//CreateEntityProperty creates entity field properties
func CreateEntityProperty(clazz reflect.Type, additionalObjectList map[string][]interface{}) EntityProperty {
	// if (clazz == null || getClassAnnotation(clazz, Dto.class) == null) {
	// 	return null
	// }

	// Dto dto = (Dto) getClassAnnotation(clazz, Dto.class)
	// final boolean ignoreBaseField = dto.ignoreBaseField()
	ignoreBaseField := true
	// final boolean isQuestionare = dto.quistionare()
	isQuestionare := false

	var entityProperty EntityProperty = EntityProperty{
		IgnoreBaseField: ignoreBaseField,
		EntityName:      clazz.Name(),
		IsQuestionare:   isQuestionare,
	}

	// try {

	var fieldList []reflect.StructField = getDeclaredFields(clazz)

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

		entityElement := EntityElement{
			Field:           field,
			IgnoreBaseField: entityProperty.IgnoreBaseField,
			entityProperty:  &entityProperty,
			AdditionalMap:   additionalObjectList,
			IsGrouped:       entityProperty.IsQuestionare,
		}
		entityElement.init()

		if false == entityElement.Build() {
			continue
		}
		if entityElement.DetailField {
			fieldToShowDetail = entityElement.ID
		}

		fieldNames = append(fieldNames, entityElement.ID)
		entityElements = append(entityElements, entityElement)
	}

	entityProperty.Alias = clazz.Name() //(dto.value().isEmpty() ? clazz.getSimpleName() : dto.value())
	entityProperty.Editable = true      //(dto.editable())
	entityProperty.setElementJsonList()
	entityProperty.Elements = entityElements
	entityProperty.DetailFieldName = (fieldToShowDetail)
	entityProperty.DateElementsJSON = (ToJSONString(&entityProperty.DateElements))
	entityProperty.FieldNames = (ToJSONString(&fieldNames))
	entityProperty.FieldNameList = (fieldNames)
	entityProperty.FormInputColumn = 1 //dto.formInputColumn().value)
	entityProperty.determineIdField()

	log.Println("============ENTITY PROPERTY: {} ", entityProperty)

	return entityProperty
	// } catch (Exception e) {
	// 	e.printStackTrace()
	// 	throw e
	// }

}
