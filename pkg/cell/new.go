package cell

// New allows multiple values to be passed along
// and creates a CellInterface based on that
//
// If FromData does not match the type of the cell
// then this will return nil & error as well
func New(
	rowIsHeader bool,
	values []interface{},
) (c CellInterface, err error) {

	c, err = FromData(values[0])
	// if there was no error create the celldata
	// then set the values
	if err == nil {
		c.SetIsRowAHeader(rowIsHeader)
		err = c.Set(values...)
	}

	return
}

func NewColumnHeader(
	values []interface{},
) (c ColumnHeadDataType[ColumnHeadData], err error) {
	c = ColumnHeadDataType[ColumnHeadData]{}
	c.SetIsRowAHeader(true)
	err = c.Set(values...)

	return
}
