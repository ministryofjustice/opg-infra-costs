package report

import (
	"fmt"
	"opg-infra-costs/pkg/debugger"

	"github.com/xuri/excelize/v2"
)

func CreateWorksheets(
	f *excelize.File,
	sheets []Sheet,
) {
	debugger.Log("Creating worksheets", debugger.DETAILED)()
	for _, s := range sheets {
		defer debugger.Log(fmt.Sprintf("Sheet [%s] created", s.GetName()), debugger.VERBOSE)()
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
