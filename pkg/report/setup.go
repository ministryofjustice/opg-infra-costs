package report

import (
	"fmt"
	"opg-infra-costs/pkg/data/convert"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/report/fx"
	"time"

	"github.com/xuri/excelize/v2"
)

// Reports provides a set of pre-configured sheets
// with columns and data sets converted
func Reports(
	start time.Time,
	end time.Time,
	rawDataset []map[string]string,
	fxFile string,
) (sheets []Sheet) {
	var name string
	fxdata := fx.Load(fxFile)
	// generate the date headers
	dateHeaders := []Column{}
	for _, d := range dates.Months(start, end, dates.YM) {
		dateHeaders = append(dateHeaders, Column{MapKey: d, Display: d})
	}
	pre := preDateHeaders(dateHeaders)
	post := postDateHeaders(dateHeaders, pre)
	// -- Detailed Breakdown
	name = "Detailed Breakdown"
	detailedBreakdown := NewSheet(name)
	detailedBreakdown.SetColumns(pre[name], ColumnsAreGroupBy)
	detailedBreakdown.SetColumns(dateHeaders, ColumnsAreDateCost)
	detailedBreakdown.SetColumns(post[name], ColumnsAreOther)
	detailedBreakdownMeta, _ := detailedBreakdown.SetDataset(
		convert.Convert(
			rawDataset,
			detailedBreakdown.GetGroupColumns(),
			detailedBreakdown.GetTransposeColumns(),
			detailedBreakdown.GetOtherColumns(),
		),
	)

	// -- totals
	name = "Totals"
	totals := NewSheet(name)
	totalData := rawDataset
	totals.SetColumns(pre[name], ColumnsAreGroupBy)
	totals.SetColumns(dateHeaders, ColumnsAreDateCost)
	totals.SetColumns(post[name], ColumnsAreOther)
	totalMeta, _ := totals.SetDataset(
		convert.Convert(
			totalData,
			totals.GetGroupColumns(),
			totals.GetTransposeColumns(),
			totals.GetOtherColumns(),
		),
	)
	// add custom row of data for showing non-vat
	exVat := map[string][]string{
		"Org": {"($) excluding Tax"},
	}
	inGBP := map[string][]string{
		"Org": {"(£) excluding Tax"},
	}
	k := detailedBreakdown.GetName()
	noTaxStartCol := detailedBreakdownMeta["Header"].Columns["Service"] + 1 //len(pre[k]) + 1
	gbpStartCol := totalMeta["Header"].Columns["Org"] + 1                   //len(pre[name]) + 1
	for _, d := range dateHeaders {
		noTaxcol, _ := excelize.ColumnNumberToName(noTaxStartCol)
		gbpCol, _ := excelize.ColumnNumberToName(gbpStartCol)
		f := fmt.Sprintf("=SUMIF('%s'!C:C,\"<>Tax\",  '%s'!%s:%s)", k, k, noTaxcol, noTaxcol)
		gbp := fmt.Sprintf("=(%s3*%f)", gbpCol, fxdata[d.Display])
		exVat[d.MapKey] = []string{f}
		inGBP[d.MapKey] = []string{gbp}
		noTaxStartCol++
		gbpStartCol++
	}
	for _, x := range post[name] {
		exVat[x.MapKey] = []string{x.Formula}
		inGBP[x.MapKey] = []string{x.Formula}
	}
	totals.AddStyle(&excelize.Style{NumFmt: 190}, 4, 0)
	totals.AddRow("excluding-vat", exVat)
	totals.AddCell(2, 1, "($) including Tax")
	totals.AddRow("excluding-vat-in-gbp", inGBP)

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
			service.GetTransposeColumns(),
			service.GetOtherColumns(),
		),
	)

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
			serviceAndEnvironment.GetTransposeColumns(),
			serviceAndEnvironment.GetOtherColumns(),
		),
	)

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
			detailedBreakdownWithRegion.GetTransposeColumns(),
			detailedBreakdownWithRegion.GetOtherColumns(),
		),
	)

	// -- Cost changes
	// 	Has some custom columns and additional styles as well
	//	hiding rows that dont match certain criteria
	name = "Cost Changes"
	post["Cost Changes"] = []Column{
		{
			MapKey:  "Increase$",
			Display: "Increase ($)",
			Formula: "=(E${r}-D${r})",
		},
		{
			MapKey:  "Increase%",
			Display: "Increase (%)",
			Formula: "=IFERROR( (E${r}/D${r})-1, 0 )",
		},
	}
	// only want the last 2 months
	lastTwoMonths := dateHeaders[len(dateHeaders)-2:]
	// adjust name to include the dates
	label := fmt.Sprintf("Changes (%s - %s)", lastTwoMonths[0].Display, lastTwoMonths[1].Display)
	costChanges := NewSheet(label)
	// make the last column %
	costChanges.AddStyle(&excelize.Style{NumFmt: 10}, 0, 7)
	costChanges.SetColumns(pre[name], ColumnsAreGroupBy)
	costChanges.SetColumns(lastTwoMonths, ColumnsAreDateCost)
	costChanges.SetColumns(post[name], ColumnsAreOther)
	costsMeta, _ := costChanges.SetDataset(
		convert.Convert(
			rawDataset,
			costChanges.GetGroupColumns(),
			costChanges.GetTransposeColumns(),
			costChanges.GetOtherColumns(),
		),
	)
	incValueCol := costsMeta["Header"].Columns["Increase$"]
	serviceCol := costsMeta["Header"].Columns["Service"]
	incPercentCol := costsMeta["Header"].Columns["Increase%"]
	hide := map[CellRef]interface{}{
		{Col: incValueCol}:         20,       // value change of more than
		{Col: incPercentCol}:       0.15,     // percentage (as a decimal) change of more than
		{Col: serviceCol, Row: 0}:  "Tax",    // exclude Tax
		{Col: serviceCol, Row: -1}: "Refund", // exclude refunds
	}
	costChanges.SetHideRowWhen(hide)

	// -- add everything in
	sheets = append(sheets, totals)
	sheets = append(sheets, service)
	sheets = append(sheets, serviceAndEnvironment)
	sheets = append(sheets, detailedBreakdown)
	sheets = append(sheets, detailedBreakdownWithRegion)
	sheets = append(sheets, costChanges)

	return
}

// preDateHeaders returns a map based on sheet name of all
// columns required before the date/cost information is included
// Typically account name, environment, service etc
func preDateHeaders(dateHeaders []Column) map[string][]Column {
	return map[string][]Column{
		"Cost Changes": {
			{MapKey: "AccountName", Display: "Account"},
			{MapKey: "AccountEnvironment", Display: "Environment"},
			{MapKey: "Service", Display: "Service"},
		},
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

// postDateHeaders provides a map keyed to the sheet name of
// all columns that come after the date/cost cols.
// Typically for totals & trend lines
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
				MapKey:  "Totals",
				Display: "Totals",
				Formula: fmt.Sprintf("=SUM(%s${r}:%s${r})", dateStart, dateEnd),
			},
			{
				MapKey:  "Trend",
				Display: "Trend",
				Formula: fmt.Sprintf("=SPARKLINE(%s${r}:%s${r}, %s)", dateStart, dateEnd, trendOptions),
			},
		}
	}

	return
}
