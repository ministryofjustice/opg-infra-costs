package report2

import (
	"testing"
)

var tnumbers []interface{} = []interface{}{"10", "100.0", "-9.1", "13.57"}
var tstrs []interface{} = []interface{}{"foo", "bar", "hello"}
var tdates []interface{} = []interface{}{"2022-01", "2022-03", "2022-10"}
var tformulas []interface{} = []interface{}{"=SUM({TST}{name})"}

func TestCellDataTypeFromValues(t *testing.T) {

	dtNum, _ := CellDataTypeFromValues(tnumbers)
	if dtNum.Type != CellDataIsNumber {
		t.Errorf("expected type [%v], actual [%v]", CellDataIsNumber, dtNum.Type)
	}

	dtStr, _ := CellDataTypeFromValues(tstrs)
	if dtStr.Type != CellDataIsString {
		t.Errorf("expected type [%v], actual [%v]", CellDataIsString, dtStr.Type)
	}

	dtDates, _ := CellDataTypeFromValues(tdates)
	if dtDates.Type != CellDataIsDate {
		t.Errorf("expected type [%v], actual [%v]", CellDataIsDate, dtDates.Type)
	}

	dtFormulas, _ := CellDataTypeFromValues(tformulas)
	if dtFormulas.Type != CellDataIsFormula {
		t.Errorf("expected type [%v], actual [%v]", CellDataIsFormula, dtFormulas.Type)
	}

}

func TestCellDataTypeValue(t *testing.T) {

	// -- numbers should be a sum
	dtNum, _ := CellDataTypeFromValues(tnumbers)
	value := dtNum.Value()
	expected := 114.47
	if value.(float64) != expected {
		t.Errorf("expected [%v], actual [%v]", expected, value)
	}

	dtDates, _ := CellDataTypeFromValues(tdates)
	valueD := dtDates.Value()
	expectedD := tdates[0]
	if valueD.(string) != expectedD {
		t.Errorf("expected [%v], actual [%v]", expectedD, valueD)
	}

	dtFormulas, _ := CellDataTypeFromValues(tformulas)
	// -- this an exact match on the formula with out subs
	valueF := dtFormulas.Value()
	expectedF := tformulas[0]
	if valueF.(string) != expectedF {
		t.Errorf("expected [%v], actual [%v]", expectedF, valueF)
	}
	// set a sub and try again
	SUBSTITUTIONS["{TST}{name}"] = "TEST WORKED"
	valueF = dtFormulas.Value()
	expectedF = "=SUM(TEST WORKED)"
	if valueF.(string) != expectedF {
		t.Errorf("expected [%v], actual [%v]", expectedF, valueF)
	}

}
