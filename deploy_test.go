package speedcurve

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var deploy *DeployAPI
var client *Client

func dmocksrv() *httptest.Server { // {{{
	f := func(w http.ResponseWriter, r *http.Request) {

		var code int
		var resp string

		switch r.URL.Path {
		case "/deploys":
			code = 200
			resp = `{
				"deploy_id": 123,
				"site_id": 11789,
				"status": "success",
				"message": "A deployment has been added",
				"info": { "tests-added": [ { "test": "blah", "browser": "Chrome", "region": {"value": "us-east-1"}, "template": 0 } ] },
				 "test-requested": 2
			}`
		case "/deploys/0":
			code = 404
			resp = `{"deploy_id": 0, "status": "no such deployment"}`
		case "/deploys/latest":
			code = 200
			resp = `{"deploy_id": 91088, "status": "completed"}`
		case "/deploys/91088":
			code = 200
			resp = `{
				"deploy_id": 91088,
				"status": "completed",
				"note": "short note",
				"detail": "long note",
				"tests-completed": [ { "test": "blah", "browser": "Chrome", "region": "us-east-1", "template": 0 } ],
				"tests-remaining": [ { "test": "blah", "browser": "Firefox", "region": "us-east-1", "template": 0 } ]
			}`
		default:
			code = 500
			resp = `{}`
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, resp)
	}

	return httptest.NewServer(http.HandlerFunc(f))
} // }}}

func TestDeployGet(t *testing.T) {
	server := dmocksrv()
	defer server.Close()

	client = NewClient(server.URL, "x")
	deploy = NewDeployAPI(client)

	var deployGetCases = []struct {
		id     string
		desc   string
		status string
	}{
		{"0", "a non-existent id", "no such deployment"},
		{"latest", "the latest deploy", "completed"},
		{"91088", "deploy id 91088", "completed"},
	}

	t.Log("Given the need to get a deploy from Speedcurve")
	for _, tc := range deployGetCases {
		t.Logf("\tWhen requesting deploy details for %s", tc.desc)
		resp, _ := deploy.Get(tc.id)
		if resp.Status != tc.status {
			t.Errorf("\t\tShould have gotten status %s but got %s.", tc.status, resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}

func TestDeployAdd(t *testing.T) {
	server := dmocksrv()
	defer server.Close()

	client = NewClient(server.URL, "x")
	deploy = NewDeployAPI(client)
	var deployAddCases = []struct {
		site    string
		note    string
		details string
		status  string
	}{
		{"11789", "note", "detail", "success"},
	}

	t.Log("Given the need to add a new deploy.")
	for _, tc := range deployAddCases {
		t.Logf("\tWhen adding a new deploy for %s", tc.site)
		resp, err := deploy.Add(tc.site, tc.note, tc.details)
		if err != nil {
			t.Errorf("\t\tFailed with an error: %s", err)
			return
		}

		if resp.Status != tc.status {
			t.Errorf("\t\tShould have gotten status %s but got %s.", tc.status, resp.Status)
			return
		}

		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}
