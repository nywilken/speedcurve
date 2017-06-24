package speedcurve

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

//AddDeploy triggers a Speedcurve deploy run for the specified site.
func (c *Client) AddDeploy(site, note, details string) (DeployResponse, error) {
	data := url.Values{}
	data.Add("site_id", site)
	data.Add("note", note)
	data.Add("details", details)

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
