package speedcurve

import (
	"encoding/json"
	"strings"
)

type customMetric struct {
	Name  string `json:"mark"`
	Value string `json:"value"`
}

//TestDetails is a type containing the details available for a specific test.
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

// TestAPI client
type TestAPI struct {
	client   *Client
	endpoint string
}

// NewTestAPI returns a API client capable of interacting with Speedcurve's /tests endpoint.
func NewTestAPI(c *Client) *TestAPI { // {{{
	t := &TestAPI{client: c, endpoint: "/tests"}
	return t
} // }}}

// Get retrieves all the details available for a specific test.
func (t TestAPI) Get(resource string) (TestDetails, error) { // {{{
	var tr TestDetails

	parts := []string{t.endpoint, resource}
	uri := strings.Join(parts, "/")

	req, _ := t.client.NewRequest("GET", uri, nil)
	resp, err := t.client.Do(req)
	if err != nil {
		return tr, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return tr, err
	}

	return tr, nil
} // }}}
