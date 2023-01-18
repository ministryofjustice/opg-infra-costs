package cell

import (
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
	"reflect"
	"testing"
)

func TestCellDataFromData(t *testing.T) {
	var ty reflect.Type

	// this should become a DateDateType
	c, err := FromData("2022-01")
	d := &datedatatype.DateDataType[string]{}

	if err != nil {
		t.Errorf("unexpected error recieved [%v]", err)
	}
	ty = reflect.TypeOf(c)
	if ty != d.Type() {
		t.Errorf("types do not match, expected [%v], actual [%v]", ty, d.Type())
	}

	c, err = FromData("0.00")
	fl := &floatdatatype.FloatDataType[float64]{}

	if err != nil {
		t.Errorf("unexpected error recieved [%v]", err)
	}
	ty = reflect.TypeOf(c)
	if ty != fl.Type() {
		t.Errorf("types do not match, expected [%v], actual [%v]", ty, fl.Type())
	}

}
