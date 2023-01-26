package main

import (
	"fmt"
	"opg-infra-costs/internal/project"
	"opg-infra-costs/pkg/data/csv"
	"opg-infra-costs/pkg/report"
	"time"

	"github.com/xuri/excelize/v2"
)

// -- Date handling
const DATEFORMAT string = "2006-01"

var FILES map[string]string = map[string]string{
	"CSV":  project.ROOT_DIR + "/files/costs.csv",
	"XLSX": fmt.Sprintf("%s/files/costs-%s.xlsx", project.ROOT_DIR, time.Now().UTC().Format(DATEFORMAT)),
}

func main() {
	start, _ := time.Parse(DATEFORMAT, "2022-01")
	end, _ := time.Parse(DATEFORMAT, "2022-12")

	raw := csv.ToMap(csv.Load(FILES["CSV"]))
	sheets := report.Reports(start, end, raw)

	f := excelize.NewFile()
	f.Path = FILES["XLSX"]
	f.SaveAs(f.Path)

	for _, s := range sheets {
		s.Write(f)
		s.AddTable(f)
		s.AddPane(f)
		f.SetSheetVisible(s.GetName(), s.GetVisible())
		f.SaveAs(f.Path)

	}
	f.DeleteSheet("Sheet1")
	f.SaveAs(f.Path)

}
