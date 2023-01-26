package report

import "strings"

func ParseFormula(
	formula string,
	replacements map[string]interface{},

) (parsed string, err error) {
	parsed = formula
	for k, r := range replacements {
		parsed = strings.ReplaceAll(parsed, "${"+k+"}", r.(string))
	}
	return
}
