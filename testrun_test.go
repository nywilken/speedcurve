package speedcurve

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockserver() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {

		var code int
		var resp string

		switch r.URL.Path {
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
			w.WriteHeader(200)
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, resp)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

var tr *TestRun

func TestGetTestRun(t *testing.T) {
	server := mockserver()
	defer server.Close()

	config := &Config{server.URL, "x"}
	client := NewClient(config)
	tr = NewTestRun(client)

	var getTestRunCases = []struct {
		id               string
		desc             string
		expectedRequests int
		expectedMarks    int
	}{
		{"0", "a non-existent test run", 0, 0},
		{"117", "a completed test run", 56, 0},
		{"989", "a completed test run with custom marks", 56, 2},
	}

	t.Log("Given the need to get a test run from Speedcurve")
	for _, tc := range getTestRunCases {
		t.Logf("\tWhen requesting test run details for %s", tc.desc)
		resp, _ := tr.Get(tc.id)
		if resp.Requests != tc.expectedRequests {
			t.Errorf("\t\tShould have a request count of %d, but got %d.", tc.expectedRequests, resp.Requests)
			return
		}
		if len(resp.Marks) != tc.expectedMarks {
			t.Errorf("\t\tShould have %d custom marks, but got %d.", tc.expectedMarks, len(resp.Marks))
			return
		}
		t.Logf("\t\tShould respond with a request count of %d", resp.Requests)
		t.Logf("\t\tAnd with %d custom mark(s).", len(resp.Marks))
	}
}
