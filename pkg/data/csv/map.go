package csv

import (
	"encoding/json"
	"time"
)

// Map uses json marshall to convert struct
// over to a map
func (r Row) Map() map[string]string {
	mapped := make(map[string]string)
	// convert to json
	asJson, _ := json.Marshal(r)
	// swap to map
	json.Unmarshal(asJson, &mapped)
	return mapped
}

func ToMap(items []Row) ([]map[string]string, time.Duration) {
	marker := time.Now().UTC()
	mapped := []map[string]string{}
	for _, item := range items {
		mapped = append(mapped, item.Map())
	}
	dur := time.Since(marker)
	return mapped, dur
}
