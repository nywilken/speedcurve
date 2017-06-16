package speedcurve

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type scTestDetails struct {
	Test     string `json:"test"`
	Browser  string `json:"browser"`
	Template int    `json:"template"`
}

type scDeployResponse struct {
	ID      int    `json:"deploy_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Note    string `json:"note"`
	Details string `json:"detail"`
	Info    struct {
		ScheduledTests []scTestDetails `json:"tests-added"`
	} `json:"info"`
	TestsRequested int             `json:"tests-requested"`
	CompletedTests []scTestDetails `json:"tests-completed"`
	RemainingTests []scTestDetails `json:"tests-remaining"`
}

// DeployAPI client
type DeployAPI struct {
	client   *Client
	endpoint string
}

// NewDeployAPI returns a API client capable of interacting with Speedcurve's /deploys endpoint.
func NewDeployAPI(c *Client) *DeployAPI {
	d := &DeployAPI{client: c, endpoint: "/deploys"}
	return d
}

//Add a deployment and trigger an additional round of testing for one of the sites.
func (d DeployAPI) Add(siteid, note, details string) (scDeployResponse, error) { // {{{
	data := url.Values{}
	data.Add("site_id", siteid)
	data.Add("note", note)
	data.Add("details", details)

	var dr scDeployResponse
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
func (d DeployAPI) Get(resource string) (scDeployResponse, error) { // {{{
	parts := []string{d.endpoint, resource}
	uri := strings.Join(parts, "/")

	var di scDeployResponse
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
func (d DeployAPI) Getlatest() (scDeployResponse, error) { // {{{
	return d.Get("latest")
} // }}}
