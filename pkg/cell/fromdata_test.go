package cell

import (
	"reflect"
	"testing"
)

func TestCellInterfaceFromData(t *testing.T) {
	var ty reflect.Type

	// this should become a DateDateType
	c, err := FromData("2022-01")
	d := &DateDataType[string]{}

	if err != nil {
		t.Errorf("unexpected error recieved [%v]", err)
	}
	ty = reflect.TypeOf(c)
	if ty != d.Type() {
		t.Errorf("types do not match, expected [%v], actual [%v]", d.Type(), ty)
	}

	c, err = FromData("0.00")
	fl := &FloatDataType[float64]{}

	if err != nil {
		t.Errorf("unexpected error recieved [%v]", err)
	}
	ty = reflect.TypeOf(c)
	if ty != fl.Type() {
		t.Errorf("types do not match, expected [%v], actual [%v]", fl.Type(), ty)
	}

	// fail with unsupport type
	c, err = FromData(true)
	if err == nil {
		t.Errorf("expected an error for unsupported type [bool], actual [%v] [%T]", err, c)
	}

	//
	c, err = FromData("generic string header")
	if err != nil {
		t.Errorf("unexpected error for [string], actual [%v] [%T]", err, c)
	}

}
