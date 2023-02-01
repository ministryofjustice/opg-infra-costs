package costs

import (
	"opg-infra-costs/pkg/aws/accounts"
	"opg-infra-costs/pkg/data/csv"
	"opg-infra-costs/pkg/dates"
	"opg-infra-costs/pkg/debug"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/gocarina/gocsv"
)

var org string = "opg"

func ToCSV(
	data map[string]*costexplorer.GetCostAndUsageOutput,
	accountList []accounts.Account,
	file string,
) (dur time.Duration, err error) {
	defer debug.Log("Wrote costs to CSV", 2)()
	marker := time.Now().UTC()
	rows := []csv.Row{}

	for accountId, usage := range data {
		account, _ := accounts.GetById(accountId, accountList)

		for _, resultByTime := range usage.ResultsByTime {
			day := *resultByTime.TimePeriod.Start
			asDate, _ := time.Parse(dates.YMD, day)
			month := asDate.Format(dates.YM)

			for _, costGroup := range resultByTime.Groups {

				for _, costMetric := range costGroup.Metrics {

					row := csv.Row{
						Org:                org,
						AccountId:          accountId,
						AccountName:        account.Name,
						AccountEnvironment: account.Environment,
						Service:            serviceNameCorrection(*costGroup.Keys[0]),
						Region:             *costGroup.Keys[1],
						Cost:               *costMetric.Amount,
						Date:               month,
					}
					rows = append(rows, row)

				}
			}
		}
	}

	f, _ := os.Create(file)
	defer f.Close()

	err = gocsv.MarshalFile(&rows, f)
	dur = time.Since(marker)
	return
}

func serviceNameCorrection(serviceName string) string {
	switch name := serviceName; name {
	case "Amazon EC2 Container Service":
		return "Amazon Elastic Container Service"
	case "Amazon Elasticsearch Service":
		return "Amazon OpenSearch Service"
	}
	return serviceName
}
