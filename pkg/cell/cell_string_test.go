package cell

import (
	"testing"
)

func TestStringDataTypeParse(t *testing.T) {
	c := StringDataType[StringData]{}
	_, err := c.Parse("string")
	if err == nil {
		t.Errorf("expected parsing invalid stringdata to faile: [%v]", err)
	}
	_, err = c.Parse(1.00)
	if err == nil {
		t.Errorf("expected an error for parsing an invalid string, recieved: [%v]", err)
	}

	_, err = c.Parse(StringData{Key: "A", Display: "ACOL"})
	if err != nil {
		t.Errorf("unexpected error for parsing a valid stringdata, recieved: [%v]", err)
	}
	_, err = c.Parse(StringData{Key: "A"})
	if err != nil {
		t.Errorf("unexpected error for parsing a valid stringdata, recieved: [%v]", err)
	}
	_, err = c.Parse(StringData{})
	if err == nil {
		t.Errorf("expected error for parsing invalid stringdata - should have a property set, recieved: [%v]", err)
	}

}
func TestStringDataTypeSet(t *testing.T) {

	c := StringDataType[StringData]{}
	// these should fail
	err := c.Set("hello")
	if err == nil {
		t.Errorf("expected error setting an invalid string: %v", err)
	}
	err = c.Set(100.01)
	if err == nil {
		t.Errorf("expected an error, recieved: %v", err)
	}
	// test setting multiple failures
	err = c.Set("foo", "bar")
	if err == nil {
		t.Errorf("expected error when setting a set of false items: [%v]", err)
	}
	// test working
	err = c.Set(StringData{Key: "A"}, StringData{Key: "B", Display: "BCOL"})
	if err != nil {
		t.Errorf("unexpected error, recieved [%v]", err)
	}

}

func TestStringDataTypeGetAll(t *testing.T) {
	var err error
	var all []interface{}
	var c StringDataType[StringData]

	c = StringDataType[StringData]{}
	// test working
	err = c.Set(StringData{Key: "A"}, StringData{Key: "B", Display: "BCOL"})
	all, err = c.GetAll()
	if err != nil {
		t.Errorf("unexpected error, recieved [%v]", err)
	}
	if len(all) != 2 {
		t.Errorf("expected get to return 2 items, actual [%v]", len(all))
	}

	c = StringDataType[StringData]{}
	err = c.Set(
		StringData{Key: "A"},
		StringData{Key: "B", Display: "BCOL"},
		StringData{},
		"string",
		StringData{Key: "C"},
	)
	if err == nil {
		t.Errorf("expected Set to return with an error, recieved [%v]", err)
	}
	all, _ = c.GetAll()
	if len(all) != 3 {
		t.Errorf("expected GetAll to return the 3 valid items, actual [%v]", len(all))
	}

}

func TestStringDataTypeGet(t *testing.T) {
	var err error
	var v interface{}
	var c StringDataType[StringData]

	c = StringDataType[StringData]{}
	// set several
	err = c.Set(StringData{Key: "A"}, StringData{Key: "B", Display: "BCOL"})
	v, err = c.Get()
	if err != nil {
		t.Errorf("expected Get to return without error, recieved [%v]", err)
	}
	if v.(string) != "A" {
		t.Errorf("expected get to return key of the first value passed to set as it has no Display")
	}

	c = StringDataType[StringData]{}
	// set several false ones
	err = c.Set(
		StringData{},
		"",
		100,
	)
	v, err = c.Get()
	if v != nil {
		t.Errorf("expected get to nothing")
	}

	c = StringDataType[StringData]{}
	c.Set(StringData{Key: "test", Display: "Test"})
	v, _ = c.Get()
	if v.(string) != "Test" {
		t.Errorf("expected get to return the Display field, [%v]", v)
	}

}
