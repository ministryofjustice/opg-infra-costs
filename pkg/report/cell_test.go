package report

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestCellReference(t *testing.T) {

	a1 := CellReference(1, 1)
	if a1 != "A1" {
		t.Errorf("cell reference failed, recieved [%v]", a1)
	}

	// -- test wrapping around 26
	c := CellReference(20, 80)
	if c != "CB20" {
		t.Errorf("cell reference failed, recieved [%v]", c)
	}
}

func TestCellValueFromType(t *testing.T) {
	var err error
	//CellValueFromType(values []string, dt ValueDataType) (value interface{}, err error)
	cv, err := CellValueFromType([]string{"100", "101", "99", "10.1", "-0.1", "-10"}, DataIsANumber)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if cv.(float64) != 300 {
		t.Errorf("expected sum total of values, actual [%v]", cv)
	}

	cv, err = CellValueFromType([]string{"foo", "bar", "hello"}, DataIsAString)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if cv.(string) != "foo" {
		t.Errorf("expected first value to be returned, actual [%v]", cv)
	}

	cv, err = CellValueFromType([]string{"2022-01", "2022-02", "2022-01"}, DataIsADate)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}
	if cv.(string) != "2022-01" {
		t.Errorf("expected first value to be returned, actual [%v]", cv)
	}

}

func TestCellWrite(t *testing.T) {
	var err error
	f := excelize.NewFile()
	f.NewSheet("testcellwrite")

	// CellWrite( ref string, values []string, sheet string, rowCount int, f *excelize.File, style *excelize.Style) (val interface{}, dt ValueDataType, s *excelize.Style, err error)

	val, dt, _, err := CellWrite("A1", []string{"100.01", "99.0"}, "testcellwrite", 1, f, defaultStyle)
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

	if val.(float64) != 199.01 {
		t.Errorf("returned value does not match, recieved [%v]", val)
	}
	if dt != DataIsANumber {
		t.Errorf("expected to be a number, recieved [%v]", dt)
	}
}
