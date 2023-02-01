package report2

import "strings"

// SUBSTITUTIONS is used to handle string replacements from
// config files and dynamic to column / row locations as
// the order of data is not always the same due to the use
// of loops maps
//   - Key will be formatted as {sheetName}{type:key}, for example
//     {Totals}{Col:Org}, or {Detailed}{Col:2022-02}
var SUBSTITUTIONS = map[string]string{}

func Substitute(original string) (updated string) {
	updated = original
	for key, replacement := range SUBSTITUTIONS {
		updated = strings.ReplaceAll(updated, key, replacement)
	}
	return
}
