package report

import (
	"strconv"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestSheetSimple(t *testing.T) {

	data := map[string]map[string][]string{
		"r2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"2022-01":            []string{"11.35", "-1.03"},
		},
		"r3": {
			"AccountName":        []string{"Test2"},
			"AccountEnvironment": []string{"Production", "Production"},
			"2022-01":            []string{"11.37", "-9.03"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "2022-01", Display: "2022-01"},
	}

	s := NewSheet("test1")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)

	c2, _ := s.Cell(2, 3)
	if c2.ValueType != DataIsANumber {
		t.Errorf("expected c2 to be a float [%v]", c2)
	}
	if c2.Value.(float64) != 10.32 {
		t.Errorf("unexpected value [%v]", c2.Value)
	}
	a3, _ := s.Cell(3, 1)
	if a3.ValueType != DataIsAString {
		t.Errorf("expected to be a string [%v]", a3)
	}
	if a3.Value.(string) != "Test2" {
		t.Errorf("unexpected value [%v]", a3.Value)
	}
}

func TestSheetWithFormula(t *testing.T) {

	data := map[string]map[string][]string{
		"r2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 1", "Service 1"},
			"2022-01":            []string{"11.35"},
			"2022-02":            []string{"13.54"},
		},
		"r3": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 2", "Service 2"},
			"2022-01":            []string{"101.12", "-10.11"},
			"2022-04":            []string{"-109.02"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "Service", Display: "Service"},
		{MapKey: "Region", Display: "Region"},
		{MapKey: "2022-01", Display: "2022-01"},
		{MapKey: "2022-02", Display: "2022-02"},
		{MapKey: "2022-03", Display: "2022-03"},
		{MapKey: "2022-04", Display: "2022-04"},
		{MapKey: "", Display: "Totals", Formula: "=SUM(E${r}:H${r})"},
		{MapKey: "", Display: "Trend", Formula: "=SPARKLINE(E${r}:H${r}, {\"charttype\",\"column\";\"empty\",\"ignore\";\"nan\",\"convert\"})"},
	}

	s := Sheet{}
	s.Init()
	s.SetName("test2")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)

	file := excelize.NewFile()
	s.Write(file)

	totals, _ := s.Cell(1, 9)
	if totals.Value.(string) != "Totals" {
		t.Errorf("unexpected cell [%v]", totals)
	}

	name, _ := s.Cell(1, 1)
	if name.Value.(string) != "Account" {
		t.Errorf("unexpected cell [%v]", name)
	}

	sum, _ := s.Cell(2, 9)
	if sum.Value.(string) != "=SUM(E2:H2)" {
		t.Errorf("unexpected formula value [%v]", sum.Value)
	}
	a, _ := strconv.ParseFloat(data["r2"]["2022-01"][0], 64)
	b, _ := strconv.ParseFloat(data["r2"]["2022-02"][0], 64)
	expected := a + b
	val, _ := sum.CalculatedValue(file, s.GetName())
	actual, _ := strconv.ParseFloat(val, 64)

	if actual != expected {
		t.Errorf("expected total to be [%v], recieved [%v]", expected, actual)
	}

}

func TestSheetWithCostChanges(t *testing.T) {

	data := map[string]map[string][]string{
		"r2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 1", "Service 1"},
			"2022-02":            []string{"13.54"},
			"2022-03":            []string{"15.78"},
		},
		"r3": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 2", "Service 2"},
			"2022-02":            []string{"103.47"},
			"2022-03":            []string{"206.94"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "Service", Display: "Service"},
		{MapKey: "2022-02", Display: "2022-02"},
		{MapKey: "2022-03", Display: "2022-03"},
		{MapKey: "", Display: "Increase ($)", Formula: "=(E${r}-D${r})"},
		{MapKey: "", Display: "Increase (%)", Formula: "=(E${r}/D${r})-1"},
	}

	s := Sheet{}
	s.Init()
	s.SetName("testchanges")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)
	// % last column
	percentCol := 7
	diffCol := 6

	file := excelize.NewFile()
	s.Write(file)
	// test the diff col matches what it should be
	diff, _ := s.Cell(3, diffCol)
	expectedDiff := 103.47
	val, _ := diff.CalculatedValue(file, s.GetName())
	actual, _ := strconv.ParseFloat(val, 64)
	if actual != expectedDiff {
		t.Errorf("expected diff to be [%v], recieved [%v]", expectedDiff, actual)
	}

	// test % col, should 100% increase, so 1
	per, _ := s.Cell(3, percentCol)
	val, _ = per.CalculatedValue(file, s.GetName())
	expectedP := "1"
	if val != expectedP {
		t.Errorf("expected percent to be [%v], recieved [%v]", expectedP, val)
	}
}

func TestSheetHideRows(t *testing.T) {

	data := map[string]map[string][]string{
		"r1": {
			"AccountName": []string{"Test1"},
			"2022-01":     []string{"11.35"},
		},
		"r2": {
			"AccountName": []string{"Test2"},
			"2022-01":     []string{"11.37"},
		},
		"r3": {
			"AccountName": []string{"Test2"},
			"2022-01":     []string{"-1.01"},
		},
		"r4": {
			"AccountName": []string{"Test2"},
			"2022-01":     []string{"12.17"},
		},
		// the comparison is done with ABS, so this will be returned
		"r5": {
			"AccountName": []string{"Test2"},
			"2022-01":     []string{"-19.17"},
		},
		"r6": {
			"AccountName": []string{"Test2"},
			"2022-01":     []string{"9.19"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "2022-01", Display: "2022-01"},
	}

	s := NewSheet("test2")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)
	// hide rows with values less than 10
	hide := map[CellRef]float64{{Col: 2}: 10}
	s.SetHideRowWhen(hide)

	file := excelize.NewFile()
	s.Write(file)
	hidden, _ := s.RowVisibility(file)

	if len(hidden) != 2 {
		t.Errorf("expected 2 rows to be hidden, actual [%v]", len(hidden))
	}

}

func TestSheetSetDataset(t *testing.T) {

	data := map[string]map[string][]string{
		"r2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"2022-01":            []string{"11.35", "-1.03"},
		},
		"r3": {
			"AccountName":        []string{"Test2"},
			"AccountEnvironment": []string{"Production", "Production"},
			"2022-01":            []string{"11.37", "-9.03"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "2022-01", Display: "2022-01"},
	}

	s := NewSheet("testing-set")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)

	c2, _ := s.Cell(2, 3)
	if c2.ValueType != DataIsANumber {
		t.Errorf("expected c2 to be a float [%v]", c2)
	}
	// if c2.Value.(float64) != 10.32 {
	// 	t.Errorf("unexpected value [%v]", c2.Value)
	// }
	// a3 := s.cells[CellRef{Row: 3, Col: 1}]
	// if a3.ValueType != DataIsAString {
	// 	t.Errorf("expected to be a string [%v]", a3)
	// }
	// if a3.Value.(string) != "Test2" {
	// 	t.Errorf("unexpected value [%v]", a3.Value)
	// }
}
