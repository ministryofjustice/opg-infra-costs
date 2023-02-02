package report2

import "opg-infra-costs/pkg/debugger"

// ColumnDefinition details where and how data for a cell is sourced.
// Also deals with expanding the transposed columns (Date -> Cost)
// into Headings for the Sheet
// Stores style information from the config
type ColumnDefinition struct {
	SourceColumn    string     `yaml:"name"`
	Display         string     `yaml:"display"`
	Style           int        `yaml:"col_style,omitempty"`
	Formula         string     `yaml:"formula,omitempty"`
	TransposeMonths MonthRange `yaml:"month_range,omitempty"`
}

// IsTransposed checks to see if the transposed data is nil or
// not.
// Used when loading from yaml
func (cd *ColumnDefinition) IsTransposed() bool {
	defer debugger.Log("ColumnDefinition.IsTransposed()", debugger.VVERBOSE)()
	return !cd.TransposeMonths.Nil()
}

// IsFormula checks to see if the struct contains
// a valid formula
func (cd *ColumnDefinition) IsFormula() bool {
	defer debugger.Log("ColumnDefinition.IsFormula()", debugger.VVERBOSE)()
	return (len(cd.Formula) > 0 && cd.Formula[0:1] == "=")
}

// ToColumns converts this defintion into a `[]Column` to
// act as sheet headings
func (cd *ColumnDefinition) ToHeadings() []Column {
	defer debugger.Log("ColumnDefinition.ToHeadings()", debugger.VVERBOSE)()
	return NewColumnsFromDefinition(*cd)
}

// -- functions that operate on a slice of ColumnDefinitions

// Transposed returns only the ColumnDefinitions that have valid
// transposed data
func Transposed(definitions map[string]ColumnDefinition) (defs []ColumnDefinition) {
	defer debugger.Log("Transposed()", debugger.VVERBOSE)()
	for _, d := range definitions {
		if d.IsTransposed() {
			defs = append(defs, d)
		}
	}
	return
}

// Formulas returns just the ColumnDefinitions that have a valid formula
func Formulas(definitions map[string]ColumnDefinition) (defs []ColumnDefinition) {
	defer debugger.Log("Formulas()", debugger.VVERBOSE)()
	for _, d := range definitions {
		if d.IsFormula() {
			defs = append(defs, d)
		}
	}
	return
}
