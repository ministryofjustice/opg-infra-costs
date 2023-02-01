package out

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func pad(l int, char rune) string {
	var buffer bytes.Buffer
	for i := 0; i < l; i++ {
		buffer.WriteRune(char)
	}
	return buffer.String()
}

// outputs to the CLI with fixed width line
// - if duration > 0 then it includes that in seconds to the right side
func CLI(s string, duration time.Duration) {
	var outputWidth = 100
	//var buffer bytes.Buffer
	var durationWidth = 12
	var gap = 4
	var lineWidth int = outputWidth - durationWidth - gap
	var strChars int = utf8.RuneCountInString(s)
	var numLines int = (strChars / lineWidth) + 1
	lines := []string{}

	for line := 0; line < numLines; line++ {
		//var b bytes.Buffer
		var start = (line * lineWidth)
		var end = start + lineWidth
		var padded = ""
		if end > strChars {
			end = strChars
		}
		chunk := s[start:end] + padded
		// if this is the first line and there is a duration, add it on
		if line == 0 && duration > 0 {
			// limit number of decimals
			sec := fmt.Sprintf("%.7f", duration.Seconds())
			pre := pad(durationWidth-utf8.RuneCountInString(sec)-3, ' ')
			chunk = fmt.Sprintf("%-*s", lineWidth, chunk) + pad(gap, ' ') + fmt.Sprintf("(%s%v%s)", pre, sec, "s")
		}

		lines = append(lines, chunk)
	}

	s = strings.Join(lines, "\n")

	fmt.Println(s)

}
