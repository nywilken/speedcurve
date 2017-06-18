package speedcurve

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// ScBaseURI specifies Speedcurve's default API URI.
const ScBaseURI = "https://api.speedcurve.com/v1"

// Client is an HTTP client wrapper for communicating with Speedcurve's API
type Client struct {
	client *http.Client

	// APIHost specifies the API URI to connect to for communicating with Speedcurve's API.
	// Defaults to ScBaseURI
	APIHost  string
	APIToken string
}

// NewClient returns a speedcurve.Client with credentials for the Speedcurve API.
func NewClient(host, token string) *Client {
	if token == "" {
		v, ok := os.LookupEnv("SPD_API_TOKEN")
		if ok != true {
			log.Fatalln("Unable to find Speedcurve API token.")
		}
		token = v
	}

	if host == "" {
		host = ScBaseURI
	}

	return &Client{client: http.DefaultClient, APIHost: host, APIToken: token}
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
