package row

import "opg-infra-costs/pkg/cell"

type RowInterface interface {
	SetIndex(i int) error
	GetIndex() (int, error)

	SetHeader(b bool) error
	GetHeader() (bool, error)

	SetVisible(v bool) error
	GetVisible() (bool, error)

	// SetCells expects data to be like:
	// {
	//		{"0.01"}, {"1.011"}
	// }
	// for cost related data
	// or:
	// {
	//		{"AWS S3"}, {"AWS S3"}
	// }
	// for string columns
	SetCells(data [][]interface{}) error
	SetDefinedCells(data []cell.CellInterface) error
	GetCells() ([]cell.CellInterface, error)
}

type Row struct {
	index   int
	visible bool
	header  bool
	cells   []cell.CellInterface
}
