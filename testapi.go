package speedcurve

import (
	"encoding/json"
	"strings"
)

type scCustomMark struct {
	Name  string `json:"mark"`
	Value string `json:"value"`
}

type scTestResponse struct {
	ID             string         `json:"test_id"`
	URL            string         `json:"url"`
	Browser        string         `json:"browser"`
	Status         int            `json:"status"`
	Requests       int            `json:"requests"`
	Render         int            `json:"render"`
	VisualComplete int            `json:"visually-complete"`
	Loaded         int            `json:"loaded"`
	Marks          []scCustomMark `json:"custom_metrics"`
}

// TestAPI ...
type TestAPI struct {
	client   *Client
	endpoint string
}

// NewTestAPI returns a API client capable of interacting with Speedcurve's /tests endpoint.
func NewTestAPI(conf *Config) *TestAPI {
	t := &TestAPI{}
	t.client = NewClient(conf)
	t.endpoint = "/tests"
	return t
}

// Get retrieves all the details available for a specific test.
func (t TestAPI) Get(resource string) (scTestResponse, error) { // {{{
	var tr scTestResponse

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
