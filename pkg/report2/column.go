package report2

import "opg-infra-costs/pkg/debugger"

// Column is used to define where and how data in cells comes from
//   - Does not require dataset loading, just config
//   - SourceColumn maps to the csv row field name
//   - Display is the value used for this column header
//   - Forumla can be set so the entire column uses a templated
//     formula (eg to work out SUM of a row)
type Column struct {
	SourceColumn string
	Display      string
	Formula      string
	Definition   *ColumnDefinition
}

// Key returns the value to use as a key against the csv data row
func (col *Column) Key() (key string) {
	defer debugger.Log("Column.Key()", debugger.VVERBOSE)()
	if len(col.SourceColumn) > 0 {
		key = col.SourceColumn
	} else if len(col.Display) > 0 {
		key = col.Display
	} else {
		key = ""
	}
	return
}

// -- New / Create methods

// NewColumnsFromDefinition uses the definition to generate a `[]Column`.
// For most, this will be singular, but for ColumnDefinition's that have
// tranpositions to do it will create one for each
//   - Formula data is carried over to each Column
//   - Defintiion styling information is not used here, but definitions
//     are attached for later user
func NewColumnsFromDefinition(definition ColumnDefinition) (cols []Column) {
	defer debugger.Log("NewColumnsFromDefinition()", debugger.VVERBOSE)()
	// if this is a transposed column, we need to generate multiple columns
	if definition.IsTransposed() {
		// create a column for each month in the transposed range
		for _, m := range definition.TransposeMonths.Months() {
			c := Column{
				SourceColumn: m,
				Display:      m,
				Formula:      definition.Formula,
				Definition:   &definition,
			}
			cols = append(cols, c)
		}
	} else {
		c := Column{
			SourceColumn: definition.SourceColumn,
			Display:      definition.Display,
			Formula:      definition.Formula,
			Definition:   &definition,
		}
		cols = append(cols, c)
	}
	return

}
