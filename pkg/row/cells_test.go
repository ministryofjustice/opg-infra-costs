package row

import (
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
	"opg-infra-costs/pkg/cell/formuladatatype"
	"testing"
)

func TestSetAndGetCells(t *testing.T) {
	r := Row{}

	r.SetVisible(true)
	r.SetHeader(true)
	r.SetIndex(1)

	testcells := [][]interface{}{
		// valid
		{100, 101.1, 0.01},
		{"200", "201.1", "0.02"},
		{"2022-02"},
		{"300", "301.1", "0.03"},
		{formuladatatype.FormulaData{Label: "A", Formula: "=SUM()"}},
		{true, false},
	}
	err := r.SetCells(testcells)
	if err == nil {
		t.Errorf("expected error from SetCells due to invalid data, recieved [%v]", err)
	}

	all := r.GetCells()
	// -1 as there is a failing version in there
	if len(all) != len(testcells)-1 {
		t.Errorf("unexpected error, GetCalls should return [%v] items, actual [%v]", len(testcells), len(all))
	}

	// check the Get on a formula returns its label (as its a header)
	h := all[4]
	hV, _ := h.Get()
	hAll, _ := h.GetAll()
	if hV != hAll[0].(formuladatatype.FormulaData).Label {
		t.Errorf("expected formula cells to return its label, actual [%T] [%v]", h, hV)
	}

	// now check some cell types
	// - float
	ty1 := &floatdatatype.FloatDataType[float64]{}
	fl := all[0]
	if fl.Type() != ty1.Type() {
		t.Errorf("epxected this cell to be of type [%T], actual [%T]", ty1, fl)
	}
	if v, _ := fl.Get(); v != 201.11 {
		t.Errorf("expected Get to return sum of values [201.11], actual [%v]", v)
	}

	// - date
	ty2 := &datedatatype.DateDataType[string]{}
	fl = all[2]
	if fl.Type() != ty2.Type() {
		t.Errorf("epxected this cell to be of type [%T], actual [%T]", ty2, fl)
	}
	if v, _ := fl.Get(); v != "2022-02" {
		t.Errorf("expected Get to return [2022-02], actual [%v]", v)
	}

	// bools - should fail
	testcells = [][]interface{}{
		{true, false},
	}
	err = r.SetCells(testcells)
	if err == nil {
		t.Errorf("expected SetCells to return an error for unsupported type, recieved [%v]", err)
	}

}
