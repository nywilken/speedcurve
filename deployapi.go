package speedcurve

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type testinfo struct {
	Test     string `json:"test"`
	Browser  string `json:"browser"`
	Template int    `json:"template"`
}

type DeployDetails struct {
	DeployID       int64      `json:"deploy_id"`
	SiteID         int64      `json:"site_id"`
	Status         string     `json:"status"`
	Note           string     `json:"note"`
	Detail         string     `json:"detail"`
	TestsCompleted []testinfo `json:"tests-completed"`
	TestsRemaining []testinfo `json:"tests-remaining"`
}

type DeployResponse struct {
	DeployID int64  `json:"deploy_id"`
	SiteID   int64  `json:"site_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	Info     struct {
		ScheduledTests []testinfo `json:"tests-added"`
	} `json:"info"`
	TestsRequested int `json:"tests-requested"`
}

// DeployAPI client
type DeployAPI struct {
	client   *Client
	endpoint string
}

// NewDeployAPI returns a API client capable of interacting with Speedcurve's /deploys endpoint.
func NewDeployAPI(c *Client) *DeployAPI { // {{{
	d := &DeployAPI{client: c, endpoint: "/deploys"}
	return d
} // }}}

//Add a deployment and trigger an additional round of testing for one of the sites.
func (d DeployAPI) Add(siteid, note, details string) (DeployResponse, error) { // {{{
	data := url.Values{}
	data.Add("site_id", siteid)
	data.Add("note", note)
	data.Add("details", details)

	var dr DeployResponse
	req, _ := d.client.NewRequest("POST", d.endpoint, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := d.client.Do(req)
	if err != nil {
		return dr, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&dr)
	if err != nil {
		return dr, err
	}

	return dr, nil
} // }}}

// Get the details for a particular deployment.
func (d DeployAPI) Get(resource string) (DeployDetails, error) { // {{{
	parts := []string{d.endpoint, resource}
	uri := strings.Join(parts, "/")

	var di DeployDetails
	req, _ := d.client.NewRequest("GET", uri, nil)
	resp, err := d.client.Do(req)
	if err != nil {
		return di, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&di)
	if err != nil {
		return di, err
	}

	return di, nil
} // }}}

//Getlatest returns details and status of testing for the latest deployment.
func (d DeployAPI) Getlatest() (DeployDetails, error) { // {{{
	return d.Get("latest")
} // }}}
