package csv

import (
	"os"

	"github.com/gocarina/gocsv"
)

func Load(file string) []Row {
	var costs []Row = []Row{}

	if f, err := os.Open(file); err == nil {
		defer f.Close()

		if loadErr := gocsv.Unmarshal(f, &costs); loadErr == nil {
			return costs
		}
	}

	return costs

}
