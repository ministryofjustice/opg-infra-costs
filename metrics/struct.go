package metrics

import (
	"fmt"
	"opg-infra-costs/dates"
	costs "opg-infra-costs/unblendedcosts"
	"strconv"
	"strings"
	"time"

	"encoding/json"
)

type MetricsData struct {
	Dimensions       string `json:"Dimensions"`
	Project          string `json:"Project"`
	Environment      string `json:"Environment"`
	Service          string `json:"Service"`
	MeasureName      string `json:"MeasureName"`
	MeasureValue     string `json:"MeasureValue"`
	MeasureValueType string `json:"MeasureValueType"`
	Time             string `json:"Time"`
}

// FromCostRow converts a CostRow struct to a MetricsData Structure for sending
// to the API
func (md *MetricsData) FromCostRow(cr costs.CostRow) {
	md.Dimensions = "dimensions"
	md.Project = cr.Account.Name
	md.Service = cr.Service
	md.Environment = cr.Account.Environment
	md.MeasureName = "cost"
	md.MeasureValue = fmt.Sprintf("%f", cr.Cost)
	md.MeasureValueType = "DOUBLE"
	mytime, _ := time.Parse(dates.AWSDateFormat(), cr.Date)
	t := mytime.UnixNano() / int64(time.Millisecond)
	md.Time = strconv.FormatInt(t, 10)
}

type MetricsRecord struct {
	Data      string `json:"data"`
	Partition string `json:"partition-key"`
}

type MetricsPutData struct {
	Records []MetricsRecord `json:"records"`
}

// FromCosts converts CostData struct to a json byte array ready for sending
// in http call
func FromCosts(costs costs.CostData, limit int) ([]byte, error) {
	mpd := MetricsPutData{}

	for _, c := range costs.Entries {
		record := MetricsRecord{}
		record.Partition = "some key"
		data := MetricsData{}
		data.FromCostRow(c)
		j, _ := json.Marshal(data)
		d := string(j)
		record.Data = strings.ReplaceAll(d, `"`, `'`)
		mpd.Records = append(mpd.Records, record)

	}

	if limit != -1 {
		mpd.Records = mpd.Records[0:limit]
	}
	return json.Marshal(mpd)
}
