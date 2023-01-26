package convert

import "encoding/json"

func Key(data map[string]string, fields []string) string {

	keys := make(map[string]string)

	for _, field := range fields {
		if val, ok := data[field]; ok {
			keys[field] = val
		}
	}

	key, _ := json.Marshal(keys)

	return string(key)

}
