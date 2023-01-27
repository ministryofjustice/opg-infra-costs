package report

import (
	"fmt"
	"math"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	name          string
	columns       []Column
	dataset       map[string]map[string][]string
	rowCount      int
	colCount      int
	visible       bool
	cells         map[CellRef]CellInfo
	styles        map[CellRef]*excelize.Style
	tableOptions  *excelize.TableOptions
	filterOptions *excelize.AutoFilterOptions
	groupColumns  []Column
	dateColumns   []Column
	otherColumns  []Column
	hideRowWhen   map[CellRef]float64
}

var defaultStyle = &excelize.Style{
	NumFmt: 177,
}
var defaultTableOptions = &excelize.TableOptions{
	StyleName: "TableStyleMedium9",
}
var defaultFilterOptions = &excelize.AutoFilterOptions{}

// SetName overwrites the .name property with n
func (s *Sheet) SetName(n string) (err error) {
	s.name = n
	return
}
func (s *Sheet) GetName() string {
	return s.name
}

// SetColumns adds the slice passed on the existing columns slice.
// This is therefore additive and can be called multiple times
func (s *Sheet) SetColumns(columns []Column, ty ColumnDataType) (err error) {
	s.columns = append(s.columns, columns...)

	if ty == ColumnsAreGroupBy {
		s.groupColumns = columns
	} else if ty == ColumnsAreDateCost {
		s.dateColumns = columns
	} else {
		s.otherColumns = append(s.otherColumns, columns...)
	}
	return
}

// GetGroupColumns gets only the column mapkeys for those columns marked
// as ColumnsAreGroupBy.
// These columns are generally used to group data into a row structure
func (s *Sheet) GetGroupColumns() (fieldNames []string) {
	for _, c := range s.groupColumns {
		fieldNames = append(fieldNames, c.MapKey)
	}
	return
}

// GetGroupColumns gets only the column mapkeys for those columns marked
// as ColumnsAreOther.
// These columns are typically the formulas and additional data at the end
func (s *Sheet) GetOtherColumns() (fieldNames []string) {
	for _, c := range s.otherColumns {
		fieldNames = append(fieldNames, c.MapKey)
	}
	return
}

// GetTransposeColumns returns a map of source & value for these columns.
// This is how we transpose column data into row data
// map key = Column heading, map value = Value of cell
func (s *Sheet) GetTransposeColumns() map[string]string {
	return map[string]string{"Date": "Cost"}
}

// SetDataset overwrites the .dataset property & then generates the
// .cells list
func (s *Sheet) SetDataset(ds map[string]map[string][]string) (err error) {
	s.dataset = ds
	s.headers()
	s.rows()
	return
}

// Write generates all the content for the sheet
// - main func to generate data for the sheet
func (s *Sheet) Write(f *excelize.File) (i int, err error) {
	i, err = f.NewSheet(s.name)
	for _, cell := range s.cells {
		err = CellWriter(cell, s.name, f)
	}
	return
}

// Cell will retrieve the value and type information for cell
// at the row|col passed
func (s *Sheet) Cell(row int, col int) (c CellInfo, ok bool) {
	ok = false
	for _, cell := range s.cells {
		if cell.Ref.Row == row && cell.Ref.Col == col {
			c = cell
			ok = true
		}
	}
	return
}

func (s *Sheet) SetVisible(visible bool) (err error) {
	s.visible = visible
	return
}
func (s *Sheet) GetVisible() bool {
	return s.visible
}

// AddStyle provides a way to overwrite the defauly cell style
// for any content based on the row/col location
// By setting just row or col you can set the style for that entire section
// - https://xuri.me/excelize/en/style.html#number_format
func (s *Sheet) AddStyle(st *excelize.Style, row int, col int) {
	var ref CellRef
	if row > 0 && col > 0 {
		ref = CellRef{Row: row, Col: col}
		s.styles[ref] = st
	} else if col > 0 {
		ref = CellRef{Col: col}
		s.styles[ref] = st
	} else if row > 0 {
		ref = CellRef{Row: row}
		s.styles[ref] = st
	}

}

// AddPane add panes to allow scrolling on long data sets
func (s *Sheet) AddPane(f *excelize.File) (err error) {
	row := 1
	col := len(s.groupColumns)
	err = f.SetPanes(s.name, &excelize.Panes{
		Freeze:      true,
		XSplit:      col,
		YSplit:      row,
		TopLeftCell: "A1",
		ActivePane:  "bottomRight",
		Panes: []excelize.PaneOptions{
			{Pane: "topLeft"},
			{Pane: "topRight"},
			{Pane: "bottomLeft"},
			{Pane: "bottomRight", ActiveCell: "A1", SQRef: "A1"},
		},
	})
	return
}

// AddTable creates a table and adds autofilters
func (s *Sheet) AddTable(f *excelize.File) (err error) {
	rangeRef := fmt.Sprintf(
		"%s:%s",
		CellReference(1, 1),
		CellReference(s.rowCount, s.colCount-1),
	)
	err = f.AddTable(s.name, rangeRef, s.tableOptions)
	f.AutoFilter(s.name, rangeRef, s.filterOptions)
	return
}

// SetFilterOptions allows changing of the options passed when creating the autofilter
// for the table
func (s *Sheet) SetFilterOptions(options *excelize.AutoFilterOptions) (err error) {
	s.filterOptions = options
	return
}

// SetTableOptions allows configuration of the table before setting
func (s *Sheet) SetTableOptions(options *excelize.TableOptions) (err error) {
	s.tableOptions = options
	return
}

// SetHideRowWhen allows configuration of how to hide certain rows
func (s *Sheet) SetHideRowWhen(criteria map[CellRef]float64) (err error) {
	s.hideRowWhen = criteria
	return
}

// RowVisibility uses hideRowWhen to hide certain rows
// - all critera has to pass for the row to be visible
func (s *Sheet) RowVisibility(f *excelize.File) (hidden []int, err error) {
	if len(s.hideRowWhen) > 0 {
		for i := 2; i <= s.rowCount; i++ {
			showRow := true

			// check criteria
			for ref, crit := range s.hideRowWhen {
				var cellVal string
				cell := CellReference(i, ref.Col)
				if cellVal, _ = f.GetCellValue(s.name, cell); len(cellVal) == 0 {
					cellVal, _ = f.CalcCellValue(s.name, cell)
				}
				val, pErr := strconv.ParseFloat(cellVal, 64)
				if pErr != nil {
					val = 0.0
				}
				val = math.Abs(val)
				greater := (val >= crit)
				if !greater {
					showRow = false
				}
				//fmt.Printf("[%s]\t (%.2f >= %v)\t(%v)\n", cell, val, crit, greater)
			}

			if !showRow {
				hidden = append(hidden, i)
				f.SetRowVisible(s.name, i, false)
			}
		}

	}
	return
}

// Init runs all the default setup for the sheet
func (s *Sheet) Init() {
	s.cells = make(map[CellRef]CellInfo)
	s.styles = make(map[CellRef]*excelize.Style)
	s.tableOptions = defaultTableOptions
	s.filterOptions = defaultFilterOptions
	s.SetVisible(true)
}

// === internal

// headers generates the cell data for the current .columns
func (s *Sheet) headers() {
	s.rowCount = 1
	s.colCount = 1
	// this writes the headers
	for _, col := range s.columns {
		ref := CellRef{Row: s.rowCount, Col: s.colCount, RowKey: "Header"}
		s.cells[ref] = CellInfo{Value: col.Display, Ref: ref}
		s.colCount++
	}
}

// rows iterates over the dataset, then loops over the columns
// get the value and writes that to the `.cells`
// -- will parse formula values as well
func (s *Sheet) rows() {

	for rowKey, row := range s.dataset {
		s.rowCount++
		s.colCount = 1

		formulaReplacements := map[string]interface{}{
			"r": strconv.Itoa(s.rowCount),
		}
		// now loop over the columns and fetch that data from the row
		for _, col := range s.columns {
			ref := CellRef{Row: s.rowCount, Col: s.colCount, RowKey: rowKey}
			style := s.style(s.rowCount, s.colCount)

			// get the set of values
			var values []string = []string{"0.0"}
			if len(col.Formula) > 0 {
				values = []string{col.Formula}
			} else if v, ok := row[col.MapKey]; ok && len(v) > 0 {
				values = v
			}

			valueType, _ := DataType(values[0])
			value, _ := CellValueFromType(values, valueType)
			if valueType == DataIsAFormula {
				value, _ = ParseFormula(value.(string), formulaReplacements)
			}

			s.cells[ref] = CellInfo{
				Value:     value,
				ValueType: valueType,
				Style:     style,
				Ref:       ref,
			}
			s.colCount++
		}

	}
}

// style gets the style for the row/col passed
func (s *Sheet) style(row int, col int) *excelize.Style {
	exact := CellRef{Row: row, Col: col}
	colMatch := CellRef{Col: col}
	rowMatch := CellRef{Row: row}

	if st, ok := s.styles[exact]; ok {
		return st
	} else if st, ok := s.styles[colMatch]; ok {
		return st
	} else if st, ok := s.styles[rowMatch]; ok {
		return st
	}

	return defaultStyle

}

// == NEW SHEET
func NewSheet(name string) Sheet {
	s := Sheet{}
	s.Init()
	s.SetName(name)
	return s
}
