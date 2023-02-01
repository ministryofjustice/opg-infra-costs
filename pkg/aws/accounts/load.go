package accounts

import (
	"io/ioutil"
	"opg-infra-costs/pkg/debugger"

	"gopkg.in/yaml.v2"
)

func Load(file string) ([]Account, error) {
	defer debugger.Log("Accounts config loaded.", debugger.DETAILED)()

	buf, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}
	list := &Accounts{}
	err = yaml.Unmarshal(buf, list)

	return list.Accounts, err
}
