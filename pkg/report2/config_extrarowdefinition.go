package report2

// ExtraRowDefinition stores data from the yaml config on if
// report needs to append extra rows.
// These rows are generally created from formulas to adjust data
// such as providing $ -> £ conversion on cost data
//   - Requires dataset to have been loaded
type ExtraRowDefinition struct {
	Name               string            `yaml:"name"`
	ColumnsByName      []string          `yaml:"columns"`
	ReplaceColumnValue map[string]string `yaml:"column_overwrite,omitempty"`
	Style              int               `yaml:"row_style,omitempty"`
}
