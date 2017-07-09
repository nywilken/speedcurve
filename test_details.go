// Copyright © 2017 Wilken Rivera

package speedcurve

import (
	"encoding/json"
)

type customMetric struct {
	Name  string `json:"mark"`
	Value string `json:"value"`
}

// TestDetails represents the response details when issuing a GET to
// Speedcurve's tests API endpoint.
type TestDetails struct {
	TestID         string         `json:"test_id"`
	URL            string         `json:"url"`
	Timezone       string         `json:"timezone"`
	Day            string         `json:"day"`
	Timestamp      int64          `json:"timestamp"`
	Region         string         `json:"region"`
	Browser        string         `json:"browser"`
	Status         int            `json:"status"`
	Requests       int            `json:"requests"`
	FirstByte      int64          `json:"byte"`
	StartRender    int64          `json:"render"`
	VisualComplete int64          `json:"visually_complete"`
	DomComplete    int64          `json:"dom"`
	Loaded         int64          `json:"loaded"`
	CustomMetrics  []customMetric `json:"custom_metrics"`
}

func (t TestDetails) String() string {
	o, _ := json.Marshal(t)
	return string(o)
}
