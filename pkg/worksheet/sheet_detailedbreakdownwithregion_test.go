package worksheet

import (
	"opg-infra-costs/pkg/cell"
	"testing"

	"github.com/k0kubun/pp"
)

func TestDetailedBreakdownWithRegionSimple(t *testing.T) {
	var err error
	var v string
	var i int
	var b bool
	var tbl SheetTableConfigurationInterface
	var f SheetFilterConfigurationInterface
	var p SheetPaneConfigurationInterface

	s := DetailedBreakdownWithRegion{}
	// -- name
	err = s.SetName("")
	v, _ = s.GetName()
	if err != nil || len(v) != 0 {
		t.Errorf("unexpected result from setting a blank name; error [%v] value [%v]", err, v)
	}

	err = s.SetName("name")
	v, _ = s.GetName()
	if err != nil || v != "name" {
		t.Errorf("unexpected result from setting a name; error [%v] value [%v]", err, v)
	}

	// -- index
	err = s.SetIndex(0)
	i, _ = s.GetIndex()
	if err != nil || i != 0 {
		t.Errorf("unexpected result from setting a index to 0; error [%v] value [%v]", err, i)
	}

	err = s.SetIndex(2)
	i, _ = s.GetIndex()
	if err != nil || i != 2 {
		t.Errorf("unexpected result from setting index; error [%v] value [%v]", err, i)
	}

	if b, _ = s.GetActive(); b {
		t.Errorf("sheet [%T] should not be active: %v", s, b)
	}
	if b, _ = s.GetVisible(); b {
		t.Errorf("sheet [%T] should not be visible: %v", s, b)
	}

	tbl, _ = s.GetTableConfiguration()
	b, _ = tbl.GetEnabled()
	if !b {
		t.Errorf("table [%T] should be enabled: %v", tbl, b)
	}

	f, _ = s.GetFilterConfiguration()
	b, _ = f.GetEnabled()
	if !b {
		t.Errorf("filter [%T] should be enabled: %v", f, b)
	}

	p, _ = s.GetPaneConfiguration()
	b, _ = p.GetEnabled()
	if !b {
		t.Errorf("pane [%T] should be enabled: %v", f, b)
	}

}

func TestDetailedBreakdownWithRegionDefaultColumns(t *testing.T) {
	var err error
	var cols []cell.ColumnHeadDataType[cell.ColumnHeadData]

	s := DetailedBreakdownWithRegion{}
	cols = s.GetDefaultColumns()

	if err != nil {
		t.Errorf("unexpected error: [%v]", err)
	}
	if len(cols) == 0 {
		t.Errorf("expected more columns: [%v]", len(cols))
	}

}

func TestDetailedBreakdownWithRegionSetAndGetColumns(t *testing.T) {
	var err error

	s := DetailedBreakdownWithRegion{}
	cols := s.GetDefaultColumns()
	// set the default columns
	s.SetColumns(cols...)
	got, err := s.GetColumns()

	if len(got) != len(cols) {
		t.Errorf("default column count and result should match [%d], recieved [%d]", len(cols), len(got))
	}
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

	// now we add some columns
	c, _ := cell.NewColumnHeader([]interface{}{"2022-01"})
	err = s.SetColumns(c)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

	got, err = s.GetColumns()
	if len(got) != len(cols)+1 {
		t.Errorf("column count should be one higher than default [%d], recieved [%d]", len(cols), len(got))
	}
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

}

func TestDetailedBreakdownWithRegionSetRawData(t *testing.T) {
	var err error
	s := DetailedBreakdownWithRegion{}

	raw := []map[string]string{
		{"A": "1"},
		{"B": "2"},
	}
	add := []map[string]string{
		{"C": "3"},
	}

	err = s.SetRawData(raw)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	err = s.SetRawData(add)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

	a, _ := s.GetRawData()
	if len(a) != 3 {
		t.Errorf("expected to get 3 raw rows back, actual [%v]", len(a))
	}

}

func TestDetailedBreakdownWithRegionConvertData(t *testing.T) {

	s := DetailedBreakdownWithRegion{}
	s.SetColumns(s.GetDefaultColumns()...)

	//some sample raw data - some columns will be ignored this time
	raw := []map[string]string{
		{
			"AccountName":        "AccountA",
			"AccountEnvironment": "E1",
			"Service":            "AWS 1",
			"Region":             "R1",
			"2022-01":            "0.17",
			"2022-02":            "3.37",
			"2022-03":            "5.37",
		},
		{
			"AccountName":        "AccountA",
			"AccountEnvironment": "E1",
			"Service":            "AWS 2",
			"Region":             "R1",
			"2022-01":            "12.57",
			"2022-02":            "101",
			"2022-03":            "-0.10",
		},
	}

	s.SetRawData(raw)
	s.ConvertData()
	pp.Println(s.rows[0])

}
