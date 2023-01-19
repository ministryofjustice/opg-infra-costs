package row

import "opg-infra-costs/pkg/cell"

type RowInterface interface {
	GetIndex() int
	SetIndex(i int)

	GetHeader() bool
	SetHeader(b bool)

	GetVisible() bool
	SetVisible(v bool)

	GetCells() []cell.CellInterface
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
	SetCells(data [][]interface{})
}

type Row struct {
	index   int
	visible bool
	header  bool
	cells   []cell.CellInterface
}
