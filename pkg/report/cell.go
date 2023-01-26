package report

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func CellReference(row int, col int) string {
	c, _ := excelize.ColumnNumberToName(col)
	return fmt.Sprintf("%s%d", c, row)
}

func CellValueFromType(values []string, dt ValueDataType) (value interface{}, err error) {
	if dt == DataIsANumber {
		sum := 0.0
		for _, str := range values {
			v, _ := strconv.ParseFloat(str, 64)
			sum += v
		}
		value = sum
	} else {
		value = values[0]
	}

	return

}

func CellWrite(
	ref string,
	values []string,
	sheet string,
	rowCount int,
	f *excelize.File,
	style *excelize.Style) (val interface{}, dt ValueDataType, s *excelize.Style, err error) {

	if len(values) <= 0 {
		err = fmt.Errorf("values is empty, nothing to write")
		return
	}
	s = style
	dt, _ = DataType(values[0])
	val, _ = CellValueFromType(values, dt)
	cellstyle, _ := f.NewStyle(style)
	// if its a number, we set the cell style first
	if dt == DataIsANumber {
		f.SetCellStyle(sheet, ref, ref, cellstyle)
		err = f.SetCellValue(sheet, ref, val)
	} else if dt == DataIsAFormula {
		// if its a formula, create the set of replacement / sub data and
		// update the value
		replacements := map[string]interface{}{
			"r": strconv.Itoa(rowCount),
		}
		val, err = ParseFormula(val.(string), replacements)
		f.SetCellStyle(sheet, ref, ref, cellstyle)
		f.SetCellFormula(sheet, ref, val.(string)[1:], excelize.FormulaOpts{})

	} else {
		err = f.SetCellValue(sheet, ref, val)
	}

	return
}
