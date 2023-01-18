package row

import (
	"opg-infra-costs/pkg/cell/datedatatype"
	"opg-infra-costs/pkg/cell/floatdatatype"
	"opg-infra-costs/pkg/cell/formuladatatype"
	"testing"

	"github.com/k0kubun/pp"
)

func TestSetAndGetCells(t *testing.T) {
	r := Row{}

	r.SetVisible(true)
	r.SetHeader(false)
	r.SetIndex(2)

	testcells := [][]interface{}{
		// valid
		{100, 101.1, 0.01},
		{"200", "201.1", "0.02"},
		{"2022-02"},
		{"300", "301.1", "0.03"},
		{formuladatatype.FormulaData{Label: "A", Formula: "=SUM()"}},
	}
	err := r.SetCells(testcells)
	if err != nil {
		t.Errorf("unexpected error from SetCells, recieved [%v]", err)
	}

	all := r.GetCells()
	// -1 as there is a failing version in there
	if len(all) != len(testcells) {
		pp.Println(all)
		t.Errorf("unexpected error, GetCalls should return [%v] items, actual [%v]", len(testcells), len(all))
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
