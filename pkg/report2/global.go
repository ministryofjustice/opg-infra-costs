package report2

import (
	"fmt"
	"opg-infra-costs/pkg/debugger"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// -- Errors
const (
	HeadingNotFound   string = "heading [%v] not found in sheet [%s] data"
	ReportKeyNotFound string = "report [%v] not found in data"
)

// SHEETDATAMAP is used to handle string replacements from
// config files and dynamic to column / row locations as
// the order of data is not always the same due to the use
// of loops maps
//   - Key will be formatted as {sheetKey}{type:key}, for example
//     {Totals}{Col:Org}, or {Detailed}{Col:2022-02}
//   - ${Col:Org} needs to be mapped to {CurrentSheet}{Col:Org}
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
const (
	// string replacements
	_currentSheet   string = "$$"
	_previousRow    string = "{row-1}"
	_currentRow     string = "{row}"
	_currentCol     string = "{col}"
	_transposeStart string = "{transposeStart}"
	_transposeEnd   string = "{transposeEnd}"

	_headingPrefix string = "Col"
)

func _reset() {
	SHEETDATAMAP = map[string]string{}
	ROWNUMBERFORMATS = map[string]int{}
	COLNUMBERFORMATS = map[string]int{}
	CELLNUMBERFORMATS = map[string]int{}
}

// setDataMapReportName tracks the name of this sheet via a
// easier string for formula parsing (so "Cost Changes" => {CostChanges})
func setDataMapReportName(sheetKey string, sheetName string) {
	defer debugger.Log("setDataMapReportName", debugger.VVERBOSE)()
	key := fmt.Sprintf("{%s}{name}", sheetKey)
	SHEETDATAMAP[key] = sheetName
}

// getDataMapReportName
func getDataMapReportName(sheetKey string) (s string, err error) {
	defer debugger.Log("getDataMapReportName", debugger.VVERBOSE)()
	key := fmt.Sprintf("{%s}", sheetKey)
	s, ok := SHEETDATAMAP[key]
	if !ok {
		err = fmt.Errorf(fmt.Sprintf(ReportKeyNotFound, sheetKey))
	}
	return
}

// setDataMapHeadings func allows setting the global SHEETDATAMAP
// with location data for the headers of the sheet
//   - Track sheet & column to an col number
//   - Add in specials for start & end of transpose data columns
func setDataMapHeadings(sheetKey string, headings []Column) {
	defer debugger.Log("setDataMapHeadings", debugger.VVERBOSE)()
	prefix := fmt.Sprintf("{%s}", sheetKey)
	transposeStart := 0
	transposeEnd := 0

	for i, h := range headings {
		col := i + 1
		key := fmt.Sprintf("%s{%s:%s}", prefix, _headingPrefix, h.Key())
		SHEETDATAMAP[key] = strconv.Itoa(col)

		if h.Definition.IsTransposed() && transposeStart == 0 {
			transposeStart = col
		}
		if h.Definition.IsTransposed() && transposeStart > 0 {
			transposeEnd = col
		}
	}

	if transposeStart > 0 && transposeEnd > transposeStart {
		tS := fmt.Sprintf("{%s}%s", sheetKey, _transposeStart)
		tE := fmt.Sprintf("{%s}%s", sheetKey, _transposeEnd)
		st, _ := excelize.ColumnNumberToName(transposeStart)
		et, _ := excelize.ColumnNumberToName(transposeEnd)
		SHEETDATAMAP[tS] = st
		SHEETDATAMAP[tE] = et
	}

}

// getDataMapHeading returns int and string versions (1 = A) of the heading
// passed for the sheet
func getDataMapHeading(sheetKey string, headingName string) (i int, s string, err error) {
	defer debugger.Log("getDataMapHeading", debugger.VVERBOSE)()
	key := fmt.Sprintf("{%s}{%s:%s}", sheetKey, _headingPrefix, headingName)
	if str, ok := SHEETDATAMAP[key]; ok {
		i, err = strconv.Atoi(str)
		s, _ = excelize.ColumnNumberToName(i)
	} else {
		err = fmt.Errorf(fmt.Sprintf(HeadingNotFound, headingName, sheetKey))
	}
	return
}

// setColumnNumberFormats logs the style for all headings in this sheet
// into the global set
func setColumnNumberFormats(sheetKey string, headings []Column) {
	defer debugger.Log("setColumnNumberFormats", debugger.VVERBOSE)()
	prefix := fmt.Sprintf("{%s}", sheetKey)
	for _, h := range headings {
		def := h.Definition
		if def.Style > 0 {
			key := fmt.Sprintf("%s{%s:%s}", prefix, _headingPrefix, h.Key())
			COLNUMBERFORMATS[key] = def.Style
		}
	}
}

// getColumnNumberFormat
func getColumnNumberFormat(sheetKey string, headingName string) (format int, err error) {
	defer debugger.Log("getColumnNumberFormat", debugger.VVERBOSE)()
	key := fmt.Sprintf("{%s}{%s:%s}", sheetKey, _headingPrefix, headingName)
	if v, ok := COLNUMBERFORMATS[key]; ok {
		format = v
	} else {
		err = fmt.Errorf(fmt.Sprintf(HeadingNotFound, headingName, sheetKey))
	}
	return
}

// == Parsing the data out for usage

func Substitute(location Location, original string) (updated string) {
	defer debugger.Log("Substitute", debugger.VVERBOSE)()
	// replace $ within the original to become the sheet name
	updated = strings.ReplaceAll(
		original,
		_currentSheet,
		fmt.Sprintf("{%s}", location.Sheet),
	)
	// handle keywords that relate the position of the cell (row/col etc)
	updated = strings.ReplaceAll(updated, _previousRow, strconv.Itoa(location.Row-1))
	updated = strings.ReplaceAll(updated, _currentRow, strconv.Itoa(location.Row))
	updated = strings.ReplaceAll(updated, _currentCol, strconv.Itoa(location.Col))

	for key, replacement := range SHEETDATAMAP {
		updated = strings.ReplaceAll(updated, key, replacement)
	}
	return
}
