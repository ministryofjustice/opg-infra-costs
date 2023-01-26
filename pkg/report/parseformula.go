package report

import "strings"

func ParseFormula(
	formula string,
	replacements map[string]interface{},

) (parsed string, err error) {
	parsed = strings.ReplaceAll(formula, "${r}", replacements["r"].(string))
	return
}
