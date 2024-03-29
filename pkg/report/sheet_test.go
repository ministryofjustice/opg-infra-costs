package report

import (
	"fmt"
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
	keyed, _ := s.SetDataset(data)

	c2, _ := s.Cell(
		keyed["r2"].Index,
		keyed["r2"].Columns["2022-01"],
	)
	if c2.ValueType != DataIsANumber {
		t.Errorf("expected c2 to be a float [%v]", c2)
	}
	if c2.Value.(float64) != 10.32 {
		t.Errorf("unexpected value [%v]", c2.Value)
	}
	a3, _ := s.Cell(
		keyed["r3"].Index,
		keyed["r3"].Columns["AccountName"],
	)
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
			"AccountId":          []string{"12309817212"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 1", "Service 1"},
			"2022-01":            []string{"11.35"},
			"2022-02":            []string{"13.54"},
		},
		"r3": {
			"AccountName":        []string{"Test1"},
			"AccountId":          []string{"12309817212"},
			"AccountEnvironment": []string{"Production", "Production"},
			"Service":            []string{"Servce 2", "Service 2"},
			"2022-01":            []string{"101.12", "-10.11"},
			"2022-04":            []string{"-109.02"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "AccountId", Display: "AccountId", ForceColumnToDisplayAsString: true},
		{MapKey: "AccountEnvironment", Display: "Environment"},
		{MapKey: "Service", Display: "Service"},
		{MapKey: "Region", Display: "Region"},
		{MapKey: "2022-01", Display: "2022-01"},
		{MapKey: "2022-02", Display: "2022-02"},
		{MapKey: "2022-03", Display: "2022-03"},
		{MapKey: "2022-04", Display: "2022-04"},
		{MapKey: "Totals", Display: "Totals", Formula: "=SUM(E${r}:H${r})"},
		{MapKey: "Trend", Display: "Trend", Formula: "=SPARKLINE(E${r}:H${r}, {\"charttype\",\"column\";\"empty\",\"ignore\";\"nan\",\"convert\"})"},
	}

	s := Sheet{}
	s.Init()
	s.SetName("test2")
	s.SetColumns(headers, ColumnsAreOther)
	keyed, _ := s.SetDataset(data)

	totals, _ := s.Cell(
		keyed["Header"].Index,
		keyed["Header"].Columns["Totals"],
	)

	if totals.Value.(string) != "Totals" {
		t.Errorf("unexpected cell [%v]", totals)
	}

	name, _ := s.Cell(
		keyed["Header"].Index,
		keyed["Header"].Columns["AccountName"],
	)
	if name.Value.(string) != "Account" {
		t.Errorf("unexpected cell [%v]", name)
	}

	file := excelize.NewFile()
	s.Write(file)

	sum, _ := s.Cell(
		keyed["r2"].Index,
		keyed["r2"].Columns["Totals"],
	)
	f := fmt.Sprintf("=SUM(E%d:H%d)", keyed["r2"].Index, keyed["r2"].Index)
	if sum.Value.(string) != f {
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

	// test to make sure account id comes out as a string for excel
	tc, _ := s.Cell(keyed["r2"].Index, keyed["r2"].Columns["AccountId"])
	if tc.Value != `"12309817212"` {
		t.Errorf("expected account id to be a string encapsulated version, recieved something else")
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
		{MapKey: "Diff", Display: "Increase ($)", Formula: "=(E${r}-D${r})"},
		{MapKey: "Percent", Display: "Increase (%)", Formula: "=(E${r}/D${r})-1"},
	}

	s := Sheet{}
	s.Init()
	s.SetName("testchanges")
	s.SetColumns(headers, ColumnsAreOther)
	keyed, _ := s.SetDataset(data)

	file := excelize.NewFile()
	s.Write(file)
	// test the diff col matches what it should be
	diff, _ := s.Cell(
		keyed["r3"].Index,
		keyed["r3"].Columns["Diff"],
	)
	expectedDiff := 103.47
	val, _ := diff.CalculatedValue(file, s.GetName())
	actual, _ := strconv.ParseFloat(val, 64)
	if actual != expectedDiff {
		t.Errorf("expected diff to be [%v], recieved [%v]", expectedDiff, actual)
	}

	// test % col, should 100% increase, so 1
	per, _ := s.Cell(
		keyed["r3"].Index,
		keyed["r3"].Columns["Percent"],
	)
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
			"Service":     []string{"Test1"},
			"2022-01":     []string{"11.35"},
		},
		"r2": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Test1"},
			"2022-01":     []string{"11.37"},
		},
		"r3": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Test1"},
			"2022-01":     []string{"-1.01"},
		},
		"r4": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Tax"},
			"2022-01":     []string{"12.17"},
		},
		// the comparison is done with ABS, so this will be returned
		"r5": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Test1"},
			"2022-01":     []string{"-19.17"},
		},
		"r6": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Test1"},
			"2022-01":     []string{"9.19"},
		},
		"r7": {
			"AccountName": []string{"Test2"},
			"Service":     []string{"Refund"},
			"2022-01":     []string{"12.17"},
		},
	}

	headers := []Column{
		{MapKey: "AccountName", Display: "Account"},
		{MapKey: "Service", Display: "Service"},
		{MapKey: "2022-01", Display: "2022-01"},
	}

	s := NewSheet("test2")
	s.SetColumns(headers, ColumnsAreOther)
	s.SetDataset(data)
	file := excelize.NewFile()
	s.Write(file)

	// hide rows with values less than 10
	hide := map[CellRef]interface{}{{Col: 3}: 10}
	s.SetHideRowWhen(hide)
	hidden, _ := s.RowVisibility(file)
	if len(hidden) != 2 {
		t.Errorf("expected 2 rows to be hidden, actual [%v]", len(hidden))
	}
	// hide rows with values less than 10
	hide = map[CellRef]interface{}{
		{Col: 2, Row: 0}:  "Tax",
		{Col: 2, Row: -1}: "Refund",
	}
	s.SetHideRowWhen(hide)
	hidden, _ = s.RowVisibility(file)
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
	keyed, _ := s.SetDataset(data)

	c2, _ := s.Cell(
		keyed["r2"].Index,
		keyed["r2"].Columns["2022-01"],
	)
	if c2.ValueType != DataIsANumber {
		t.Errorf("expected c2 to be a float [%v]", c2)
	}
	if c2.Value.(float64) != 10.32 {
		t.Errorf("unexpected value [%v]", c2.Value)
	}

	a3, _ := s.Cell(
		keyed["r3"].Index,
		keyed["r3"].Columns["AccountName"],
	)
	if a3.ValueType != DataIsAString {
		t.Errorf("expected to be a string [%v]", a3)
	}
	if a3.Value.(string) != "Test2" {
		t.Errorf("unexpected value [%v]", a3.Value)
	}

	// add a new row to the data
	keyed, _ = s.AddRow("add-a-row-test", map[string][]string{
		"AccountName":        {"Test3"},
		"AccountEnvironment": {"Dev"},
		"2022-01":            {"1000.13"},
	})

	c4, _ := s.Cell(
		keyed["add-a-row-test"].Index,
		keyed["add-a-row-test"].Columns["2022-01"],
	)
	if c4.Value.(float64) != 1000.13 {
		t.Errorf("unexpected value [%v]", c4.Value)
	}

	s.AddCell(2, 1, "NewAccount")
	cA, _ := s.Cell(2, 1)
	if cA.Value.(string) != "NewAccount" {
		t.Errorf("Cell overwrite did not work [%v]", cA.Value)
	}

}
