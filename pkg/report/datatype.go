package report

import (
	"fmt"
	"strconv"
	"time"
)

// DataType looks at the value and tries to covnert it
func DataType(value interface{}) (dt ValueDataType, err error) {

	err = fmt.Errorf("unable to match data type of [%v]", value)

	switch value.(type) {
	case float64:
	case int:
		err = nil
		dt = DataIsANumber
	case string:
		err = nil
		dt = DataIsAString
		str := value.(string)
		// check if its a string that could be a number, date or formula
		if _, e := strconv.ParseFloat(str, 64); e == nil {
			dt = DataIsANumber
		} else if _, e := strconv.Atoi(str); e == nil {
			dt = DataIsANumber
		} else if _, e := time.Parse(DATEFORMAT, str); e == nil {
			dt = DataIsADate
		} else if len(str) > 0 && str[0:1] == "=" {
			dt = DataIsAFormula
		}
	}

	return
}
