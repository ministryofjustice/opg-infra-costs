package fx

import (
	"io/ioutil"
	"opg-infra-costs/pkg/dates"
	"os"

	"gopkg.in/yaml.v2"
)

func Load(file string) (mapped map[string]float64) {
	fx := Rates{}
	mapped = make(map[string]float64)
	if f, err := os.Open(file); err == nil {
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		if loadErr := yaml.Unmarshal(content, &fx); loadErr == nil {
			for _, r := range fx.Rates {
				for _, m := range dates.Months(r.Start, r.End, dates.YM) {
					mapped[m] = r.Rate
				}
			}
		}
	}

	return

}
