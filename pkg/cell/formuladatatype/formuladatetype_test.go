package formuladatatype

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

	//pp.Println(res)
	pft = res.(FormulaData)
	if pft.Label != "A" || pft.Formula != "=SUM()" {
		t.Errorf("expected Parse to return a matching FormulaData from json, recieved [%v]", pft)
	}

}
