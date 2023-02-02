package report2

import (
	"testing"
)

var dummyCfg = `
reports:
  CostChanges:
    name: Cost Changes
    columns:
      - AccountName
      - AccountEnvironment
      - Service
      - TransposeCostChanges
      - FormulaIncrease$
      - FormulaIncrease%
  DetailedBreakdown:
    name: Detailed Breakdown
    columns:
      - AccountName
      - AccountEnvironment
      - Service
      - TransposeCosts
      - FormulaTotals
      - FormulaTrend
  Totals:
    name: Totals
    visible: false
    columns:
      - Org
      - TransposeCosts
      - FormulaTotals
      - FormulaTrend
    extra_rows:
      - ExcludeTax
      - To£
    overwrite_cells:
      '${Col:Org}2': ($) inc. Tax
extra_row_definitions:
  ExcludeTax:
    name: ExcludeTax
    overwrite_columns:
      Org: ($) exc. Tax
    columns:
      - Org
      - FormulaYearlyCostsNoTax
      - FormulaTotals
      - FormulaTrend
  To£:
    name: To£
    column_overwite:
      Org: (£) excluding Tax
    row_style: 190
    columns:
      - Org
      - FormulaYearlyCostsGBP
      - FormulaTotals
      - FormulaTrend
column_definitions:
  Org:
    name: Org
    display: ' '
  AccountName:
    name: AccountName
    display: Account
  AccountEnvironment:
    name: AccountEnvironment
    display: Environment
  Service:
    name: Service
    display: Service
  Region:
    name: Region
    display: Region
  FormulaTotals:
    name: FormulaTotals
    display: Totals
    formula: '=SUM(${transposeStart}${row}:${transposeEnd}${row})'
    col_style: 177
  FormulaTrend:
    name: FormulaTrend
    display: Trend
    formula: >-
      =SPARKLINE(${transposeStart}${row}:${transposeEnd}${row},
      "{"charttype","column";"empty","ignore";"nan","convert"}")
  FormulaIncrease$:
    name: FormulaIncrease$
    display: Increase ($)
    formula: '=(${transposeEnd}${row}-${transposeEnd}${row})'
    col_style: 177
  FormulaIncrease%:
    name: FormulaIncrease%
    display: Increase (%)
    formula: '=IFERROR( (${transposeEnd}${row}/${transposeEnd}${row})-1, 0 )'
    col_style: 10
  TransposeCosts:
    name: TransposeCosts
    col_style: 177
    month_range:
      start: -12
      end: -1
  TransposeCostChanges:
    name: TransposeCostChanges
    col_style: 177
    month_range:
      start: -2
      end: -1
  FormulaYearlyCostsNoTax:
    name: FormulaYearlyCostsNoTax
    display: no tax
    formula: >-
      =SUMIF('${DetailedBreakdown}{name}'!${DetailedBreakdown}{Col:Service},"<>Tax",
      '${DetailedBreakdown}{name}'!${col}:${col})
    col_style: 177
    month_range:
      start: -12
      end: -1
  FormulaYearlyCostsGBP:
    name: FormulaYearlyCostsGBP
    display: no tax
    formula: '=(${col}${row-1}*0.7)'
    col_style: 190
`

func TestUnmarshalContent(t *testing.T) {

	_, err := unmarshalConfig([]byte(dummyCfg))
	if err != nil {
		t.Errorf("unexpected error [%v]", err)
	}

}
