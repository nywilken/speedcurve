// Copyright © 2017 Wilken Rivera

package speedcurve_test

import (
	"fmt"
	"github.com/nywilken/speedcurve"
	"net/http"
	"net/http/httptest"
	"testing"
)

var client *speedcurve.Client

func mocksrv() *httptest.Server { // {{{
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
		case "/tests/0":
			code = 404
			resp = `{"error": "Test not found"}`
		case "/tests/117":
			code = 200
			resp = `{"deploy_id": 117, "status":0, "requests":56}`
		case "/tests/989":
			code = 200
			resp = `{
				"deploy_id": 989,
				"status":0,
				"requests":56,
				"custom_metrics": [
						{
							"mark": "hls_request_m3u8",
							"value": "250"
						},
						{
							"mark": "hls_request_segement_1",
							"value": "1240"
						}
					]
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

func TestGetDeploy(t *testing.T) {
	server := mocksrv()
	defer server.Close()

	client = speedcurve.NewClient("x", server.URL)

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
		resp, _ := client.GetDeploy(tc.id)
		if resp.Status != tc.status {
			t.Errorf("\t\tShould have gotten status %s but got %s.", tc.status, resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}

func TestAddDeploy(t *testing.T) {
	server := mocksrv()
	defer server.Close()

	client = speedcurve.NewClient("x", server.URL)

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
		resp, err := client.AddDeploy(tc.site, tc.note, tc.details)
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

func TestGetTest(t *testing.T) {
	server := mocksrv()
	defer server.Close()

	client = speedcurve.NewClient("x", server.URL)

	var tt = []struct {
		id               string
		desc             string
		expectedRequests int
		expectedMetrics  int
	}{
		{"0", "a non-existent test run", 0, 0},
		{"117", "a completed test run", 56, 0},
		{"989", "a completed test run with custom marks", 56, 2},
	}

	t.Log("Given the need to get a test run from Speedcurve")
	for _, tc := range tt {
		t.Logf("\tWhen requesting test run details for %s", tc.desc)
		resp, _ := client.GetTest(tc.id)
		if resp.Requests != tc.expectedRequests {
			t.Errorf("\t\tShould have a request count of %d, but got %d.", tc.expectedRequests, resp.Requests)
			return
		}
		if len(resp.CustomMetrics) != tc.expectedMetrics {
			t.Errorf("\t\tShould have %d custom marks, but got %d.", tc.expectedMetrics, len(resp.CustomMetrics))
			return
		}
		t.Logf("\t\tShould respond with a request count of %d", resp.Requests)
		t.Logf("\t\tAnd with %d custom mark(s).", len(resp.CustomMetrics))
	}
}
