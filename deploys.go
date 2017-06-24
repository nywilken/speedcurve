package speedcurve

type testinfo struct {
	Test     string `json:"test"`
	Browser  string `json:"browser"`
	Template int    `json:"template"`
}

//DeployDetails ...
type DeployDetails struct {
	DeployID       int64      `json:"deploy_id"`
	SiteID         int64      `json:"site_id"`
	Status         string     `json:"status"`
	Note           string     `json:"note"`
	Detail         string     `json:"detail"`
	TestsCompleted []testinfo `json:"tests-completed"`
	TestsRemaining []testinfo `json:"tests-remaining"`
}

//DeployResponse ...
type DeployResponse struct {
	DeployID int64  `json:"deploy_id"`
	SiteID   int64  `json:"site_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	Info     struct {
		ScheduledTests []testinfo `json:"tests-added"`
	} `json:"info"`
	TestsRequested int `json:"tests-requested"`
}
