package speedcurve

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"
)

const endpoint = "/deploys"

type (
	testDetails struct {
		Test     string `json:"test"`
		Browser  string `json:"browser"`
		Region   string `json:"region"`
		Template int    `json:"template"`
	}

	deployDetails struct {
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
)

var c *Client

func init() {
	conf := &Config{}
	c = NewClient(conf)
}

func Add(site, note, details string) (deployResponse, error) { // {{{
	var r deployResponse

	payload := url.Values{}
	payload.Add("site_id", site)
	payload.Add("note", note)
	payload.Add("details", details)

	req, _ := c.NewRequest("POST", endpoint, bytes.NewBufferString(payload.Encode()))
	//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return r, errors.New("request responded with errors")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
		return r, err
	}

	return r, nil
} // }}}

func Get(resource string) (deployDetails, error) { // {{{
	var d deployDetails

	parts := []string{endpoint, resource}
	uri := strings.Join(parts, "/")

	req, _ := c.NewRequest("GET", uri, nil)
	resp, err := c.Do(req)
	if err != nil {
		return d, errors.New("request responded with errors")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
		return d, err
	}

	return d, nil
} // }}}

func Getlatest() (deployDetails, error) { // {{{
	return Get("latest")
} // }}}
