package report

import (
	"fmt"
	"opg-infra-costs/pkg/data/convert"
	"opg-infra-costs/pkg/dates"
	"time"

	"github.com/xuri/excelize/v2"
)

func Reports(
	start time.Time,
	end time.Time,
	rawDataset []map[string]string,
) (sheets []Sheet) {
	var name string

	// generate the date headers
	dateHeaders := []Column{}
	for _, d := range dates.Months(start, end, DATEFORMAT) {
		dateHeaders = append(dateHeaders, Column{MapKey: d, Display: d})
	}
	pre := preDateHeaders(dateHeaders)
	post := postDateHeaders(dateHeaders, pre)

	// -- totals
	name = "Totals"
	totals := NewSheet(name)
	totals.SetColumns(pre[name], ColumnsAreGroupBy)
	totals.SetColumns(dateHeaders, ColumnsAreDateCost)
	totals.SetColumns(post[name], ColumnsAreOther)
	totals.SetDataset(
		convert.Convert(
			rawDataset,
			totals.GetGroupColumns(),
			totals.GetDateCostColumns(),
			totals.GetOtherColumns(),
		),
	)
	sheets = append(sheets, totals)
	// -- Service
	name = "Service"
	service := NewSheet(name)
	service.SetColumns(pre[name], ColumnsAreGroupBy)
	service.SetColumns(dateHeaders, ColumnsAreDateCost)
	service.SetColumns(post[name], ColumnsAreOther)
	service.SetDataset(
		convert.Convert(
			rawDataset,
			service.GetGroupColumns(),
			service.GetDateCostColumns(),
			service.GetOtherColumns(),
		),
	)
	sheets = append(sheets, service)
	// -- Service And Environment
	name = "Service And Environment"
	serviceAndEnvironment := NewSheet(name)
	serviceAndEnvironment.SetColumns(pre[name], ColumnsAreGroupBy)
	serviceAndEnvironment.SetColumns(dateHeaders, ColumnsAreDateCost)
	serviceAndEnvironment.SetColumns(post[name], ColumnsAreOther)
	serviceAndEnvironment.SetDataset(
		convert.Convert(
			rawDataset,
			serviceAndEnvironment.GetGroupColumns(),
			serviceAndEnvironment.GetDateCostColumns(),
			serviceAndEnvironment.GetOtherColumns(),
		),
	)
	sheets = append(sheets, serviceAndEnvironment)
	// -- Detailed Breakdown
	name = "Detailed Breakdown"
	detailedBreakdown := NewSheet(name)
	detailedBreakdown.SetColumns(pre[name], ColumnsAreGroupBy)
	detailedBreakdown.SetColumns(dateHeaders, ColumnsAreDateCost)
	detailedBreakdown.SetColumns(post[name], ColumnsAreOther)
	detailedBreakdown.SetDataset(
		convert.Convert(
			rawDataset,
			detailedBreakdown.GetGroupColumns(),
			detailedBreakdown.GetDateCostColumns(),
			detailedBreakdown.GetOtherColumns(),
		),
	)
	sheets = append(sheets, detailedBreakdown)
	// -- Detailed Breakdown With Region
	name = "Detailed Breakdown With Region"
	detailedBreakdownWithRegion := NewSheet(name)
	detailedBreakdownWithRegion.SetVisible(false)
	detailedBreakdownWithRegion.SetColumns(pre[name], ColumnsAreGroupBy)
	detailedBreakdownWithRegion.SetColumns(dateHeaders, ColumnsAreDateCost)
	detailedBreakdownWithRegion.SetColumns(post[name], ColumnsAreOther)
	detailedBreakdownWithRegion.SetDataset(
		convert.Convert(
			rawDataset,
			detailedBreakdownWithRegion.GetGroupColumns(),
			detailedBreakdownWithRegion.GetDateCostColumns(),
			detailedBreakdownWithRegion.GetOtherColumns(),
		),
	)
	sheets = append(sheets, detailedBreakdownWithRegion)

	return
}

func preDateHeaders(dateHeaders []Column) map[string][]Column {
	return map[string][]Column{
		"Detailed Breakdown With Region": {
			{MapKey: "AccountName", Display: "Account"},
			{MapKey: "AccountEnvironment", Display: "Environment"},
			{MapKey: "Service", Display: "Service"},
			{MapKey: "Region", Display: "Region"},
		},
		"Detailed Breakdown": {
			{MapKey: "AccountName", Display: "Account"},
			{MapKey: "AccountEnvironment", Display: "Environment"},
			{MapKey: "Service", Display: "Service"},
		},
		"Service And Environment": {
			{MapKey: "AccountName", Display: "Account"},
			{MapKey: "AccountEnvironment", Display: "Environment"},
		},
		"Service": {
			{MapKey: "AccountName", Display: "Account"},
		},
		"Totals": {
			{MapKey: "Org", Display: " "},
		},
	}

}

func postDateHeaders(dateHeaders []Column, preHeaders map[string][]Column) (post map[string][]Column) {
	trendOptions := "{\"charttype\",\"column\";\"empty\",\"ignore\";\"nan\",\"convert\"}"
	post = map[string][]Column{}

	standard := []string{
		"Detailed Breakdown With Region",
		"Detailed Breakdown",
		"Service And Environment",
		"Service",
		"Totals",
	}

	for _, s := range standard {
		dateStart, _ := excelize.ColumnNumberToName(len(preHeaders[s]) + 1)
		dateEnd, _ := excelize.ColumnNumberToName(len(preHeaders[s]) + len(dateHeaders))
		post[s] = []Column{
			{
				Display: "Totals",
				Formula: fmt.Sprintf("=SUM(%s${r}:%s${r})", dateStart, dateEnd),
			},
			{
				Display: "Trend",
				Formula: fmt.Sprintf("=SPARKLINE(%s${r}:%s${r}, %s)", dateStart, dateEnd, trendOptions),
			},
		}
	}

	return
}
