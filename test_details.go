// Copyright © 2017 Wilken Rivera

package speedcurve

type customMetric struct {
	Name  string `json:"mark"`
	Value string `json:"value"`
}

//TestDetails ...
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
