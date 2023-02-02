package report2

import (
	"testing"
)

func TestSetReportName(t *testing.T) {
	defer _reset()
	name := "test sheet"
	key := "testSheet"
	setDataMapReportName(key, name)
	actual, _ := getDataMapReportName(key)
	if actual != name {
		t.Errorf("expected report to be called [%v], actual [%v]", name, actual)
	}

}

func TestSetGetDataMapHeadings(t *testing.T) {
	defer _reset()
	var err error
	sheet := "test-map-headings"
	headings := []Column{
		{SourceColumn: "AccountName"},
		{SourceColumn: "AccountEnvironment", Display: "Env"},
	}

	setDataMapHeadings(sheet, headings)
	expected := 1
	actual, _, err := getDataMapHeading(sheet, "AccountName")
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if actual != expected {
		t.Errorf("expected column to be mapped as [%v], actual [%v]", expected, actual)
	}

	expectedL := "B"
	_, actualL, err := getDataMapHeading(sheet, "AccountEnvironment")
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if actualL != expectedL {
		t.Errorf("expected column to be mapped as [%v], actual [%v]", expectedL, actual)
	}

}

func TestSetGetNumFmts(t *testing.T) {
	defer _reset()
	var err error
	sheet := "test-map-headings"
	def := ColumnDefinition{Style: 10}
	headings := []Column{
		{SourceColumn: "AccountName"},
		{SourceColumn: "AccountEnvironment", Display: "Env"},
		{SourceColumn: "Percent", Display: "P", Definition: def},
	}

	setColumnNumberFormats(sheet, headings)
	expected := def.Style
	actual, err := getColumnNumberFormat(sheet, "Percent")

	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if actual != expected {
		t.Errorf("expected numfmt to be [%v], actual [%v]", expected, actual)
	}

}
