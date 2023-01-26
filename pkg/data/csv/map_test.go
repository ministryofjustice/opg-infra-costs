package csv

import (
	"testing"
)

func TestMap(t *testing.T) {

	dummy := Row{
		Org:                "org",
		AccountId:          "1",
		AccountName:        "A",
		AccountEnvironment: "Dev",
		Service:            "AWS 1",
		Region:             "region-1",
		Date:               "2022-01",
		Cost:               "0.17",
	}

	mapped := dummy.Map()

	if mapped["Org"] != dummy.Org &&
		mapped["AccountId"] != dummy.AccountId &&
		mapped["AccountName"] != dummy.AccountName &&
		mapped["AccountEnvironment"] != dummy.AccountEnvironment &&
		mapped["Service"] != dummy.Service &&
		mapped["Region"] != dummy.Region &&
		mapped["Date"] != dummy.Date &&
		mapped["Cost"] != dummy.Cost {

		t.Errorf("mapped version does not match original\noriginal:\n%v\nmap:\n%v\n", dummy, mapped)
	}

}
