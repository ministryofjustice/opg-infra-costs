package report2

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestSetGetDataMapHeadings(t *testing.T) {
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

	pp.Println(COLNUMBERFORMATS)
}
