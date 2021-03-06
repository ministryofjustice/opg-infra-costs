package summary

import (
	"fmt"
	"opg-infra-costs/dates"
	costs "opg-infra-costs/unblendedcosts"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Summary(
	costData costs.CostData,
	startDate time.Time,
	endDate time.Time,
	account string,
	env string) {
	// work out total
	total := costData.Total()
	days := (endDate.Sub(startDate).Hours() / 24)
	average := (total / days)
	p := message.NewPrinter(language.English)
	message := p.Sprintf("Total cost: ($) %f\n  Daily average: ($) %f", total, average)

	// filters that have been used
	basedOn := fmt.Sprintf("  Between [%v] and [%v]",
		startDate.Format(dates.AWSDateFormat()),
		endDate.Format(dates.AWSDateFormat()))

	if len(account) > 0 {
		basedOn = basedOn + fmt.Sprintf(" for account [%s]", account)
	}
	if len(env) > 0 {
		basedOn = basedOn + fmt.Sprintf(" for environment [%s]", env)
	}

	fmt.Println(message)
	fmt.Println(basedOn)
}
