package report

import "github.com/xuri/excelize/v2"

// -- Date handling
const DATEFORMAT string = "2006-01"

// -- Sheet main interface
type SheetInterface interface {
	SetName(name string) error
	GetName() string

	SetVisible(v bool) error
	GetVisible() bool

	SetColumns(cols []Column, ty ColumnDataType) error
	GetGroupColumns() []string
	GetTransposeColumns() map[string]string
	GetOtherColumns() []string

	AddStyle(st *excelize.Style, row int, col int)
	SetFilterOptions(options *excelize.AutoFilterOptions) error
	SetTableOptions(options *excelize.TableOptions) error
	SetHideRowWhen(criteria map[CellRef]float64) (err error)

	SetDataset(ds map[string]map[string][]string) (mapped map[string]RowKeyIndexSet, err error)
	AddRow(key string, cells map[string][]string) (mapped map[string]RowKeyIndexSet, err error)
	AddCell(row int, col int, value string) error

	Write(f excelize.File) (int, error)
	AddTable(f *excelize.File) error
	AddPane(f *excelize.File) error
	RowVisibility(f *excelize.File) (hidden []int, err error)

	Cell(row int, col int) (c CellInfo, ok bool)
}

// -- COLUMNS
type Column struct {
	MapKey  string
	Display string
	Formula string
}
type ColumnDataType string

const (
	ColumnsAreGroupBy  ColumnDataType = "group-by"
	ColumnsAreDateCost ColumnDataType = "date-cost"
	ColumnsAreOther    ColumnDataType = "other"
)

// -- ROWS
type RowKeyIndexSet struct {
	Index   int
	Columns map[string]int
}

// -- CELLS
type CellRef struct {
	Row    int
	Col    int
	RowKey string
}
type CellInfo struct {
	Value     interface{}
	Ref       CellRef
	ValueType ValueDataType
	Style     *excelize.Style
}

type ValueDataType string

const (
	DataIsADate    ValueDataType = "is-a-date"
	DataIsAFormula ValueDataType = "is-a-formula"
	DataIsANumber  ValueDataType = "is-a-number"
	DataIsAString  ValueDataType = "is-a-string"
)
