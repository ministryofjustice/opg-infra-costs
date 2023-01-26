package report

import "github.com/xuri/excelize/v2"

// -- Date handling
const DATEFORMAT string = "2006-01"

// -- Sheet main interface
type SheetInterface interface {
	SetName(name string) error
	GetName() string

	SetColumns(cols []Column, ty ColumnDataType) error
	GetGroupColumns() []string
	GetDateCostColumns() map[string]string
	GetOtherColumns() []string

	SetDataset(ds map[string]map[string][]string) error

	SetVisible(v bool) error
	GetVisible() bool
	AddStyle(st *excelize.Style, row int, col int)
	AddTable(f *excelize.File) error
	AddPane(f *excelize.File) (err error)

	Write(f excelize.File) error
	Cell(ref string) (CellInfo, bool)
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

// -- CELLS
type CellRef struct {
	Row int
	Col int
}
type CellInfo struct {
	Value     interface{}
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
