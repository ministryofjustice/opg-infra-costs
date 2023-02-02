package report2

import (
	"testing"
)

func TestConfigReportHeadings(t *testing.T) {
	var cfg Configuration
	var reportCfg ReportConfiguration
	var colDefs map[string]ColumnDefinition
	var err error

	nm := "dDetailedBreakdown"
	cfg, _ = unmarshalConfig([]byte(dummyCfg))
	reportCfg = cfg.Reports[nm]
	colDefs, _ = reportCfg.ColumnNamesToDefinitions(cfg.ColumnDefinitions)

	// detailed breakdown is standard version, should have simple
	// transpose count to test
	reportCfg = cfg.Reports[nm]
	extraCols := 0
	transposed := 0
	for _, tCol := range Transposed(colDefs) {
		extraCols += len(tCol.ToHeadings())
		transposed++
	}
	expected := len(reportCfg.ColumnsByName) - transposed + extraCols
	headings, err := reportCfg.Headings(colDefs)
	actual := len(headings)

	if actual != expected {
		t.Errorf("expected [%v] headings, actual [%v]", expected, actual)
	}

	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

}

func TestConfigReportSetHeadings(t *testing.T) {
	var err error

	cfg, _ := unmarshalConfig([]byte(dummyCfg))
	sheet := "dDetailedBreakdown"
	reportCfg := cfg.Reports[sheet]
	colDefs, _ := reportCfg.ColumnNamesToDefinitions(cfg.ColumnDefinitions)

	headings, err := reportCfg.Headings(colDefs)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
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
