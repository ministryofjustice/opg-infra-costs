package cell

import (
	"encoding/json"
	"testing"
)

func TestFormulaDataTypeParse(t *testing.T) {
	var err error

	// check with the standard type
	ft := FormulaData{Label: "A", Formula: "=SUM()"}

	c := FormulaDataType[FormulaData]{}
	p, err := c.Parse(ft)
	if err != nil {
		t.Errorf("unexpected error from parsing FormulaData, [%v]", err)
	}

	pft := p.(FormulaData)
	if pft.Label != "A" || pft.Formula != "=SUM()" {
		t.Errorf("expected Parse to return a matching FormulaData, recieved [%v]", pft)
	}

	// check json/string parsing
	jsB, _ := json.Marshal(ft)
	jsF := string(jsB)
	jsC := FormulaDataType[string]{}

	res, err := jsC.Parse(jsF)

	if err != nil {
		t.Errorf("unexpected error from parsing FormulaData as json, [%v]", err)
	}
	pft = res.(FormulaData)
	if pft.Label != "A" || pft.Formula != "=SUM()" {
		t.Errorf("expected Parse to return a matching FormulaData from json, recieved [%v]", pft)
	}

}

func TestFormulaDataTypeSet(t *testing.T) {
	var err error

	// test sets with the struct
	c := FormulaDataType[FormulaData]{}
	err = c.Set(
		FormulaData{Label: "A", Formula: "=SUM()"},
		FormulaData{Label: "B", Formula: "=SUM()"},
	)
	if err != nil {
		t.Errorf("unexpected error returned by Set: [%v]", err)
	}

	// this should throw an error as it should have a Label
	err = c.Set(
		FormulaData{Formula: "=SUM()"},
	)
	if err == nil {
		t.Errorf("expected an error to be returned by Set, recieved: [%v]", err)
	}

	err = c.Set(
		FormulaData{Label: "A", Formula: ""},
	)
	if err == nil {
		t.Errorf("expected an error to be returned by Set, recieved: [%v]", err)
	}

}

func TestFormulaDataTypeGetAll(t *testing.T) {
	c := FormulaDataType[FormulaData]{}
	c.Set(
		FormulaData{},
		FormulaData{Label: "1"},
		FormulaData{Label: "A", Formula: "=SUM()"},
		FormulaData{Label: "B", Formula: "=SUM()"},
	)

	vals, _ := c.GetAll()
	if len(vals) != 2 {
		t.Errorf("expected GetAll to only return 2 items, actual [%v]", len(vals))
	}
}

func TestFormulaDataTypeGet(t *testing.T) {
	// these are body rows, so should return formula
	c := FormulaDataType[FormulaData]{}
	c.Set(
		FormulaData{},
		FormulaData{Label: "1"},
		FormulaData{Label: "A", Formula: "=SUMA()"},
		FormulaData{Label: "B", Formula: "=SUMB()"},
	)

	val, _ := c.Get()
	if val.(string) != "=SUMA()" {
		t.Errorf("expected Get to return matching formula, recieved: [%v]", val)
	}

	c = FormulaDataType[FormulaData]{}
	c.Set(FormulaData{Label: "TOTALS", Formula: "=SUM()"})
	c.SetIsRowAHeader(true)

	val, _ = c.Get()

	if val.(string) != "TOTALS" {
		t.Errorf("expected Get to return matching formula, recieved: [%v]", val)
	}

}
