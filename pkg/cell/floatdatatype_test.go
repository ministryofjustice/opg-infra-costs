package cell

import (
	"testing"
)

func TestFloatDataTypeParse(t *testing.T) {
	c := FloatDataType[float64]{}

	v, err := c.Parse("not a float")
	if err == nil {
		t.Errorf("expected Parse to return an error for non-float, recieved [%v]", err)
	}
	if v != nil {
		t.Errorf("expected Parse to return a nil for non-float, recieved [%v]", v)
	}

	v, err = c.Parse("10.11")
	if v != 10.11 {
		t.Errorf("expected Parse to return matching float, recieved [%v]", v)
	}
	if err != nil {
		t.Errorf("unexpected error from Parse: [%v]", err)
	}

	v, err = c.Parse(101.111)
	if v != 101.111 {
		t.Errorf("expected Parse to return  matching float, recieved [%v]", v)
	}
	if err != nil {
		t.Errorf("unexpected error from Parse: [%v]", err)
	}

	v, err = c.Parse(301)
	if v != 301.00 {
		t.Errorf("expected Parse to return  matching float, recieved [%v]", v)
	}
	if v == 301 {
		t.Errorf("expected Parse to return  matching float, but matched against an int [%v]", v)
	}
	if err != nil {
		t.Errorf("unexpected error from Parse: [%v]", err)
	}

}

func TestFloatDataTypeSet(t *testing.T) {

	c := FloatDataType[float64]{}

	err := c.Set("not a float")
	if err == nil {
		t.Errorf("expected Set to return an error for a non numeric value")
	}

	err = c.Set("not a float", 100)
	if err == nil {
		t.Errorf("expected Set to return an error for a non numeric value")
	}

	err = c.Set(100)
	if err != nil {
		t.Errorf("unexpected error from Set: [%v]", err)
	}
	err = c.Set(101, 1.01, 0.01)
	if err != nil {
		t.Errorf("unexpected error from Set: [%v]", err)
	}

}

func TestFloatDataTypeGetAll(t *testing.T) {

	c := FloatDataType[float64]{}
	c.Set("not a float", 100, 10.01, 0.11, 0.00, "not a float", 404)
	vals, _ := c.GetAll()

	if len(vals) != 5 {
		t.Errorf("expected GetAll to only return 5 items, actual [%v]", len(vals))
	}

}

func TestFloatDataTypeGet(t *testing.T) {

	c := FloatDataType[float64]{}
	c.Set("not a float", 100, 10.01, 0.11, 0.00, "not a float", 404, -1)
	val, _ := c.Get()

	if val != 513.12 {
		t.Errorf("expected Get to be [514.12], actual [%v]", val)
	}

}
