package main

import (
	"fmt"
	"opg-infra-costs/internal/project"
	"opg-infra-costs/pkg/aws/accounts"
	"opg-infra-costs/pkg/aws/costs"
	"opg-infra-costs/pkg/data/csv"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/debugger"
	"opg-infra-costs/pkg/report"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func excludeTax() (excludeTax bool) {
	excludeTax = false

	args := os.Args
	for _, v := range args {
		if strings.Contains(v, "no-tax") {
			excludeTax = true
		}
	}
	return
}

// -- Date handling

var FILES map[string]string = map[string]string{
	"CSV":      fmt.Sprintf("%s/files/costs-%s.csv", project.ROOT_DIR, time.Now().UTC().Format(dates.YM)),
	"XLSX":     fmt.Sprintf("%s/files/costs-%s.xlsx", project.ROOT_DIR, time.Now().UTC().Format(dates.YM)),
	"ACCOUNTS": project.ROOT_DIR + "/accounts.yml",
	"FX":       project.ROOT_DIR + "/exchange-rates.yml",
}

func main() {

	defer debugger.Log("Complete.", debugger.INFO)()

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month()-12, 1, 0, 0, 0, 0, now.Location())
	// TimePeriod EndDate is exclusive see:
	// https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html#awscostmanagement-GetCostAndUsage-request-TimePeriod
	// so this needs to be first of the month
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	debugger.Log(fmt.Sprintf("Report for [%s]-[%s]", start.Format(dates.YMD), end.Format(dates.YMD)), debugger.INFO)()

	accountList, _ := accounts.Load(FILES["ACCOUNTS"])

	costUsageData, _ := costs.Costs(
		accountList,
		start,
		end,
		excludeTax(),
	)

	rows, err := costs.ToCSVRows(costUsageData, accountList) //, FILES["CSV"])
	if err != nil {
		panic(err)
	}
	costs.SaveCSVRowsToFile(rows, FILES["CSV"])

	raw := csv.ToMap(csv.Load(FILES["CSV"]))
	// set the end date to be the previous complete month
	sheetEnd := end.Add(time.Duration(-1) * time.Second)
	sheets := report.Reports(start, sheetEnd, raw, FILES["FX"])

	f := excelize.NewFile()
	f.Path = FILES["XLSX"]
	f.SaveAs(f.Path)

	report.CreateWorksheets(f, sheets)

	f.DeleteSheet("Sheet1")
	f.SaveAs(f.Path)

}
