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

	SetDataset(ds map[string]map[string][]string) error
	AddStyle(st *excelize.Style, row int, col int)

	SetFilterOptions(options *excelize.AutoFilterOptions) error
	SetTableOptions(options *excelize.TableOptions) error

	Write(f excelize.File) (int, error)

	AddTable(f *excelize.File) error
	AddPane(f *excelize.File) error

	SetHideRowWhen(criteria map[CellRef]float64) (err error)
	RowVisibility(f *excelize.File) (hidden []int, err error)

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
