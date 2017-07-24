// Copyright © 2017 Wilken Rivera

/*
Package speedcurve provides a Go client for working with Speedcurve's Web Page Test API.
For more information on Speedcurve's V1 API, see https://api.speedcurve.com.

Example Usage:

	// Create client with Speedcurve Token
	token := "SpeedcurveAPITokenString"
	sc := speedcurve.NewClient(token, "")

	// Get the latest deploy information
	d, _ := sc.GetLatestDeploy()
	fmt.Println(d)
*/
package speedcurve

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client is an HTTP client wrapper for communicating with Speedcurve's API
type Client struct {
	// APIToken is the account secret for authenticating to Speedcurve's API.
	APIToken string
	// APIHost specifies the API URI to connect to for communicating with Speedcurve's API.
	// Defaults to ScBaseURI
	APIHost string

	client *http.Client
}

// NewClient returns a speedcurve.Client with credentials for the Speedcurve API.
func NewClient(token, host string) *Client {
	if host == "" {
		host = "https://api.speedcurve.com/v1"
	}

	return &Client{APIToken: token, APIHost: host, client: http.DefaultClient}
}

// NewRequest returns an http.Request with information for the Speedcurve API.
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	parts := []string{c.APIHost, url}
	uri := strings.Join(parts, "")

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIToken, "x")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// Do performs an http.Request to Speedcurve API endpoint.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//AddDeploy triggers a Speedcurve deploy run for the specified site.
func (c *Client) AddDeploy(site, note, details string) (DeployResponse, error) {
	data := url.Values{}
	data.Add("site_id", site)
	data.Add("note", note)
	data.Add("detail", details)

	var d DeployResponse
	req, _ := c.NewRequest("POST", "/deploys", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil
}

//GetDeploy retrieves all the details available for a specific deploy.
func (c *Client) GetDeploy(resource string) (DeployDetails, error) {
	parts := []string{"/deploys", resource}
	uri := strings.Join(parts, "/")

	var d DeployDetails
	req, _ := c.NewRequest("GET", uri, nil)
	resp, err := c.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil
}

//GetLatestDeploy retrieves all the details available for the last deploy.
func (c *Client) GetLatestDeploy() (DeployDetails, error) {
	return c.GetDeploy("latest")
}

//GetTest retrieves all the details available for a specific test.
func (c *Client) GetTest(resource string) (TestDetails, error) {
	var tr TestDetails

	parts := []string{"/tests", resource}
	uri := strings.Join(parts, "/")

	req, _ := c.NewRequest("GET", uri, nil)
	resp, err := c.Do(req)
	if err != nil {
		return tr, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return tr, err
	}

	return tr, nil
}
