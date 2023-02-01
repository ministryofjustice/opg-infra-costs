package accounts

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

func Load(file string) ([]Account, time.Duration, error) {
	marker := time.Now().UTC()
	buf, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, 0, err
	}
	list := &Accounts{}
	err = yaml.Unmarshal(buf, list)

	return list.Accounts, time.Since(marker), err
}
