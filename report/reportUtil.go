package report

import (
	"fmt"
	"log"
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
	log.Println("values: ", len(values))

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

func getTableContentMap(columCount int, values ...interface{}) map[int][]interface{} {
	log.Println("list size: ", len(values))
	var tableContent map[int][]interface{} = map[int][]interface{}{}
	var rowNum int
	for i, value := range values {

		if tableContent[rowNum] == nil {
			tableContent[rowNum] = []interface{}{}
		}
		tempSlize := tableContent[rowNum]
		tempSlize = append(tempSlize, value)
		tableContent[rowNum] = tempSlize

		if (i+1)%columCount == 0 {
			rowNum++
		}
	}

	log.Println("rowNum: ", rowNum, " columCount: ", columCount)
	log.Println("tableContent: ", tableContent)
	return tableContent
}

func createCell(row excelRow, col int) excelCell {
	return excelCell{
		vIndex: row.index,
		hIndex: col,
	}
}

/**
 * fill row with values
 *
 * @param parentRow
 * @param offsetIndex
 * @param sourceStyle
 * @param values
 */
func fillRows(parentRow excelRow, offsetIndex int, values ...interface{}) (row excelRow) {
	// DataFormat fmt = parentRow.getSheet().getWorkbook().createDataFormat();
	// XSSFCell[] cells = new XSSFCell[values.length];
	var cells []excelCell
	for i, cellValue := range values {
		if cellValue == nil {
			cellValue = ""
		}
		cell := createCell(parentRow, offsetIndex+i)
		cell.value = cellValue
		// CellStyle cellStyle = createCellStyle(parentRow.getSheet().getWorkbook());

		// if (sourceStyle != null) {
		// 	cellStyle.cloneStyleFrom(sourceStyle);
		// 	cell.setCellStyle(cellStyle);
		// }

		cell.value = cellValue
		cells = append(cells, cell)
	}
	parentRow.cells = cells
	return parentRow
}

func createRow(sheetName string, rownum int, offsetIndex int,
	values ...interface{}) (row excelRow) {

	// final XSSFRow existingRow = sheet.getRow(rownum);
	// XSSFRow row = existingRow == null ? sheet.createRow(rownum) : existingRow;
	row = excelRow{index: rownum}
	// XSSFCellStyle style = sheet.getWorkbook().createCellStyle();
	// setAllBorder(style, THIN);
	row = fillRows(row, offsetIndex /*style, */, values...)

	// for (int i = 0; i < values.length; i++) {
	// 	sheet.autoSizeColumn(i);
	// }

	return row
}
