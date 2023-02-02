package report2

import (
	"opg-infra-costs/pkg/debugger"

	"github.com/xuri/excelize/v2"
)

// CellLocation stores the column
// and row coordinates
type CellLocation struct {
	Row int
	Col int
}

func (l *CellLocation) String() (str string) {
	defer debugger.Log("CellLocation.String()", debugger.VVERBOSE)()
	str, _ = excelize.CoordinatesToCellName(l.Col, l.Row, false)
	return
}
