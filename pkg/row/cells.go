package row

import "opg-infra-costs/pkg/cell"

// GetCells simply returns the cells internal
// slice of CellInterface interfaces structs
func (r *Row) GetCells() ([]cell.CellInterface, error) {
	return r.cells, nil
}

// SetCells takes a slice of mixed (likely strings)
// covnerts that to a cell and adds to the internal
// CellInterface slice
func (r *Row) SetCells(dataset [][]interface{}) (err error) {
	var c cell.CellInterface

	var isThisRowAHeader bool = false
	if v, err := r.GetHeader(); err == nil {
		isThisRowAHeader = v
	}

	for _, values := range dataset {

		c, err = cell.New(isThisRowAHeader, values)

		if err == nil {
			r.cells = append(r.cells, c)
		}

	}
	return
}
