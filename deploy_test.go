package speedcurve

import (
	"testing"
)

func TestGet(t *testing.T) {
	var getTestCases = []struct {
		id       string
		desc     string
		expected string
	}{
		{"0", "a non-existent id", "no such deployment"},
		{"latest", "the latest deploy", "completed"},
		{"91088", "deploy id 91088", "completed"},
	}

	t.Log("Given the need to get a deploy from Speedcurve")
	for _, tc := range getTestCases {
		t.Logf("\tWhen requesting deploy details for %s", tc.desc)
		resp, _ := Get(tc.id)
		if resp.Status != tc.expected {
			t.Errorf("\t\tShould have gotten status %s but got %s.", tc.expected, resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}

func TestAdd(t *testing.T) {
	var addTestCases = []struct {
		site     string
		note     string
		details  string
		desc     string
		expected string
	}{
		{"BOGUS", "", "", "a non-existent site id", "failure"},
	}

	t.Log("Given the need to add a new deploy.")
	for _, tc := range addTestCases {
		t.Logf("\tWhen adding a new deploy for %s", tc.site)
		resp, _ := Add(tc.site, tc.note, tc.details)
		if resp.Status != tc.expected {
			t.Errorf("\t\tShould have gotten status completed but got %s.", resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}
