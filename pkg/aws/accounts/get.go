package accounts

import "fmt"

func GetById(accountId string, accounts []Account) (foundAccount Account, err error) {

	for _, acc := range accounts {
		if acc.Id == accountId {
			return acc, nil
		}
	}
	return Account{}, fmt.Errorf("could not find account with matching id: (%s)", accountId)
}
