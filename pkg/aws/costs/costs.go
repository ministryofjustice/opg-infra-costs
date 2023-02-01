package costs

import (
	"opg-infra-costs/pkg/aws/accounts"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/gammazero/workerpool"
)

const PoolSize = 20

// Costs handles the async calls to aws api to
// fetch cost results for each account in the slice
// passed
func Costs(
	accounts []accounts.Account,
	start time.Time,
	end time.Time,
) (resultsByAccountId map[string]*costexplorer.GetCostAndUsageOutput, dur time.Duration, err error) {
	marker := time.Now().UTC()
	mu := &sync.Mutex{}
	errors := []error{}
	resultsByAccountId = map[string]*costexplorer.GetCostAndUsageOutput{}
	poolSize := PoolSize

	if len(accounts) < PoolSize {
		poolSize = len(accounts)
	}

	workerPool := workerpool.New(poolSize)

	for _, a := range accounts {
		account := a
		// push the call to get costs to the worker pool
		workerPool.Submit(func() {
			res, err := CostsForAccount(account, start, end)
			mu.Lock()
			if err != nil {
				errors = append(errors, err)
			} else {
				resultsByAccountId[account.Id] = res
			}
			mu.Unlock()
		})

	}

	workerPool.StopWait()
	if len(errors) > 0 {
		err = errors[0]
	}
	dur = time.Since(marker)
	return
}
