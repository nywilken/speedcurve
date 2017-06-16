package speedcurve

import (
	"encoding/json"
	"strings"
)

const testEndpoint = "/tests"

type (
	mark struct {
		Name  string `json:"mark"`
		Value string `json:"value"`
	}

	testInfo struct {
		Id             string `json:"test_id"`
		Url            string `json:"url"`
		Browser        string `json:"browser"`
		Status         int    `json:"status"`
		Requests       int    `json:"requests"`
		Render         int    `json:"render"`
		VisualComplete int    `json:"visually-complete"`
		Loaded         int    `json:"loaded"`
		Marks          []mark `json:"custom_metrics"`
	}

	TestRun struct {
		client *Client
	}
)

func NewTestRun(client *Client) *TestRun {
	return &TestRun{client}
}

func (t TestRun) Get(resource string) (testInfo, error) { // {{{
	var ti testInfo

	parts := []string{testEndpoint, resource}
	uri := strings.Join(parts, "/")

	req, _ := t.client.NewRequest("GET", uri, nil)
	resp, err := t.client.Do(req)
	if err != nil {
		return ti, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ti)
	if err != nil {
		return ti, err
	}

	return ti, nil
} // }}}
