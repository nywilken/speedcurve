// Copyright © 2017 Wilken Rivera

package speedcurve

import (
	"encoding/json"
)

type testinfo struct {
	Test     string `json:"test"`
	Browser  string `json:"browser"`
	Template int    `json:"template"`
}

// Deploy represents the response obtained when issuing a POST to
// Speedcurve's deploys API endpoint.
type Deploy struct {
	DeployID int64  `json:"deploy_id"`
	SiteID   int64  `json:"site_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	Info     struct {
		ScheduledTests []testinfo `json:"tests-added"`
	} `json:"info"`
	TestsRequested int `json:"tests-requested"`
}

// DeployDetails represents the response details when issuing a GET to
// Speedcurve's deploys API endpoint.
type DeployDetails struct {
	DeployID       int64      `json:"deploy_id"`
	SiteID         int64      `json:"site_id"`
	Status         string     `json:"status"`
	Note           string     `json:"note"`
	Detail         string     `json:"detail"`
	TestsCompleted []testinfo `json:"tests-completed"`
	TestsRemaining []testinfo `json:"tests-remaining"`
}

func (d Deploy) String() string {
	out, _ := jsonOutput(d)
	return string(out)
}

func (d DeployDetails) String() string {
	out, _ := jsonOutput(d)
	return string(out)
}

func jsonOutput(in interface{}) ([]byte, error) {
	o, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	return o, nil
}
