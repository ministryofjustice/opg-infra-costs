package csv

import (
	"encoding/json"
	"opg-infra-costs/pkg/debug"
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

func ToMap(items []Row) []map[string]string {
	defer debug.Log("Converted CSV costs to map", 2)()
	mapped := []map[string]string{}
	for _, item := range items {
		mapped = append(mapped, item.Map())
	}

	return mapped
}
