package speedcurve

import (
	"io"
	"net/http"
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
