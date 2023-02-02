package report2

import (
	"fmt"
)

type ReportConfiguration struct {
	Name            string            `yaml:"name"`
	Visible         bool              `yaml:"visible,omitempty"`
	ColumnsByName   []string          `yaml:"columns"`
	ExtraRowsByName []string          `yaml:"extra_rows,omitempty"`
	CellOverwrites  map[string]string `yaml:"overwrite_cells,omitempty"`
}

// ColumnNamesToDefinitions finds all the definitions from this reports `ColumnsByName` value
//   - used to find only this reports definitions and not the full set
func (r *ReportConfiguration) ColumnNamesToDefinitions(definitions map[string]ColumnDefinition) (found map[string]ColumnDefinition, err error) {
	found = make(map[string]ColumnDefinition)
	for _, name := range r.ColumnsByName {
		if def, ok := definitions[name]; ok {
			found[name] = def
		} else {
			err = fmt.Errorf(fmt.Sprintf(ConfigurationColumnDefinitionNotFound, name))
		}
	}

	return
}

// Headings uses `ColumnsByName` and the definition data to create all
// the headings
//   - this handles transposing date ranges into columns
//   - does not do any value subsitutions
func (r *ReportConfiguration) Headings(definitions map[string]ColumnDefinition) (headings []Column, err error) {
	if len(r.ColumnsByName) > 0 {
		for _, name := range r.ColumnsByName {
			if def, ok := definitions[name]; ok {
				headings = append(headings, def.ToHeadings()...)
			} else {
				err = fmt.Errorf(fmt.Sprintf(ConfigurationColumnDefinitionNotFound, name))
			}

		}
	} else {
		err = fmt.Errorf(ConfigurationReportHasNoColumns)
	}
	return
}
