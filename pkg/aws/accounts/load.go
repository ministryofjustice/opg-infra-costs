package accounts

import (
	"io/ioutil"
	"opg-infra-costs/pkg/debug"

	"gopkg.in/yaml.v2"
)

func Load(file string) ([]Account, error) {
	defer debug.Log("Accounts config loaded.", 2)()

	buf, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}
	list := &Accounts{}
	err = yaml.Unmarshal(buf, list)

	return list.Accounts, err
}
