package csv

import (
	"encoding/json"
	"opg-infra-costs/pkg/debugger"
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
	defer debugger.Log("Converted CSV costs to map", debugger.DETAILED)()
	mapped := []map[string]string{}
	for _, item := range items {
		mapped = append(mapped, item.Map())
	}

	return mapped
}
