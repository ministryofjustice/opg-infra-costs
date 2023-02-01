package report2

import (
	"opg-infra-costs/pkg/dates"
	"time"
)

type MonthRange struct {
	MonthsAgoStart int
	MonthsAgoEnd   int
}

func (mr *MonthRange) Months() []string {
	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	start = start.AddDate(0, mr.MonthsAgoStart, 0)
	end := now.AddDate(0, mr.MonthsAgoEnd, 0)
	return dates.Months(start, end, dates.YM)
}
