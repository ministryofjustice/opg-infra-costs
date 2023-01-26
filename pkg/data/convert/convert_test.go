package convert

import (
	"testing"
)

func TestConvert(t *testing.T) {

	set := []map[string]string{
		{"Org": "", "AccountId": "1", "AccountName": "A", "Service": "AWS 1", "Region": "region1", "Date": "2022-01", "Cost": "5.17"},
		{"Org": "", "AccountId": "1", "AccountName": "A", "Service": "AWS 2", "Region": "region2", "Date": "2022-02", "Cost": "10.13"},
		{"Org": "", "AccountId": "1", "AccountName": "A", "Service": "AWS 3", "Region": "region3", "Date": "2022-02", "Cost": "11.01"},
		{"Org": "", "AccountId": "1", "AccountName": "A", "Service": "AWS 1", "Region": "region4", "Date": "2022-02", "Cost": "9.01"},
		{"Org": "", "AccountId": "2", "AccountName": "B", "Service": "AWS 1", "Region": "region5", "Date": "2022-02", "Cost": "13.01"},
		{"Org": "", "AccountId": "2", "AccountName": "B", "Service": "AWS 1", "Region": "region6", "Date": "2022-03", "Cost": "1.05"},
		{"Org": "", "AccountId": "3", "AccountName": "C", "Service": "AWS 3", "Region": "region7", "Date": "2022-02", "Cost": "15.47"},
		{"Org": "", "AccountId": "3", "AccountName": "C", "Service": "AWS 3", "Region": "region8", "Date": "2022-02", "Cost": "1.07"},
		{"Org": "", "AccountId": "3", "AccountName": "C", "Service": "AWS 3", "Region": "region9", "Date": "2022-02", "Cost": "0.01"},
	}

	// typical to find totals of all per month

	result := Convert(
		set,
		[]string{"Org"},
		map[string]string{"Date": "Cost"},
		[]string{},
	)
	expectedL := 1
	if expectedL != len(result) {
		t.Errorf("expected to find [%d] unique row, actual [%v]", expectedL, len(result))
	}

	// check the cell merge works
	month := result["{\"Org\":\"\"}"]["2022-02"]
	if len(month) != 7 {
		t.Errorf("grouping failed, expected [7] enteries, actual [%d]", len(month))
	}

	// typical to find totals of all per month per service team
	result = Convert(
		set,
		[]string{"Org", "AccountId", "AccountName"},
		map[string]string{"Date": "Cost"},
		[]string{},
	)
	expectedL = 3
	if expectedL != len(result) {
		t.Errorf("expected to find [%d] unique row, actual [%v]", expectedL, len(result))
	}

	// now tweak the convertor to split by more
	result = Convert(
		set,
		[]string{"Org", "AccountId", "AccountName", "Service"},
		map[string]string{"Date": "Cost"},
		[]string{},
	)
	expectedL = 5
	if expectedL != len(result) {
		t.Errorf("expected to find [%d] unique row, actual [%v]", expectedL, len(result))
	}

	result = Convert(
		set,
		[]string{"Org", "AccountId", "AccountName", "Service", "Region"},
		map[string]string{"Date": "Cost"},
		[]string{},
	)
	expectedL = 9 // all regions are unique
	if expectedL != len(result) {
		t.Errorf("expected to find [%d] unique row, actual [%v]", expectedL, len(result))
	}
}
