package report

import "github.com/xuri/excelize/v2"

// -- Date handling
const DATEFORMAT string = "2006-01"

// -- Data struct for csv read and converted data
type RawDataset map[string]RawRow
type RawRow map[string][]string

// -- Sheet main interface
type SheetInterface interface {
	SetName(name string) error
	SetColumns(cols []Column) error
	SetDataset(ds RawDataset) error

	AddStyle(st *excelize.Style, row int, col int)
	AddTable(f *excelize.File, rangeRef string) error
	AddPane(f *excelize.File, row int, col int) (err error)

	Write(f excelize.File) error
	Cell(ref string) (CellInfo, bool)
}

// -- COLUMNS
type Column struct {
	MapKey  string
	Display string
	Formula string
}

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
