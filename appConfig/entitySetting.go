package appConfig

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/fajaralmu/go_part4_web/reflections"
)

type EntityElement struct {
	ID                  string
	Type                string
	ClassName           string
	Identity            bool
	Required            bool
	IDField             bool
	LableName           string
	Options             []interface{}
	JSONList            string
	OptionItemName      string
	OptionValueName     string
	OptionValueType     string
	EntityReferenceName string
	EtityReferenceClass string
	Multiple            bool
	ShowDetail          bool
	DetailFields        string
	DefaultValues       []string
	PlainListValues     []interface{}

	IsGrouped        bool
	InputGroupname   string
	OptionJSONString string

	DetailField bool

	// not shown in view

	Field           reflect.StructField
	IgnoreBaseField bool
	entityProperty  *EntityProperty
	AdditionalMap   map[string][]interface{}

	FormField map[string]string
	// BaseField baseField
	SkipBaseField bool
	HasJoinColumn bool
}

//private
func (e *EntityElement) init() {
	fieldTag, fieldTagOK := reflections.GetMapOfTag(e.Field, "custom") //field.getAnnotation(FormField.class);

	log.Println("e.Field.Name: ", e.Field.Name, "e.Field.Type: ", e.Field.Type)
	log.Printf("Custom fieldTag [%v] : %v  \n", fieldTagOK, fieldTag)

	/////FOR ID FIELD////
	if e.Field.Name == "ID" {
		fieldTag = map[string]string{
			"type": "FIELD_TYPE_NUMBER",
		}
		fieldTagOK = true
	}

	if fieldTagOK {
		e.FormField = fieldTag
	} else {

		e.FormField = nil
		return
	}
	// baseField = field.getAnnotation(BaseField.class);

	e.IDField = e.Field.Name == "ID"
	e.SkipBaseField = false // true // !idField && (baseField != null && ignoreBaseField);

	e.Identity = e.IDField
	e.HasJoinColumn = e.FormField["foreignKey"] != ""

	e.checkIfGroupedInput()
}

func (e *EntityElement) GetJsonListString(removeBeginningAndEndIndex bool) string {

	jsonBytes, _ := json.Marshal(&e.JSONList)
	jsonStringified := ""
	if removeBeginningAndEndIndex {
		var dummyString string = string(jsonBytes[1:])
		dummyString = string(dummyString[:len(dummyString)-1])
		jsonStringified = dummyString
		log.Println("jsonStringified:", jsonStringified)
	}
	jsonStringified = strings.Replace(jsonStringified, "\\t", "", 0)
	jsonStringified = strings.Replace(jsonStringified, "\\r", "", 0)
	jsonStringified = strings.Replace(jsonStringified, "\\n", "", 0)
	log.Println("RETURN jsonStringified: ", jsonStringified)
	return jsonStringified

}

func (e *EntityElement) checkIfGroupedInput() {

	if e.IsGrouped {
		// AdditionalQuestionField annotation = field.getAnnotation(AdditionalQuestionField.class);
		// inputGroupname = annotation != null ? annotation.value() : AdditionalQuestionField.DEFAULT_GROUP_NAME;
		e.InputGroupname = "DEFAULT_GROUP_NAME"
	}
}

func (e *EntityElement) Build() bool {
	result := e.doBuild()
	e.entityProperty = nil // &EntityProperty{}
	return result
}

func (e *EntityElement) doBuild() bool {

	var formFieldIsNull bool = (e.FormField == nil) // || e.SkipBaseField)

	log.Printf("formFieldIsNullOrSkip: %v %v", formFieldIsNull, e.FormField)

	if formFieldIsNull {
		return false
	}

	var lableName string
	if e.FormField["lableName"] == ("") {
		lableName = e.Field.Name
	} else {
		lableName = strings.Replace(e.FormField["lableName"], "_", " ", -1)
	}
	var determinedFieldType string = e.determineFieldType()

	e.checkFieldType(determinedFieldType)
	var hasJoinColumn bool = e.FormField["foreignKey"] != ""

	if hasJoinColumn {
		e.processJoinColumn(determinedFieldType)
	}

	e.checkDetailField()
	e.ID = (e.Field.Name)
	e.Identity = (e.IDField)
	e.LableName = reflections.ExtractCamelCase(lableName)
	e.Required = e.FormField["required"] == "TRUE"
	e.Type = determinedFieldType
	e.Multiple = e.FormField["multiple"] == "TRUE"
	e.ClassName = e.Field.Type.Name()
	e.ShowDetail = e.FormField["showDetail"] == "TRUE"

	//setting field type so can be read by browser
	switch e.Type {
	case "FIELD_TYPE_TEXT":
		e.Type = "text"
	case "FIELD_TYPE_NUMBER":
		e.Type = "number"
	case "FIELD_TYPE_COLOR":
		e.Type = "color"
	case "FIELD_TYPE_DATE":
		e.Type = "date"
	case "FIELD_TYPE_IMAGE":
		e.Type = "img"
	}

	if e.JSONList != "" {
		e.OptionJSONString = e.GetJsonListString(true)
	}

	return true
}

func (e *EntityElement) checkDetailField() {

	var detailFieldVals []string
	if e.FormField["detailFields"] != "" {
		detailFieldVals = strings.Split(e.FormField["detailFields"], ",")
	}

	if len(detailFieldVals) > 0 {
		e.DetailFields = (strings.Join(detailFieldVals, "~"))
	}
	if e.FormField["showDetail"] == "TRUE" {
		e.OptionItemName = e.FormField["optionItemName"]
		e.DetailField = (true)
	}
}

func (e *EntityElement) checkFieldType(fieldType string) {

	if fieldType == ("FIELD_TYPE_IMAGE") {
		e.processImageType()

	} else if fieldType == ("FIELD_TYPE_CURRENCY") {
		e.processCurrencyType()

	} else if fieldType == ("FIELD_TYPE_DATE") {
		e.processDateType()

	} else if fieldType == ("FIELD_TYPE_PLAIN_LIST") {
		e.processPlainListType()

	} else if fieldType == ("FIELD_TYPE_NUMBER") {
		e.processNumberType()
	}

}

func (e *EntityElement) processNumberType() {
	e.entityProperty.NumberElements = append(e.entityProperty.NumberElements, e.Field.Name)
	// e.EntityProperty.getCurrencyElements().add(field.getName());
}

func (e *EntityElement) processCurrencyType() {
	e.entityProperty.CurrencyElements = append(e.entityProperty.CurrencyElements, e.Field.Name)
	// e.EntityProperty.getCurrencyElements().add(field.getName());
}

func (e *EntityElement) processImageType() {
	e.entityProperty.ImageElements = append(e.entityProperty.ImageElements, e.Field.Name)

	// entityProperty.getImageElements().add(field.getName());
}

func (e *EntityElement) processDateType() {
	e.entityProperty.DateElements = append(e.entityProperty.DateElements, e.Field.Name)

	// entityProperty.getDateElements().add(field.getName());
}

func (e *EntityElement) processPlainListType() {

	var availableValues []string
	if e.FormField["availableValues"] != "" {
		availableValues = strings.Split(e.FormField["availableValues"], ",")
		e.PlainListValues = reflections.ToInterfaceSlice(&availableValues)
		// } else if (e.Field.getType().isEnum()) {
		// 	Object[] enumConstants = field.getType().getEnumConstants();
		// 	setPlainListValues(Arrays.asList(enumConstants));

	} else {
		log.Println("Ivalid PlainListT: ", e.Field.Name)
		// throw new Exception("Invalid PlainListT");
	}
}

func (e *EntityElement) determineFieldType() string {

	var fieldType string

	if reflections.IsNumericType(e.Field.Type) {
		fieldType = "FIELD_TYPE_NUMBER"

	} else if e.Field.Type == reflect.TypeOf(time.Time{}) {
		fieldType = "FIELD_TYPE_DATE"

	} else if e.IDField {
		fieldType = "FIELD_TYPE_HIDDEN"
	} else {
		fieldType = e.FormField["type"]
	}
	return fieldType
}

func (e *EntityElement) processJoinColumn(fieldType string) {
	// log.info("field {} of {} is join column, type: {}", e.Field.Name, fieldType)

	referenceEntityClass := e.Field.Type
	// referenceEntityIdField := "ID" // Get EntityUtil.getIdFieldOfAnObject(referenceEntityClass);

	// if (referenceEntityIdField == null) {
	// 	throw new Exception("ID Field Not Found");
	// }

	foreignKeyType := e.FormField["foreignKeyType"]
	if foreignKeyType == "" {
		foreignKeyType = "text"
	}

	if fieldType == ("FIELD_TYPE_FIXED_LIST") && e.AdditionalMap != nil {

		referenceEntityList := e.AdditionalMap[e.Field.Name]
		if nil == referenceEntityList || len(referenceEntityList) == 0 {
			errorStr := "Invalid object list provided for key: " + e.Field.Name + " in EntityElement.AdditionalMap"
			log.Println(errorStr)
		}
		log.Printf("Additional map with key: %v} . Length: %v", e.Field.Name, len(referenceEntityList))
		if referenceEntityList != nil {
			e.Options = (referenceEntityList)
			jsonListStr, _ := json.Marshal(&referenceEntityList)
			e.JSONList = string(jsonListStr)
		}

	} else if fieldType == ("FIELD_TYPE_DYNAMIC_LIST") {
		e.EtityReferenceClass = referenceEntityClass.Name()
	}

	e.OptionValueName = "ID" //(referenceEntityIdField.getName());
	e.OptionItemName = (e.FormField["optionItemName"])
	e.OptionValueType = foreignKeyType
}

///////////////////ENTITY PROPERTY///////////////////

type EntityProperty struct {
	EntityName           string
	Alias                string
	FieldNames           string
	IDField              string
	FormInputColumn      int
	Editable             bool
	WithDetail           bool
	DetailFieldName      string
	ImageElementsJSON    string
	DateElementsJSON     string
	CurrencyElementsJSON string
	NumberElementsJSON   string
	DateElements         []string
	ImageElements        []string
	CurrencyElements     []string
	NumberElements       []string
	Elements             []EntityElement
	FieldNameList        []string
	IgnoreBaseField      bool
	IsQuestionare        bool
	GroupNames           string
}

func (e *EntityProperty) setElementJsonList() {

	dateJSON, _ := json.Marshal(&e.DateElements)
	e.DateElementsJSON = string(dateJSON)
	imgJSON, _ := json.Marshal(&e.ImageElements)
	e.ImageElementsJSON = string(imgJSON)
	currJSON, _ := json.Marshal(&e.CurrencyElements)
	e.CurrencyElementsJSON = string(currJSON)
	numJSON, _ := json.Marshal(&e.NumberElements)
	e.NumberElementsJSON = string(numJSON)
}

func (e *EntityProperty) removeElements(fieldNames ...string) {
	if e.Elements == nil || len(e.Elements) == 0 {
		return
	}

	for i := 0; i < len(fieldNames); i++ {
		var fieldName string = fieldNames[i]
	loop:
		for j, fName := range e.FieldNameList {
			if fieldName == (fName) {

				e.FieldNameList[j] = e.FieldNameList[len(e.FieldNameList)-1] // Copy last element to index i.
				e.FieldNameList[len(e.FieldNameList)-1] = ""                 // Erase last element (write zero value).
				e.FieldNameList = e.FieldNameList[:len(e.FieldNameList)-1]
				// dele
				// fieldNameList.remove(fName)
				break loop
			}
		}
	loop2:
		for j, entityElement := range e.Elements {
			if entityElement.ID == (fieldName) {
				e.Elements[j] = e.Elements[len(e.Elements)-1]   // Copy last element to index i.
				e.Elements[len(e.Elements)-1] = EntityElement{} // Erase last element (write zero value).
				e.Elements = e.Elements[:len(e.Elements)-1]

				break loop2
			}
		}
	}
	fieldNamesJSON, _ := json.Marshal(&e.FieldNameList)
	e.FieldNames = string(fieldNamesJSON)
}

func (e *EntityProperty) setGroupNames(groupNamesArray []string) {
	// var removedIndex int = 0
	for i, groupNameArr := range groupNamesArray {
		if groupNameArr == "DEFAULT_GROUP_NAME" {
			// removedIndex = i

			groupNamesArray[i] = groupNamesArray[len(groupNamesArray)-1] // Copy last element to index i.
			groupNamesArray[len(groupNamesArray)-1] = ""                 // Erase last element (write zero value).
			groupNamesArray = groupNamesArray[:len(groupNamesArray)-1]
		}
	}
	// groupNamesArray = ArrayUtils.remove(groupNamesArray, removedIndex)
	e.GroupNames = strings.Join(groupNamesArray, ",")
	e.GroupNames = e.GroupNames + "," + "DEFAULT_GROUP_NAME"
}

//	static  main(String[] args) {
//		args =new String[] {"OO", "ff", "fff22"}
//		for (int i = 0 i < args.length i++) {
//			if(args[i] == "OO")
//		}
//	}

func (e *EntityProperty) getGridTemplateColumns() string {
	if e.FormInputColumn == 2 {
		return "20% 70%"
	}
	return strings.Repeat("auto ", e.FormInputColumn)
}

func (e *EntityProperty) determineIdField() {
	if nil == e.Elements {
		log.Println("Entity ELements is NULL")
		return
	}
	for _, entityElement := range e.Elements {
		if entityElement.IDField && e.IDField == "" {
			e.IDField = (entityElement.ID)
		}
	}
}
