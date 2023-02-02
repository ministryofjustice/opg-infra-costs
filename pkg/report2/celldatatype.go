package report2

import (
	"fmt"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/debugger"
	"strconv"
	"time"
)

// -- Errors
const (
	CellDataTypeErrorEmpty   string = "param `values` is empty"
	CellDataTypeErrorNoMatch string = "unable to match data type of [%v]"
)

const (
	CellDataIsDate    string = "date"
	CellDataIsFormula string = "formula"
	CellDataIsNumber  string = "number"
	CellDataIsString  string = "string"
)

type CellDataType struct {
	Values []interface{}
	Type   string
}

// Value handles covnerting set of string values for the cell
// into a single value
//   - For CellDataIsNumber this is a sum
//   - called from Cell.Value()
//   - needs the cell location to handle formula string replacements
func (cdt *CellDataType) Value(location Location) (value interface{}) {
	defer debugger.Log("CellDataType.Value()", debugger.VVERBOSE)()
	if len(cdt.Values) > 0 {
		if cdt.Type == CellDataIsNumber {
			// If this is a number we run a sum on all the contents
			sum := 0.0
			for _, str := range cdt.Values {
				v, _ := strconv.ParseFloat(str.(string), 64)
				sum += v
			}
			value = sum

		} else if cdt.Type == CellDataIsFormula {
			// If this is a formula, that parse out the values
			value = Substitute(location, cdt.Values[0].(string))
		} else {
			value = cdt.Values[0]
		}
	}
	return
}

// CellDataTypeFromValues is used to change how we handle different string values
// Those that look like a number should be handled differently to items that
// appear to be a YYYY-MM date
//   - uses the first item in values passed, checks its `(.type)`
//     property and sets flags
//   - values have to be an slice interfaces for type switch to work
//   - Type match order - int, float, string might be float, string might be int, string might be
//     date, string might be a formula
func CellDataTypeFromValues(values []interface{}) (cellDataType CellDataType, err error) {
	defer debugger.Log("CellDataTypeFromValues", debugger.VVERBOSE)()
	if len(values) > 0 {
		cellDataType.Values = values
		value := values[0]
		// set a default error
		err = fmt.Errorf(CellDataTypeErrorNoMatch, value)
		switch value.(type) {
		case float64:
			err = nil
			cellDataType.Type = CellDataIsNumber
		case int:
			err = nil
			cellDataType.Type = CellDataIsNumber
		case string:
			err = nil
			str := value.(string)
			cellDataType.Type = CellDataIsString

			if _, e := strconv.ParseFloat(str, 64); e == nil {
				cellDataType.Type = CellDataIsNumber
			} else if _, e := strconv.Atoi(str); e == nil {
				cellDataType.Type = CellDataIsNumber
			} else if _, e := time.Parse(dates.YM, str); e == nil {
				cellDataType.Type = CellDataIsDate
			} else if len(str) > 0 && str[0:1] == "=" {
				cellDataType.Type = CellDataIsFormula
			}
		}

	} else {
		err = fmt.Errorf(CellDataTypeErrorEmpty)
	}

	return
}
