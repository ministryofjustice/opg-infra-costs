package row

import "opg-infra-costs/pkg/cell"

// GetCells simply returns the cells internal
// slice of CellData interfaces structs
func (r *Row) GetCells() []cell.CellData {
	return r.cells
}

// SetCells takes a slice of mixed (likely strings)
// covnerts that to a cell and adds to the internal
// CellData slice
func (r *Row) SetCells(dataset [][]interface{}) (err error) {
	var c cell.CellData

	var isThisRowAHeader bool = false
	if r.GetHeader() {
		isThisRowAHeader = true
	}

	for _, values := range dataset {

		c, err = cell.New(isThisRowAHeader, values)

		if err == nil {
			r.cells = append(r.cells, c)
		}

	}
	return
}
