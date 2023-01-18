package formats

import "time"

var DATES map[string]string = map[string]string{
	"ymd": "2006-01-22",
	"ym":  "2006-01",
	"RFC": time.RFC3339,
}
