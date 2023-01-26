package report

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestSheetSimple(t *testing.T) {

	data := map[string]map[string][]string{
		"field-key-1": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"2022-01":            []string{"11.35", "-1.03"},
		},
		"field-key-2": {
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

	file := excelize.NewFile()

	s.Write(file)

	c2 := s.cells[CellRef{Row: 2, Col: 3}]
	if c2.ValueType != DataIsANumber {
		t.Errorf("expected c2 to be a float [%v]", c2)
	}
	if c2.Value.(float64) != 10.32 {
		t.Errorf("unexpected value [%v]", c2.Value)
	}
	a3 := s.cells[CellRef{Row: 3, Col: 1}]
	if a3.ValueType != DataIsAString {
		t.Errorf("expected to be a string [%v]", a3)
	}
	if a3.Value.(string) != "Test2" {
		t.Errorf("unexpected value [%v]", a3.Value)
	}
}

func TestSheetWithCostChanges(t *testing.T) {

	data := map[string]map[string][]string{
		"field-key-1": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 1", "Service 1"},
			"2022-01":            []string{"11.35", "-1.03"},
			"2022-02":            []string{"13.54"},
			"2022-03":            []string{"15.78"},
			"2022-04":            []string{"-9.99"},
		},
		"field-key-2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 2", "Service 2"},
			"2022-01":            []string{"101.12", "-10.11"},
			"2022-02":            []string{"103.47"},
			"2022-03":            []string{"206.94"},
			"2022-04":            []string{"-109.02"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "Service", Display: "Service"},
		{MapKey: "Region", Display: "Region"},
		{MapKey: "2022-02", Display: "2022-02"},
		{MapKey: "2022-03", Display: "2022-03"},
		{MapKey: "", Display: "Increase ($)", Formula: "=(F${r}-E${r})"},
		{MapKey: "", Display: "Increase (%)", Formula: "=(F${r}/E${r})-1"},
	}

	file := excelize.NewFile()

	s := Sheet{}
	s.Init()
	s.SetName("testchanges")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)

	s.AddStyle(&excelize.Style{
		NumFmt: 10,
	}, 0, 8)

	s.Write(file)

	s.AddTable(file)
	s.AddPane(file)

	//file.SaveAs(project.ROOT_DIR + "/files/costs-testcostchanges.xlsx")

}

func TestSheetWithFormula(t *testing.T) {

	data := map[string]map[string][]string{
		"field-key-1": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 1", "Service 1"},
			"2022-01":            []string{"11.35", "-1.03"},
			"2022-02":            []string{"13.54"},
			"2022-03":            []string{"15.78"},
			"2022-04":            []string{"-9.99"},
		},
		"field-key-2": {
			"AccountName":        []string{"Test1"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 2", "Service 2"},
			"2022-01":            []string{"101.12", "-10.11"},
			"2022-02":            []string{"103.47"},
			"2022-03":            []string{"105.17"},
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

	file := excelize.NewFile()

	s := Sheet{}
	s.Init()
	s.SetName("test2")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)
	s.Write(file)

	totals := s.cells[CellRef{Row: 1, Col: 9}]
	if totals.Value.(string) != "Totals" {
		t.Errorf("unexpected cell [%v]", totals)
	}

	name := s.cells[CellRef{Row: 1, Col: 1}]
	if name.Value.(string) != "Account" {
		t.Errorf("unexpected cell [%v]", name)
	}

	sum := s.cells[CellRef{Row: 2, Col: 9}]
	if sum.Value.(string) != "=SUM(E2:H2)" {
		t.Errorf("unexpected formula value [%v]", sum)
	}

}
