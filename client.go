package speedcurve

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const scAPIURL = "https://api.speedcurve.com/v1"

// Client ...
type Client struct {
	client   *http.Client
	APIHost  string
	APIToken string
}

// NewClient returns a speedcurve.Client with credentials for the Speedcurve API.
func NewClient(host, token string) *Client {
	t := token
	h := host

	if t == "" {
		v, ok := os.LookupEnv("SPD_API_TOKEN")
		if ok != true {
			log.Fatalln("Unable to find Speedcurve API token.")
		}
		t = v
	}

	if h == "" {
		h = scAPIURL
	}

	return &Client{client: http.DefaultClient, APIHost: h, APIToken: t}
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
