package main

import (
	"fmt"
	"opg-infra-costs/internal/project"
	"opg-infra-costs/pkg/aws/accounts"
	"opg-infra-costs/pkg/aws/costs"
	"opg-infra-costs/pkg/data/csv"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/out"
	"opg-infra-costs/pkg/report"
	"time"

	"github.com/xuri/excelize/v2"
)

// -- Date handling

var FILES map[string]string = map[string]string{
	"CSV":      fmt.Sprintf("%s/files/costs-%s.csv", project.ROOT_DIR, time.Now().UTC().Format(dates.YM)),
	"XLSX":     fmt.Sprintf("%s/files/costs-%s.xlsx", project.ROOT_DIR, time.Now().UTC().Format(dates.YM)),
	"ACCOUNTS": project.ROOT_DIR + "/accounts.yml",
	"FX":       project.ROOT_DIR + "/exchange-rates.yml",
}

func main() {
	var duration time.Duration

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month()-12, 1, 0, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, -1, now.Location())

	out.CLI(
		fmt.Sprintf("Report period: [%s to %s]", start.Format(dates.YMD), end.Format(dates.YMD)),
		-1)

	accountList, duration, _ := accounts.Load(FILES["ACCOUNTS"])
	out.CLI("- Loaded Accounts", duration)

	costUsageData, duration, _ := costs.Costs(
		accountList,
		start,
		end,
	)
	out.CLI("- Fetched cost data", duration)

	duration, _ = costs.ToCSV(costUsageData, accountList, FILES["CSV"])
	out.CLI("- Wrote cost data to csv", duration)

	raw, duration := csv.ToMap(csv.Load(FILES["CSV"]))
	out.CLI("- Converted csv data to map", duration)

	sheets, duration := report.Reports(start, end, raw, FILES["FX"])
	out.CLI("- Converted data & created report configurations", duration)

	f := excelize.NewFile()
	f.Path = FILES["XLSX"]
	f.SaveAs(f.Path)

	marker := time.Now().UTC()
	for _, s := range sheets {
		s.Write(f)
		s.AddTable(f)
		s.AddPane(f)
		s.RowVisibility(f)
		f.SetSheetVisible(s.GetName(), s.GetVisible())
		f.SaveAs(f.Path)
		duration = time.Since(marker)
		marker = time.Now().UTC()
		out.CLI(fmt.Sprintf("- [%s] Worksheet saved", s.GetName()), duration)
	}
	f.DeleteSheet("Sheet1")
	f.SaveAs(f.Path)

	duration = time.Since(now)
	out.CLI("", 0)
	out.CLI("Completed.", duration)
	out.CLI("=====", 0)
}
