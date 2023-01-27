package report

import "strings"

func ParseFormula(
	subject string,
	replacements map[string]interface{},

) (parsed string, err error) {
	parsed = subject
	for k, r := range replacements {
		parsed = strings.ReplaceAll(parsed, "${"+k+"}", r.(string))
	}
	return
}
