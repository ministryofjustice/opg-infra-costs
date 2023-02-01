package debugger

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var outputWidth int = 100
var durationWidth int = 12
var gap int = 4

func pad(l int, char rune) string {
	var buffer bytes.Buffer
	for i := 0; i < l; i++ {
		buffer.WriteRune(char)
	}
	return buffer.String()
}

func formatForCli(message string, duration time.Duration) string {
	var lineWidth int = outputWidth - durationWidth - gap
	var strChars int = utf8.RuneCountInString(message)
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
		chunk := message[start:end] + padded
		// if this is the first line and there is a duration, add it on
		if line == 0 && duration > 0 {
			// limit number of decimals
			sec := fmt.Sprintf("%.7f", duration.Seconds())
			pre := pad(durationWidth-utf8.RuneCountInString(sec)-3, ' ')
			chunk = fmt.Sprintf("%-*s", lineWidth, chunk) + pad(gap, ' ') + fmt.Sprintf("(%s%v%s)", pre, sec, "s")
		}

		lines = append(lines, chunk)
	}

	return strings.Join(lines, "\n")

}

func Log(message string, level int) func() {
	t := time.Now().UTC()
	return func() {
		p, _ := strconv.Atoi(fmt.Sprintf("%d", level))
		p = (p / 10) - 1
		pre := pad(p*2, ' ')
		message = fmt.Sprintf("%s%s", pre, message)
		str := formatForCli(message, time.Since(t))

		show := (level <= LEVEL)
		//fmt.Printf("(%v) (%v) [%v]\n", LEVEL, level, show)
		if show {
			fmt.Println(str)
		}
	}
}
