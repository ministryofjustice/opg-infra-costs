package report

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func (c *CellInfo) CalculatedValue(f *excelize.File, sheet string) (string, error) {
	return f.CalcCellValue(sheet, c.Ref.String())
}

func (r *CellRef) String() string {
	return CellReference(r.Row, r.Col)
}

// CellReference converts a row & col int into a excel colname and row
// -- 1,1 => A1
// -- 2,3 => C2
func CellReference(row int, col int) string {
	c, _ := excelize.ColumnNumberToName(col)
	return fmt.Sprintf("%s%d", c, row)
}

// CellValueFromType uses the value set passed and the guessed data type
// to return the value we want
// - for float/int it should be the sum, everything else is the values[0]
func CellValueFromType(values []string, dt ValueDataType) (value interface{}, err error) {

	if len(values) <= 0 {
		err = fmt.Errorf("values is empty")
		return
	}

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

// CellWriter takes a cell (CellInfo) and write to the sheet & file
// passed
//
// Number & Formula (DataIsANumber/DataIsAFormula) have additional
// handling to set cell style or cell formulas
func CellWriter(
	cell CellInfo,
	sheet string,
	f *excelize.File,
) (err error) {
	refStr := cell.Ref.String()
	cellStyle, _ := f.NewStyle(cell.Style)

	if cell.ValueType == DataIsANumber {
		f.SetCellStyle(sheet, refStr, refStr, cellStyle)
		err = f.SetCellValue(sheet, refStr, cell.Value)
	} else if cell.ValueType == DataIsAFormula {
		f.SetCellStyle(sheet, refStr, refStr, cellStyle)
		f.SetCellFormula(sheet, refStr, cell.Value.(string)[1:], excelize.FormulaOpts{})
	} else {
		err = f.SetCellValue(sheet, refStr, cell.Value)
	}
	return
}
