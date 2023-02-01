package report2

import "github.com/xuri/excelize/v2"

// CellLocation stores the column
// and row data as well as tracking
// the RowKey value
type CellLocation struct {
	Row    int
	Col    int
	RowKey string
}

func (l *CellLocation) String() (str string) {
	str, _ = excelize.CoordinatesToCellName(l.Col, l.Row, false)
	return
}
