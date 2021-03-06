package main

import (
	"fmt"
	"opg-infra-costs/commands"
	"opg-infra-costs/commands/detail"
	"opg-infra-costs/commands/excel"
	"opg-infra-costs/commands/increases"
	"opg-infra-costs/commands/monthtodate"
	"opg-infra-costs/commands/sendtometrics"
	"opg-infra-costs/commands/yeartodate"
	"os"
)

func usage(commands []commands.Command) {
	fmt.Println("Available commands listed below:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf(" *%s*:\n", cmd.Name)
		cmd.Set.PrintDefaults()
		fmt.Println()
	}
	fmt.Println()
	os.Exit(1)

}
func main() {

	detailCmd, _ := detail.Command()
	excelCmd, _ := excel.Command()
	mtdCmd, _ := monthtodate.Command()
	ytdCmd, _ := yeartodate.Command()
	increasesCmd, _ := increases.Command()
	metricsCmd, _ := sendtometrics.Command()
	allCmds := []commands.Command{
		detailCmd,
		excelCmd,
		increasesCmd,
		mtdCmd,
		ytdCmd,
		metricsCmd}

	if len(os.Args) < 2 {
		usage(allCmds)
	}
	var err error

	switch os.Args[1] {
	case increasesCmd.Name:
		err = increases.Run(increasesCmd)
	case detailCmd.Name:
		err = detail.Run(detailCmd)
	case excelCmd.Name:
		err = excel.Run(excelCmd)
	case mtdCmd.Name:
		err = monthtodate.Run(mtdCmd)
	case ytdCmd.Name:
		err = yeartodate.Run(ytdCmd)
	case metricsCmd.Name:
		err = sendtometrics.Run(metricsCmd)
	default:
		usage(allCmds)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
