package report2

import (
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/debugger"
	"time"
)

type MonthRange struct {
	MonthsAgoStart int `yaml:"start"`
	MonthsAgoEnd   int `yaml:"end"`
}

func (mr *MonthRange) Months() []string {
	defer debugger.Log("MonthRange.Months()", debugger.VVERBOSE)()
	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	start = start.AddDate(0, mr.MonthsAgoStart, 0)
	end := now.AddDate(0, mr.MonthsAgoEnd, 0)
	return dates.Months(start, end, dates.YM)
}

// IsNil checks that the range information is set to a
// valid number (non-zero)
//   - Used in loading from yaml to check if transpose data
//     is real of defaults
func (mr *MonthRange) Nil() (isNil bool) {
	defer debugger.Log("MonthRange.Nil()", debugger.VVERBOSE)()
	isNil = true
	if mr.MonthsAgoEnd != 0 && mr.MonthsAgoStart != 0 {
		isNil = false
	}
	return
}
