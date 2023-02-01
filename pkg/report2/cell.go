package report2

type Cell struct {
	DataType CellDataType
	Location CellLocation
}

func (c *Cell) Value() interface{} {
	return c.DataType.Value()
}

// -- Create a new cell
func NewCell(
	location CellLocation,
	values []interface{},
) (cell Cell, err error) {

	cell.Location = location
	cell.DataType, err = CellDataTypeFromValues(values)

	return
}
