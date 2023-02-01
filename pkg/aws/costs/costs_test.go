package costs

import (
	"opg-infra-costs/internal/project"
	"opg-infra-costs/pkg/aws/accounts"
	"os"
	"testing"
	"time"
)

func TestCosts(t *testing.T) {

	if sess := os.Getenv("AWS_SESSION_TOKEN"); len(sess) > 0 {

		now := time.Now().UTC()
		start := now.AddDate(0, 0, -2)
		end := now.AddDate(0, 0, -1)
		accountList := []accounts.Account{
			{Name: "test", Id: "050256574573", Environment: "Identity", Role: "breakglass"},
		}
		r, _, _ := Costs(
			accountList,
			start,
			end,
		)

		f, _ := os.CreateTemp(project.ROOT_DIR+"/files/", "*-test.csv")
		ToCSV(r, accountList, f.Name())
	}
}
