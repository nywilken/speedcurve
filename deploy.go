package speedcurve

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

const deployEndpoint = "/deploys"

type (
	testDetails struct {
		Test     string `json:"test"`
		Browser  string `json:"browser"`
		Template int    `json:"template"`
	}

	deployInfo struct {
		Id             int           `json:"deploy_id"`
		Status         string        `json:"status"`
		Message        string        `json:"message"`
		Note           string        `json:"note"`
		Details        string        `json:"detail"`
		CompletedTests []testDetails `json:"tests-completed"`
		RemainingTests []testDetails `json:"tests-remaining"`
	}

	deployResponse struct {
		Id      int    `json:"deploy_id"`
		SiteId  int    `json:"site_id"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Info    struct {
			ScheduledTests []testDetails `json:"tests-added"`
		} `json:"info"`
		TestsRequested int `json:"tests-requested"`
	}

	Deploy struct {
		client *Client
	}
)

func NewDeploy(client *Client) *Deploy {
	return &Deploy{client}
}

func (d Deploy) Add(site, note, details string) (deployResponse, error) { // {{{
	var dr deployResponse

	data := url.Values{}
	data.Add("site_id", site)
	data.Add("note", note)
	data.Add("details", details)

	payload := bytes.NewBufferString(data.Encode())
	req, _ := d.client.NewRequest("POST", deployEndpoint, payload)
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

func (d Deploy) Get(resource string) (deployInfo, error) { // {{{
	var di deployInfo

	parts := []string{deployEndpoint, resource}
	uri := strings.Join(parts, "/")

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

func (d Deploy) Getlatest() (deployInfo, error) { // {{{
	return d.Get("latest")
} // }}}
