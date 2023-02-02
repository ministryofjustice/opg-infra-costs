package report2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// -- Errors
const (
	HeadingNotFound string = "heading [%v] not found in sheet [%s] data"
)

// SHEETDATAMAP is used to handle string replacements from
// config files and dynamic to column / row locations as
// the order of data is not always the same due to the use
// of loops maps
//   - Key will be formatted as {sheetName}{type:key}, for example
//     {Totals}{Col:Org}, or {Detailed}{Col:2022-02}
var SHEETDATAMAP = map[string]string{}

// ROWNUMBERFORMATS is used to track style data for
// every row between each sheet
// This is so any report can adjust / use styles from others
//   - map["{Totals}{Row:1}"] = 177 (0.00)
var ROWNUMBERFORMATS = map[string]int{}

// COLNUMBERFORMATS is used to track style data for
// every column
// This is so any report can adjust / use styles from others
//   - map["{CostChanges}{Col:FormulaIncrease%}"] = 10 (%.00)
var COLNUMBERFORMATS = map[string]int{}

// CELLNUMBERFORMATS
//   - map["{Totals}{Col:2022-01}{Row:2}"]
var CELLNUMBERFORMATS = map[string]int{}

// == Update the global data map
var headingPrefix string = "Col"

// setDataMapHeadings func allows setting the global SHEETDATAMAP
// with location data for the headers of the sheet
func setDataMapHeadings(sheetName string, headings []Column) {
	prefix := fmt.Sprintf("{%s}", sheetName)
	for i, h := range headings {
		// set the column index
		key := fmt.Sprintf("%s{%s:%s}", prefix, headingPrefix, h.Key())
		SHEETDATAMAP[key] = strconv.Itoa(i + 1)
	}
}

// getDataMapHeading returns int and string versions (1 = A) of the heading
// passed for the sheet
func getDataMapHeading(sheetName string, headingName string) (i int, s string, err error) {
	key := fmt.Sprintf("{%s}{%s:%s}", sheetName, headingPrefix, headingName)
	if str, ok := SHEETDATAMAP[key]; ok {
		i, err = strconv.Atoi(str)
		s, _ = excelize.ColumnNumberToName(i)
	} else {
		err = fmt.Errorf(fmt.Sprintf(HeadingNotFound, headingName, sheetName))
	}
	return
}

// setColumnNumberFormats logs the style for all headings in this sheet
// into the global set
func setColumnNumberFormats(sheetName string, headings []Column) {
	prefix := fmt.Sprintf("{%s}", sheetName)
	for _, h := range headings {
		def := h.Definition
		if def.Style > 0 {
			key := fmt.Sprintf("%s{%s:%s}", prefix, headingPrefix, h.Key())
			COLNUMBERFORMATS[key] = def.Style
		}
	}
}

// getColumnNumberFormat
func getColumnNumberFormat(sheetName string, headingName string) (format int, err error) {
	key := fmt.Sprintf("{%s}{%s:%s}", sheetName, headingPrefix, headingName)
	if v, ok := COLNUMBERFORMATS[key]; ok {
		format = v
	} else {
		err = fmt.Errorf(fmt.Sprintf(HeadingNotFound, headingName, sheetName))
	}
	return
}

// == Parsing the data out for usage
func Substitute(original string) (updated string) {
	updated = original
	for key, replacement := range SHEETDATAMAP {
		updated = strings.ReplaceAll(updated, key, replacement)
	}
	return
}
