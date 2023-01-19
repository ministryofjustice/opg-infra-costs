package cell

import (
	"testing"
)

func TestStringDataTypeParse(t *testing.T) {
	c := StringDataType[string]{}
	// correct
	v, err := c.Parse("string")
	if err != nil {
		t.Errorf("parsing valid string failed: %v", err)
	}
	if v.(string) != "string" {
		t.Errorf("valid string did not match expected value: %v", v)
	}

	_, err = c.Parse(1.00)
	if err == nil {
		t.Errorf("expected an error for parsing an invalid string, recieved: %v", err)
	}

}
func TestStringDataTypeSet(t *testing.T) {

	c := StringDataType[string]{}
	// correct date
	err := c.Set("hello")
	if err != nil {
		t.Errorf("unexpected error setting a string: %v", err)
	}
	// these should fail
	err = c.Set(100.01)
	if err == nil {
		t.Errorf("expected an error, recieved: %v", err)
	}

	// test setting multiple that work
	err = c.Set("foo", "bar")
	if err != nil {
		t.Errorf("unexpected error setting a set: %v", err)
	}

	// test setting multiple with a failure
	err = c.Set("2022-01", "2022-02", 0.13)
	if err == nil {
		t.Errorf("expected error with one of the passed strings failed, recieved: %v", err)
	}
	err = c.Set("2022-01", 15.01, "2022-01")
	if err == nil {
		t.Errorf("expected Set to return with an error, recieved [%v]", err)
	}

}

func TestStringDataTypeGetAll(t *testing.T) {
	c := StringDataType[string]{}
	// set many, getall should match
	c.Set("hello", "world")

	all, err := c.GetAll()
	if len(all) != 2 {
		t.Errorf("expected GetAll to return 2 items, actual [%v]", len(all))
	}
	if err != nil {
		t.Errorf("expected GetAll to return without error, recieved [%v]", err)
	}
	// include some bad data in the middle
	c = StringDataType[string]{}

	err = c.Set("2022-01", false, "2022-01")
	if err == nil {
		t.Errorf("expected Set to return with an error, recieved [%v]", err)
	}

	all, _ = c.GetAll()
	if len(all) != 2 {
		t.Errorf("expected GetAll to return 2 items, actual [%v]", len(all))
	}

}

func TestStringDataTypeGet(t *testing.T) {

	c := StringDataType[string]{}
	// set several
	c.Set("foo", "bar", "hello", "world")
	// get should return just the first
	v, e := c.Get()
	if e != nil {
		t.Errorf("expected Get to return without error, recieved [%v]", e)
	}
	if v.(StringData).Display != "foo" {
		t.Errorf("expected get to return the display field of the first value passed to set")
	}

	c = StringDataType[string]{}
	// set several
	c.Set(true, false, 1.23)
	// get should return just the first
	v, e = c.Get()
	if e != nil {
		t.Errorf("expected Get to return without error, even when now items [%v]", e)
	}
	if v != nil {
		t.Errorf("expected Get to return empty string, recieved [%v]", v)
	}

	c = StringDataType[string]{}
	// set several
	c.Set(StringData{Key: "test", Display: "Test"})
	// get should return just the first
	v, _ = c.Get()
	if v.(StringData).Display != "Test" {
		t.Errorf("expected get to return the Display field, [%v]", v)
	}

}
