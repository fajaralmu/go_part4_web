package report

import "github.com/360EntSecGroup-Skylar/excelize"

type excelCell struct {
	hIndex int
	vIndex int
}

type customCell interface {
	getValue() interface{}
	getCell() excelCell
}

func test() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.

	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
}

type currencyCell struct {
	value int64
	cell  excelCell
}

func (c *currencyCell) getCell() excelCell {
	return c.cell
}
func (c *currencyCell) getValue() interface{} {
	return c.value
}

type dateCell struct {
	value int64
	cell  excelCell
}

func (c *dateCell) getCell() excelCell {
	return c.cell
}
func (c *dateCell) getValue() interface{} {
	return c.value
}

type numericCell struct {
	value int64
	cell  excelCell
}

func (c *numericCell) getCell() excelCell {
	return c.cell
}
func (c *numericCell) getValue() interface{} {
	return c.value
}
