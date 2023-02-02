package report2

import (
	"fmt"
	"opg-infra-costs/pkg/debugger"
)

type Sheet struct {
	// -- pre data load
	Name     string // the worksheet name
	Key      string // used in formula parsing, no spaces etc
	Visible  bool
	Headings []Column

	dataset    map[string]map[string][]string
	dataLoaded bool
	rowCounter int
	colCounter int

	Cells []Cell
}

func (s *Sheet) SetDataset(dataset map[string]map[string][]string) {
	defer debugger.Log("Sheet.SetDataset()", debugger.VERBOSE)()
	s.dataset = dataset
	s.rowCounter = 1
	s.colCounter = 1

	// -- headings
	setDataMapHeadings(s.Key, s.Headings)
	// convert headings into cells
	for _, h := range s.Headings {
		location := Location{Row: s.rowCounter, Col: s.colCounter, Sheet: s.Key}
		cell, err := NewCell(location, []interface{}{h.Name()})
		if err == nil {
			s.Cells = append(s.Cells, cell)
		}
		s.colCounter++
	}

	// -- data
	s.rowCounter++
	s.colCounter = 1
	// now deal with the main dataset

	s.dataLoaded = true
}

// -- New
func NewSheet(name string, key string, report ReportConfiguration, cfg *Configuration) Sheet {
	defer debugger.Log(fmt.Sprintf("Created NewSheet(n:%s, k:%s)", name, key), debugger.INFO)()
	s := Sheet{
		Name:       name,
		Key:        key,
		Visible:    report.Visible,
		dataLoaded: false,
	}

	// -- set all the data we can before loading the dataset
	setDataMapReportName(key, name)
	headings, _ := report.Headings(cfg.ColumnDefinitions)
	s.Headings = headings
	setDataMapHeadings(key, headings)
	setColumnNumberFormats(key, headings)

	return s
}
