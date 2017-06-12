package speedcurve

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const APIHost = "https://api.speedcurve.com/v1"

type Config struct {
	APIHost  string
	APIToken string
}

type Client struct {
	client *http.Client
	host   string
	token  string
}

// NewClient returns a speedcurve.Client with credentials for the Speedcurve API.
func NewClient(conf *Config) *Client {
	t := conf.APIToken
	h := conf.APIHost

	if conf.APIToken == "" {
		v, ok := os.LookupEnv("SPD_API_TOKEN")
		if ok != true {
			log.Fatalln("Unable to find Speedcurve API token.")
		}
		t = v
	}

	if conf.APIHost == "" {
		h = APIHost
	}

	return &Client{client: http.DefaultClient, host: h, token: t}
}

// NewRequest returns an http.Request with information for the Speedcurve API.
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	parts := []string{c.host, url}
	uri := strings.Join(parts, "")

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.token, "x")
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
