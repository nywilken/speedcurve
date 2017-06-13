package speedcurve

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"
)

const deployEndpoint = "/deploys"

type (
	testDetails struct {
		Test     string `json:"test"`
		Browser  string `json:"browser"`
		Region   string `json:"region"`
		Template int    `json:"template"`
	}

	deployInfo struct {
		Id             int           `json:"deploy_id"`
		Status         string        `json:"status"`
		Message        string        `json:"message"`
		Note           string        `json:"note"`
		Details        string        `json:"detail"`
		CompletedTests []testDetails `json:"tests-completed"`
		RemainingTests []testDetails `json:"test-remaining"`
	}

	deployResponse struct {
		Id      int           `json:"deploy_id"`
		SiteId  int           `json:"site_id"`
		Status  string        `json:"status"`
		Message string        `json:"message"`
		Info    []testDetails `json:"info"`
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

	payload := url.Values{}
	payload.Add("site_id", site)
	payload.Add("note", note)
	payload.Add("details", details)

	req, _ := d.client.NewRequest("POST", deployEndpoint, bytes.NewBufferString(payload.Encode()))
	//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := d.client.Do(req)
	if err != nil {
		return dr, errors.New("request responded with errors")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&dr)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
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
		return di, errors.New("request responded with errors")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&di)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
		return di, err
	}

	return di, nil
} // }}}

func (d Deploy) Getlatest() (deployInfo, error) { // {{{
	return d.Get("latest")
} // }}}
