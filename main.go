package main

import (
	"fmt"
	"opg-infra-costs/internal/project"
	"opg-infra-costs/pkg/aws/accounts"
	"opg-infra-costs/pkg/aws/costs"
	"opg-infra-costs/pkg/data/csv"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/debug"
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
	defer debug.Log("Complete.")()

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month()-12, 1, 0, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, -1, now.Location())

	accountList, _ := accounts.Load(FILES["ACCOUNTS"])

	costUsageData, _ := costs.Costs(
		accountList,
		start,
		end,
	)

	costs.ToCSV(costUsageData, accountList, FILES["CSV"])

	raw := csv.ToMap(csv.Load(FILES["CSV"]))

	sheets := report.Reports(start, end, raw, FILES["FX"])

	f := excelize.NewFile()
	f.Path = FILES["XLSX"]
	f.SaveAs(f.Path)

	debug.Log("Creating worksheets")()
	d := debug.DEPTH + 2
	for _, s := range sheets {
		defer debug.LogAtDepth(fmt.Sprintf("Sheet [%s] created", s.GetName()), d)()
		s.Write(f)
		s.AddTable(f)
		s.AddPane(f)
		s.RowVisibility(f)
		f.SetSheetVisible(s.GetName(), s.GetVisible())
		f.SaveAs(f.Path)
	}
	f.DeleteSheet("Sheet1")
	f.SaveAs(f.Path)

}
