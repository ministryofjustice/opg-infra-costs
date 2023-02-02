package report2

import "opg-infra-costs/pkg/debugger"

type Cell struct {
	DataType CellDataType
	Location Location
	Column   Column
}

func (c *Cell) Value() interface{} {
	defer debugger.Log("Cell.Value()", debugger.VVERBOSE)()
	return c.DataType.Value(c.Location)
}

// -- Create a new cell
func NewCell(
	location Location,
	values []interface{},
	col Column,
) (cell Cell, err error) {
	defer debugger.Log("NewCell()", debugger.VVERBOSE)()
	cell.Location = location
	cell.Column = col
	// if the column is a formula, change the values
	if col.Definition.IsFormula() {
		values = []interface{}{col.Definition.Formula}
	}
	cell.DataType, err = CellDataTypeFromValues(values)

	return
}
