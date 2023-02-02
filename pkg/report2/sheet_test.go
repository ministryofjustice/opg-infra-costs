package report2

import (
	"opg-infra-costs/pkg/debugger"
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

var simpleDataset = map[string]map[string][]string{
	"r2": {
		"AccountName":        []string{"Test1"},
		"AccountEnvironment": []string{"Production", "Production"},
		"Service":            []string{"Service 1"},
		"2022-01":            []string{"11.35", "-1.03"},
		"2022-02":            []string{"100.35", "-9.15", "+123.01"},
		"2022-03":            []string{"21.01", "0.003"},
	},
	"r3": {
		"AccountName":        []string{"Test1"},
		"AccountEnvironment": []string{"Production", "Production"},
		"Service":            []string{"Service 2"},
		"2022-01":            []string{"11.37", "-9.03"},
	},
	"r4": {
		"AccountName":        []string{"Test2"},
		"AccountEnvironment": []string{"Production", "Production"},
		"Service":            []string{"Service 2"},
		"2022-02":            []string{"7.95"},
	},
	"r5": {
		"AccountName":        []string{"Test2"},
		"AccountEnvironment": []string{"Development"},
		"Service":            []string{"Service 2"},
		"2022-02":            []string{"27.85"},
	},
	"r6": {
		"AccountName":        []string{"Test2"},
		"AccountEnvironment": []string{"Development"},
		"Service":            []string{"Service 2"},
		"2022-02":            []string{"27.85"},
		"2022-03":            []string{"57.73"},
	},
}

func TestNewSheet(t *testing.T) {
	defer _reset()
	cfg, _ := unmarshalConfig([]byte(dummyCfg))
	key := "dDetailedBreakdown"
	report := cfg.Reports[key]
	s := NewSheet("Detailed Breakdown", key, report, &cfg)

	s.SetDataset(simpleDataset)

	//row := s.Cells.Row(2)
	//pp.Println(s.Headings)

	pp.Println(SHEETDATAMAP)

}

func TestMain(m *testing.M) {
	debugger.LEVEL = debugger.ERR
	_reset()
	code := m.Run()
	_reset()
	os.Exit(code)
}
