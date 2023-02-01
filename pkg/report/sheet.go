package report

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	name           string
	columns        []Column
	dataset        map[string]map[string][]string
	rowCount       int
	colCount       int
	visible        bool
	cells          map[CellRef]CellInfo
	styles         map[CellRef]*excelize.Style
	tableOptions   *excelize.TableOptions
	filterOptions  *excelize.AutoFilterOptions
	groupColumns   []Column
	dateColumns    []Column
	otherColumns   []Column
	hideRowWhen    map[CellRef]interface{}
	rowKeyIndexMap map[string]RowKeyIndexSet
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
func (s *Sheet) SetDataset(ds map[string]map[string][]string) (mapped map[string]RowKeyIndexSet, err error) {
	s.dataset = ds
	s.headers()
	s.rows()

	mapped = s.rowKeyIndexMap
	return
}

// Write generates all the content for the sheet
// - main func to generate data for the sheet
func (s *Sheet) Write(f *excelize.File) (i int, err error) {
	i, err = f.NewSheet(s.name)
	for _, cell := range s.cells {
		err = CellWriter(cell, s.name, f)
	}
	// this is slow call
	// s.adjustColWidth(f)

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
func (s *Sheet) SetHideRowWhen(criteria map[CellRef]interface{}) (err error) {
	s.hideRowWhen = criteria
	return
}

// RowVisibility uses hideRowWhen to hide certain rows
// The row is shown by default, but is hidden if any of the criteria are true
//   - When float64 version of the cell value is less than the (int|float64)
//     version of the comparison the row is hidden
//   - When the string version of the cell is an exact match to the string
//     version of the comparison the row is also hidden
//   - CellRef.Row is never used, so make this a 0 or lower. This allows
//     multiple checks against same column
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
				// float parsing of the cell value
				val, pErr := strconv.ParseFloat(cellVal, 64)
				if pErr != nil {
					val = 0.0
				}

				// For int / float comparisons - when the cell value is less than (<)
				//	the criteria value this row will be hidden
				// For string comparisons - when the value of the cell matches (==)
				//	the criteria value the row will be hidden
				switch crit.(type) {
				case int:
					comp := float64(crit.(int))
					val = math.Abs(val)
					if val < comp {
						showRow = false
					}
				case float64:
					comp := crit.(float64)
					val = math.Abs(val)
					if val < comp {
						showRow = false
					}
				case string:
					if cellVal == crit.(string) {
						showRow = false
					}
				}
			}

			if !showRow {
				hidden = append(hidden, i)
				f.SetRowVisible(s.name, i, false)
			}
		}

	}
	return
}

// AddCell pushes new values into the cells slice
// -- will overwrite
func (s *Sheet) AddCell(row int, col int, value string) (CellInfo, error) {
	var foundAt CellRef
	var found bool = false
	var c CellInfo

	if col > len(s.columns) {
		return c, fmt.Errorf("column out of range")
	}
	if row > s.rowCount {
		return c, fmt.Errorf("row out of range")
	}

	formulaReplacements := map[string]interface{}{
		"r": strconv.Itoa(row),
		"c": strconv.Itoa(col),
	}
	// look for an existing cell
	for k, cell := range s.cells {
		if cell.Ref.Row == row && cell.Ref.Col == col {
			foundAt = k
			found = true
		}
	}
	if !found {
		foundAt = CellRef{Row: row, Col: col, RowKey: "additional"}
	}
	c = NewCellInfo(foundAt, s.style(row, col))
	header := s.columns[col]
	c.SetValue(header, []string{value}, formulaReplacements)

	s.cells[foundAt] = c
	return c, nil
}

// AddRow takes a map os values and a row key and pushes a set if new cells into .cells
func (s *Sheet) AddRow(rowKey string, row map[string][]string) (mapped map[string]RowKeyIndexSet, err error) {
	s.rowCount++
	s.colCount = 1
	formulaReplacements := map[string]interface{}{
		"r": strconv.Itoa(s.rowCount),
		"c": "1",
	}
	//key to index, as range over map is not consistent
	s.rowKeyIndexMap[rowKey] = RowKeyIndexSet{Index: s.rowCount, Columns: make(map[string]int)}
	// now loop over the columns and fetch that data from the row
	for _, col := range s.columns {
		// meta data
		s.rowKeyIndexMap[rowKey].Columns[col.Key()] = s.colCount
		// update the column number
		formulaReplacements["c"] = strconv.Itoa(s.colCount)
		ref := CellRef{Row: s.rowCount, Col: s.colCount, RowKey: rowKey}
		c := NewCellInfo(ref, s.style(s.rowCount, s.colCount))
		c.SetValue(col, row[col.MapKey], formulaReplacements)
		s.cells[ref] = c

		s.colCount++
	}

	mapped = s.rowKeyIndexMap
	return
}

// Init runs all the default setup for the sheet
func (s *Sheet) Init() {
	s.cells = make(map[CellRef]CellInfo)
	s.styles = make(map[CellRef]*excelize.Style)
	s.tableOptions = defaultTableOptions
	s.filterOptions = defaultFilterOptions
	s.rowKeyIndexMap = make(map[string]RowKeyIndexSet)
	s.SetVisible(true)
}

// === internal
func (s *Sheet) adjustColWidth(f *excelize.File) {
	l := len(s.groupColumns)
	cols, _ := f.GetCols(s.GetName())
	// loop over columns
	for idx, column := range cols {
		maxWidth := 0
		if idx < l {
			// loop over cells in the column to find longest
			for _, cell := range column {
				width := utf8.RuneCountInString(cell) + 4
				if width > maxWidth {
					maxWidth = width
				}
			}
			columnName, _ := excelize.ColumnNumberToName(idx + 1)
			f.SetColWidth(s.GetName(), columnName, columnName, float64(maxWidth))
		}
	}
}

// headers generates the cell data for the current .columns
func (s *Sheet) headers() map[string]RowKeyIndexSet {
	s.rowCount = 1
	s.colCount = 1
	// this writes the headers
	rowKey := "Header"
	s.rowKeyIndexMap[rowKey] = RowKeyIndexSet{Index: s.rowCount, Columns: make(map[string]int)}
	for _, col := range s.columns {
		ref := CellRef{Row: s.rowCount, Col: s.colCount, RowKey: rowKey}
		s.cells[ref] = CellInfo{Value: col.Display, Ref: ref}

		s.rowKeyIndexMap[rowKey].Columns[col.Key()] = s.colCount
		s.colCount++
	}
	return s.rowKeyIndexMap
}

// rows iterates over the dataset, then loops over the columns
// get the value and writes that to the `.cells`
// -- will parse formula values as well
func (s *Sheet) rows() map[string]RowKeyIndexSet {
	for rowKey, row := range s.dataset {
		s.AddRow(rowKey, row)
	}
	return s.rowKeyIndexMap
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
