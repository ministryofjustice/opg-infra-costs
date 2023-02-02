package report2

import "strings"

// SHEETDATAMAP is used to handle string replacements from
// config files and dynamic to column / row locations as
// the order of data is not always the same due to the use
// of loops maps
//   - Key will be formatted as {sheetName}{type:key}, for example
//     {Totals}{Col:Org}, or {Detailed}{Col:2022-02}
var SHEETDATAMAP = map[string]string{}

func Substitute(original string) (updated string) {
	updated = original
	for key, replacement := range SHEETDATAMAP {
		updated = strings.ReplaceAll(updated, key, replacement)
	}
	return
}

// ROWNUMBERFORMATS is used to track style data for
// every row between each sheet
// This is so any report can adjust / use styles from others
//   - map["{Totals}{Row:1}"] = 177 (0.00)
var ROWNUMBERFORMATS = map[string]int{}

// COLNUMBERFORMATS is used to track style data for
// every column
// This is so any report can adjust / use styles from others
//   - map["{CostChanges}{Col:FormulaIncrease%}"] = 10 (%.00)
var COLNUMBERFORMATS = map[string]map[string]int{}

// CELLNUMBERFORMATS
//   - map["{Totals}{Col:2022-01}{Row:2}"]
var CELLNUMBERFORMATS = map[string]map[string]int{}
