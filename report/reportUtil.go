package report

import (
	"fmt"
	"strings"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/reflections"
)

func getEntitiesTableValues(entities []interface{}, entityProperty appConfig.EntityProperty) []interface{} {
	entityElements := entityProperty.Elements

	var values []interface{}
	seqNum := 0

	/**
	 * column header
	 */
	values = append(values, "No")
	seqNum++
	for _, element := range entityElements {
		values = append(values, element.LableName)
		seqNum++
	}

	/**
	 * table content
	 */
	for e, entity := range entities {

		values = append(values, e+1) // numbering
		seqNum++

		/**
		 * checking the value type
		 */
		for _, element := range entityElements {

			value := mapEntityValue(entity, element)
			values = append(values, value)
			seqNum++

		}
	}

	return values
}

func mapEntityValue(entity interface{}, element appConfig.EntityElement) interface{} {

	var value interface{}

	value, ok := reflections.GetFieldValue(element.ID, entity)
	if !ok {
		return value
	}

	fieldType := element.Type
	if nil != value && "" != value {

		if objectEquals(fieldType, "FIELD_TYPE_DYNAMIC_LIST", "FIELD_TYPE_FIXED_LIST") {

			optionItemName := element.OptionItemName

			if "" != optionItemName {

				// Field converterField = getDeclaredField(field.getType(), optionItemName);
				// Object converterValue = converterField.get(value);
				converterValue, ok := reflections.GetFieldValue(optionItemName, value)
				if ok {
					value = converterValue
				} else {
					value = "[error]"
				}

			} else {
				// value = value.toString();
			}

		} else if objectEquals(fieldType, "FIELD_TYPE_IMAGE") {
			strVal := fmt.Sprintf("%v", value)
			value = strings.Split(strVal, "~")[0]
			//					values[seqNum] = ComponentBuilder.imageLabel(UrlConstants.URL_IMAGE+value, 100, 100);
			//					continue elementLoop;

		} else if objectEquals(fieldType, "FIELD_TYPE_DATE") {

			value = value //DateUtil.formatDate((Date) value, DATE_PATTERN);

		} else if objectEquals(fieldType, "FIELD_TYPE_NUMBER") {

			value = value //Double.parseDouble(value.toString())
		}
	}

	return value

}

func objectEquals(obj interface{}, compares ...interface{}) bool {

	for _, val := range compares {
		if obj == val {
			return true
		}
	}

	return false
}
