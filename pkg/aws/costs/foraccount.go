package costs

import (
	"fmt"
	"opg-infra-costs/pkg/aws/accounts"
	"opg-infra-costs/pkg/aws/session"
	"opg-infra-costs/pkg/dates"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/k0kubun/pp"
)

var granularity string = costexplorer.GranularityMonthly
var costExplorerInput = &costexplorer.GetCostAndUsageInput{
	Granularity: aws.String(granularity),
	Metrics: []*string{
		aws.String("UNBLENDED_COST"),
	},
	GroupBy: []*costexplorer.GroupDefinition{
		{
			Type: aws.String("DIMENSION"),
			Key:  aws.String("SERVICE"),
		},
		{
			Type: aws.String("DIMENSION"),
			Key:  aws.String("REGION"),
		},
	},
}

func CostsForAccount(
	account accounts.Account,
	start time.Time,
	end time.Time,
) (*costexplorer.GetCostAndUsageOutput, error) {
	dateFormat := dates.YMD

	arn := fmt.Sprintf("arn:aws:iam::%s:role/%s", account.Id, account.Role)
	region := "eu-west-1"
	// auth
	costExplorerSession, err := session.Session(arn, region)
	if err != nil {
		pp.Println(err)
		return &costexplorer.GetCostAndUsageOutput{}, err
	}
	// call
	sdkInput := costExplorerInput
	sdkInput.TimePeriod = &costexplorer.DateInterval{
		Start: aws.String(start.Format(dateFormat)),
		End:   aws.String(end.Format(dateFormat)),
	}
	request, response := costExplorerSession.GetCostAndUsageRequest(sdkInput)
	err = request.Send()
	if err != nil {
		return &costexplorer.GetCostAndUsageOutput{}, err
	}

	return response, nil

}
