package monthtodate

import (
	"flag"
	"fmt"
	"opg-infra-costs/accounts"
	"opg-infra-costs/commands"
	"opg-infra-costs/dates"
	"opg-infra-costs/summary"
	"opg-infra-costs/tabular"
	costs "opg-infra-costs/unblendedcosts"
	"os"
	"time"
)

// DetailCommand set up the input and name
func Command() (commands.Command, error) {
	cmd := commands.Command{Name: "mtd"}
	set := flag.NewFlagSet(cmd.Name, flag.ExitOnError)

	commands.ArgumentStandardFilters(set, true)
	set.Bool("breakdown", false, "Include a tablurar breakdown grouped by account and envrionment")

	cmd.Set = set
	return cmd, nil
}

// parseCommand uses the flag set in Command to handle and parse
// the arguments
func parseCommand(
	cmd commands.Command) (
	accountName string,
	environment string,
	service string,
	breakdown bool,
	err error) {

	cmdSet := cmd.Set

	accountName = cmdSet.Lookup("account").Value.String()
	environment = cmdSet.Lookup("env").Value.String()
	service = cmdSet.Lookup("service").Value.String()

	bdown := cmdSet.Lookup("breakdown").Value.String()
	if bdown == "true" {
		breakdown = true
	} else {
		breakdown = false
	}

	return

}

// Run the command
func Run(cmd commands.Command) error {
	// parse the args, skipping the 'detail' namespace
	cmd.Set.Parse(os.Args[2:])
	// get all the command arguments
	account, env, service, breakdown, err := parseCommand(cmd)
	if err != nil {
		return err
	}
	endDate := time.Now()
	month := int(endDate.Month())

	startStr := fmt.Sprintf("%d-%02d-%s", endDate.Year(), month, "01")
	startDate, _ := time.Parse(dates.AWSDateFormat(), startStr)

	allAccounts := accounts.Filtered(account, env)
	period := "MONTHLY"
	costData, e := costs.AsyncCosts(
		&allAccounts,
		startDate,
		endDate,
		period,
		service)

	if len(e) > 0 {
		return e[0]
	}
	// how do we output this - table is default
	if breakdown {
		headers := []string{"AccountName", "Environment", "Cost"}
		row := []string{"Account.Name", "Account.Environment", "Cost"}
		group := []string{"Account.Name", "Account.Environment"}

		grouped := costData.GroupByKeys(group)
		tabular.Table(grouped, headers, row)
		summary.Summary(grouped, startDate, endDate, account, env)
	} else {
		summary.Summary(costData, startDate, endDate, account, env)
	}

	return nil

}
