package cell

// New allows multiple values to be passed along
// and creates a CellData based on that
//
// If FromData does not match the type of the cell
// then this will return nil & error as well
func New(values []interface{}) (c CellData, err error) {
	c, err = FromData(values[0])
	c.Set(values...)
	return
}
