package fx

import "time"

type Rates struct {
	Rates []Rate
}
type Rate struct {
	Start time.Time
	End   time.Time
	Rate  float64
}
