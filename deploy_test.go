package speedcurve

import (
	"testing"
)

var deploy *Deploy

func init() {
	config := &Config{}
	client := NewClient(config)
	deploy = NewDeploy(client)
}

func TestDeployGet(t *testing.T) {

	var deployGetCases = []struct {
		id       string
		desc     string
		expected string
	}{
		{"0", "a non-existent id", "no such deployment"},
		{"latest", "the latest deploy", "completed"},
		{"91088", "deploy id 91088", "completed"},
	}

	t.Log("Given the need to get a deploy from Speedcurve")
	for _, tc := range deployGetCases {
		t.Logf("\tWhen requesting deploy details for %s", tc.desc)
		resp, _ := deploy.Get(tc.id)
		if resp.Status != tc.expected {
			t.Errorf("\t\tShould have gotten status %s but got %s.", tc.expected, resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}

func TestDeployAdd(t *testing.T) {
	var deployAddCases = []struct {
		site     string
		note     string
		details  string
		desc     string
		expected string
	}{
		{"bogus-site-id", "", "", "a non-existent site id", "failure"},
	}

	t.Log("Given the need to add a new deploy.")
	for _, tc := range deployAddCases {
		t.Logf("\tWhen adding a new deploy for %s", tc.site)
		resp, _ := deploy.Add(tc.site, tc.note, tc.details)
		if resp.Status != tc.expected {
			t.Errorf("\t\tShould have gotten status completed but got %s.", resp.Status)
			return
		}
		t.Logf("\t\tShould respond with status %s.", resp.Status)
	}
}
