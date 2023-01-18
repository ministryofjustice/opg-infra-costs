package datedatatype

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestDateDataTypeParse(t *testing.T) {
	c := DateDataType[string]{}
	// correct date
	v, err := c.Parse("2022-01")
	if err != nil {
		t.Errorf("parsing valid date failed: %v", err)
	}
	if v.(string) != "2022-01" {
		t.Errorf("valid date did not match expected value: %v", v)
	}

	_, err = c.Parse("Not a date")
	if err == nil {
		t.Errorf("expected an error for parsing an invalid date, recieved: %v", err)
	}

}
func TestDateDataTypeSet(t *testing.T) {

	c := DateDataType[string]{}
	// correct date
	err := c.Set("2022-01")
	if err != nil {
		t.Errorf("unexpected error setting a date: %v", err)
	}
	// these should fail as they are not the correct format
	err = c.Set("2022")
	if err == nil {
		t.Errorf("expected an error, recieved: %v", err)
	}
	err = c.Set("2022-01-23")
	if err == nil {
		t.Errorf("expected an error, recieved: %v", err)
	}
	// this should fail as its not a date
	err = c.Set("Not a date at all")
	if err == nil {
		t.Errorf("expected an error, recieved: %v", err)
	}
	// test setting multiple that work
	err = c.Set("2022-01", "2022-02")
	if err != nil {
		t.Errorf("unexpected error setting a set of dates: %v", err)
	}

	// test setting multiple with a failure
	err = c.Set("2022-01", "2022-02", "not a date")
	if err == nil {
		t.Errorf("expected error with one of the passed dates failed, recieved: %v", err)
	}
	err = c.Set("2022-01", "2022", "2022-01")
	if err == nil {
		t.Errorf("expected Set to return with an error, recieved [%v]", err)
	}

}

func TestDateDataTypeGetAll(t *testing.T) {
	c := DateDataType[string]{}
	// set many, getall should match
	c.Set("2022-01", "2022-06")

	all, err := c.GetAll()
	if len(all) != 2 {
		t.Errorf("expected GetAll to return 2 items, actual [%v]", len(all))
		pp.Println(all)
	}
	if err != nil {
		t.Errorf("expected GetAll to return without error, recieved [%v]", err)
	}
	// include some bad data in the middle
	c = DateDataType[string]{}

	err = c.Set("2022-01", "2022", "2022-01")
	if err == nil {
		t.Errorf("expected Set to return with an error, recieved [%v]", err)
	}

	all, _ = c.GetAll()
	if len(all) != 2 {
		t.Errorf("expected GetAll to return 2 items, actual [%v]", len(all))
		pp.Println(all)
	}

}

func TestDateDataTypeGet(t *testing.T) {

	c := DateDataType[string]{}
	// set several
	c.Set("2022-01", "2022-06", "2022-02")
	// get should return just the first
	v, e := c.Get()
	if e != nil {
		t.Errorf("expected Get to return without error, recieved [%v]", e)
	}
	if v != "2022-01" {
		t.Errorf("expected get to return the first value passed to set")
	}

	c = DateDataType[string]{}
	// set several
	c.Set("not a date")
	// get should return just the first
	v, e = c.Get()
	if e != nil {
		t.Errorf("expected Get to return without error even when empty, recieved [%v]", e)
	}
	if v != nil {
		t.Errorf("expected Get to return empty string, recieved [%v]", v)
	}

}
