package report2

import (
	"opg-infra-costs/pkg/debugger"

	"github.com/xuri/excelize/v2"
)

// Location stores the column
// and row coordinates
type Location struct {
	Row   int
	Col   int
	Sheet string
}

func (l *Location) String() (str string) {
	defer debugger.Log("Location.String()", debugger.VVERBOSE)()
	str, _ = excelize.CoordinatesToCellName(l.Col, l.Row, false)
	return
}
