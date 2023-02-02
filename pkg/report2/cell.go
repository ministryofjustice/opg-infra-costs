package report2

import "opg-infra-costs/pkg/debugger"

type Cell struct {
	DataType CellDataType
	Location Location
}

func (c *Cell) Value() interface{} {
	defer debugger.Log("Cell.Value()", debugger.VVERBOSE)()
	return c.DataType.Value(c.Location)
}

// -- Create a new cell
func NewCell(
	location Location,
	values []interface{},
) (cell Cell, err error) {
	defer debugger.Log("NewCell()", debugger.VVERBOSE)()
	cell.Location = location
	cell.DataType, err = CellDataTypeFromValues(values)

	return
}
