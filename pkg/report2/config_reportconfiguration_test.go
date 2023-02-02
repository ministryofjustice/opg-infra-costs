package report2

import (
	"testing"
)

func TestConfigReportHeadings(t *testing.T) {
	var cfg Configuration
	var reportCfg ReportConfiguration
	var colDefs map[string]ColumnDefinition
	var err error

	cfg, _ = unmarshalConfig([]byte(dummyCfg))
	reportCfg = cfg.Reports["DetailedBreakdown"]
	colDefs, _ = reportCfg.ColumnNamesToDefinitions(cfg.ColumnDefinitions)

	// detailed breakdown is standard version, should have simple
	// transpose count to test
	reportCfg = cfg.Reports["DetailedBreakdown"]
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
