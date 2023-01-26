package convert

import (
	"testing"
)

func TestSkeleton(t *testing.T) {

	set := []map[string]string{
		{"AccountId": "1", "AccountName": "A", "Service": "AWS 1", "Region": "region1", "Date": "2022-01", "Cost": "5.17"},
		{"AccountId": "1", "AccountName": "A", "Service": "AWS 1", "Region": "region1", "Date": "2022-02", "Cost": "10.13"},
		{"AccountId": "2", "AccountName": "B", "Service": "AWS 1", "Region": "region1", "Date": "2022-02", "Cost": "13.01"},
	}

	groupings := []string{"AccountId", "AccountName"}
	dateCost := map[string]string{"Date": "Cost"}
	others := []string{"Totals"}

	skel := Skeleton(set, groupings, dateCost, others)
	eLen := 2
	if len(skel) != eLen {
		t.Errorf("expected to find [%d] unique rows, actual [%v]", eLen, len(skel))
	}

}
