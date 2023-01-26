package dates

import "time"

func Months(
	start time.Time,
	end time.Time,
	format string,
) []string {
	asString := []string{}
	for d := start; !d.After(end); d = d.AddDate(0, 1, 0) {
		str := d.Format(format)
		asString = append(asString, str)
	}

	return asString
}
