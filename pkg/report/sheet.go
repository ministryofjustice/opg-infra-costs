package report

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	name         string
	columns      []Column
	dataset      map[string]map[string][]string
	rowCount     int
	colCount     int
	visible      bool
	cells        map[CellRef]CellInfo
	styles       map[CellRef]*excelize.Style
	groupColumns []Column
	dateColumns  []Column
	otherColumns []Column
}

var defaultStyle = &excelize.Style{
	NumFmt: 177,
}

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

func (s *Sheet) GetGroupColumns() (fieldNames []string) {
	for _, c := range s.groupColumns {
		fieldNames = append(fieldNames, c.MapKey)
	}
	return
}
func (s *Sheet) GetOtherColumns() (fieldNames []string) {
	for _, c := range s.otherColumns {
		fieldNames = append(fieldNames, c.MapKey)
	}
	return
}

func (s *Sheet) GetDateCostColumns() map[string]string {
	return map[string]string{"Date": "Cost"}
}

// SetDataset overwrites the .dataset property
func (s *Sheet) SetDataset(ds map[string]map[string][]string) (err error) {
	s.dataset = ds
	return
}

// Write generates all the content for the sheet
// - main func to call!
func (s *Sheet) Write(f *excelize.File) (err error) {
	i, err := f.NewSheet(s.name)
	s.headers(f)
	s.rows(f)
	f.SetActiveSheet(i)
	return
}

// Cell will retrieve the value and type information for cell
// at the row|col passed
func (s *Sheet) Cell(row int, col int) (c CellInfo, ok bool) {
	c, ok = s.cells[CellRef{Row: row, Col: col}]
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
	err = f.AddTable(s.name, rangeRef, &excelize.TableOptions{
		StyleName: "TableStyleMedium9",
	})
	f.AutoFilter(s.name, rangeRef, &excelize.AutoFilterOptions{})
	return
}

// === internal
func (s *Sheet) Init() {
	s.cells = make(map[CellRef]CellInfo)
	s.styles = make(map[CellRef]*excelize.Style)
	s.SetVisible(true)
}

// headers writes the column data into the file passed
func (s *Sheet) headers(f *excelize.File) {
	s.rowCount = 1
	s.colCount = 1
	// this writes the headers
	for _, col := range s.columns {
		cell := CellReference(s.rowCount, s.colCount)
		f.SetCellValue(s.name, cell, col.Display)
		// store the cell
		s.cells[CellRef{Row: s.rowCount, Col: s.colCount}] = CellInfo{Value: col.Display}
		s.colCount++
	}
}

// rows iterates over the dataset, then loops over the columns
// get the value and writes that to the file
// -- handles formula overwrite for values
func (s *Sheet) rows(f *excelize.File) {
	// now add the data
	for _, row := range s.dataset {
		s.rowCount++
		s.colCount = 1
		// now loop over the columns and fetch that data from the row
		for _, col := range s.columns {
			cell := CellReference(s.rowCount, s.colCount)
			style := s.style(s.rowCount, s.colCount)
			var values []string
			// formula check here, overwrite the values to be the formula
			if len(col.Formula) > 0 {
				values = []string{col.Formula}
			} else {
				values = row[col.MapKey]
			}
			v, t, st, _ := CellWrite(cell, values, s.name, s.rowCount, f, style)
			// track the cell
			s.cells[CellRef{Row: s.rowCount, Col: s.colCount}] = CellInfo{Value: v, ValueType: t, Style: st}
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

func NewSheet(name string) Sheet {
	s := Sheet{}
	s.Init()
	s.SetName(name)
	return s
}
