package report

import (
	"log"
	"strconv"

	"github.com/fajaralmu/go_part4_web/reflections"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fajaralmu/go_part4_web/appConfig"
)

var upperCased []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//GetCellName returns column name
func GetCellName(col, row int) string {

	arrLen := len(upperCased)
	name := ""
	if col < arrLen {
		name = upperCased[col]
	} else {
		idx := 0
		idxCounter := 0
		for i := 0; i < col-arrLen; i++ {
			if i > 0 && i%arrLen == 0 {
				idx++
				idxCounter = 0
			}
			idxCounter++
		}
		name = upperCased[idx] + upperCased[idxCounter]
	}

	return name + strconv.Itoa(row)

}

//GetEntityReport generate excel report for list of models
func GetEntityReport(entities []interface{}, entityProp appConfig.EntityProperty) {
	f := excelize.NewFile()
	// Create a new sheet.

	// Set value of a cell.
	// f.SetCellValue(sheetName, "A2", "Hello world.")
	// Set active sheet of the workbook.
	writeCellValues(f, entities, entityProp)
	// Save xlsx file by the given path.
	random := reflections.RandomNum(30)
	if err := f.SaveAs("./reports/" + entityProp.EntityName + random + ".xlsx"); err != nil {
		println(err.Error())
	}
}

func writeCellValues(file *excelize.File, entities []interface{}, entityProperty appConfig.EntityProperty) {

	sheetName := entityProperty.EntityName
	index := file.NewSheet(sheetName)
	entityValues := getEntitiesTableValues(entities, entityProperty)
	log.Println("entityValues: ", len(entityValues))

	rows := createTable(sheetName, len(entityProperty.Elements)+1, 2, 2, entityValues...)
	log.Println("rows: ", len(rows))

	for _, row := range rows {
		cells := row.cells
		for _, col := range cells {
			cellName := GetCellName(col.hIndex, row.index)
			file.SetCellValue(sheetName, cellName, col.value)
		}
	}

	file.SetActiveSheet(index)
}

func createTable(sheetName string, columCount int, xOffset int, yOffset int, values ...interface{}) []excelRow {
	tableContent := getTableContentMap(columCount, values...)
	var rows []excelRow
	for key, value := range tableContent {
		log.Printf("row: %v of %v \n", key, len(tableContent))

		row := createRow(sheetName, key+yOffset, xOffset, value...)
		// autosizeColumn(row, rowValues.length, BorderStyle.THIN, null);
		rows = append(rows, row)
	}

	return rows
}
