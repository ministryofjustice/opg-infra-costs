package worksheet

import (
	"opg-infra-costs/pkg/cell"
	"opg-infra-costs/pkg/row"

	"github.com/k0kubun/pp"
)

type DetailedBreakdownWithRegion struct {
	name       string
	index      int
	rowCounter int

	defaultHeaderColumns []cell.ColumnHeadData
	columns              []cell.ColumnHeadDataType[cell.ColumnHeadData]

	rawRows []map[string]string
	rows    []row.RowInterface
}

// SetName uses the string passed to update the internal .name property
// -- todo: should this validate / strip characters out of `n`?
func (s *DetailedBreakdownWithRegion) SetName(n string) error {
	s.name = n
	return nil
}

// GetName returns the internal .name property
func (s *DetailedBreakdownWithRegion) GetName() (string, error) {
	return s.name, nil
}

// SetIndex sets this sheets index within the spreadsheet
// - typically the value from excelize.NewSheet()
func (s *DetailedBreakdownWithRegion) SetIndex(i int) error {
	s.index = i
	return nil
}

// GetIndex returns the stored value
func (s *DetailedBreakdownWithRegion) GetIndex() (int, error) {
	return s.index, nil
}

// GetVisible for this type of sheet is always false
// - we dont want to have this sheet visible by default
func (s *DetailedBreakdownWithRegion) GetVisible() (bool, error) {
	return false, nil
}

// GetActive always returns false
// - as we dont want this sheet to be visible, it can never be marked as active
func (s *DetailedBreakdownWithRegion) GetActive() (bool, error) {
	return false, nil
}

// GetTableConfiguration always returns an enabled version
func (s *DetailedBreakdownWithRegion) GetTableConfiguration() (SheetTableConfigurationInterface, error) {
	return &EnabledTable{}, nil
}

// GetFilterConfiguration always returns an enabled version
func (s *DetailedBreakdownWithRegion) GetFilterConfiguration() (SheetFilterConfigurationInterface, error) {
	return &EnabledFilter{}, nil
}

// GetPaneConfiguration always returns an enabled version
func (s *DetailedBreakdownWithRegion) GetPaneConfiguration() (SheetPaneConfigurationInterface, error) {
	return &EnabledPane{}, nil
}

// GetRowCount returns the current value of the rowCounter
func (s *DetailedBreakdownWithRegion) GetRowCount() (int, error) {
	return s.rowCounter, nil
}

// DefaultColumns returns the standard columns for this type of sheet
// -- AccountName
// -- AccountEnvironment
// -- Service
// -- Region
// other fields such as month, total, trend are added differently
func (s *DetailedBreakdownWithRegion) GetDefaultColumns() (defaults []cell.ColumnHeadDataType[cell.ColumnHeadData]) {
	var err error
	var c cell.ColumnHeadDataType[cell.ColumnHeadData]

	s.defaultHeaderColumns = []cell.ColumnHeadData{
		{Key: "AccountName", Display: "Account"},
		{Key: "AccountEnvironment", Display: "Environment"},
		{Key: "Service", Display: "Service"},
		{Key: "Region", Display: "Region"},
	}

	for _, h := range s.defaultHeaderColumns {
		var val []interface{} = []interface{}{h}
		c, err = cell.NewColumnHeader(val)
		if err == nil {
			defaults = append(defaults, c)
		}
	}

	return
}

// SetColumns allows additional columns to be attached to this sheet
//   - typically the date columns which are dynammic or the total / trend
//     cols that get added to the end
func (s *DetailedBreakdownWithRegion) SetColumns(columns ...cell.ColumnHeadDataType[cell.ColumnHeadData]) (err error) {
	s.columns = append(s.columns, columns...)
	return
}

// GetColumns simply returns the column set
func (s *DetailedBreakdownWithRegion) GetColumns() (columns []cell.ColumnHeadDataType[cell.ColumnHeadData], err error) {
	columns = s.columns
	return
}

// SetRawData take the row as raw strings assigns it to internal `.rawRows`
// - This is additive and uses append, so multiple calls will add more raw data
func (s *DetailedBreakdownWithRegion) SetRawData(rows []map[string]string) (err error) {
	s.rawRows = append(s.rawRows, rows...)
	return
}

// GetRawData returns all raw data
func (s *DetailedBreakdownWithRegion) GetRawData() ([]map[string]string, error) {
	return s.rawRows, nil
}

func (s *DetailedBreakdownWithRegion) ConvertData() (err error) {
	// reset counter and row data
	s.rows = []row.RowInterface{}
	s.rowCounter = 1

	columns, _ := s.GetColumns()
	pp.Println(columns)
	// create the header with the cols
	headerRow, _ := row.New(s.rowCounter, true, true)
	hCells := []cell.CellInterface{}
	for _, c := range columns {
		hCells = append(hCells, &c)
	}
	headerRow.SetDefinedCells(hCells)
	s.rows = append(s.rows, headerRow)

	// now loop over the raw data and create rows for each
	rawRows, _ := s.GetRawData()
	for _, raw := range rawRows {
		s.rowCounter++
		// new row!
		r, _ := row.New(s.rowCounter, false, true)
		// now we loop over columns and create new cells based on their content
		for _, col := range columns {
			// create the cell
			if v, ok := raw[col.Key()]; ok {
				c, _ := cell.New(false, []interface{}{v})
				r.SetDefinedCells([]cell.CellInterface{c})
			}
		}
		// add rows!
		s.rows = append(s.rows, r)
	}

	return
}
