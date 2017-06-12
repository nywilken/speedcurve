package speedcurve

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

const uri = "https://api.speedcurve.com/v1/deploys"

type (
	testInfo struct {
		Test     string `json:"test"`
		Browser  string `json:"browser"`
		Region   string `json:"region"`
		Template int    `json:"template"`
	}

	deployInfo struct {
		DeployId      int        `json:"deploy_id"`
		Status        string     `json:"status"`
		Message       string     `json:"message"`
		Note          string     `json:"note"`
		Details       string     `json:"detail"`
		TestCompleted []testInfo `json:"tests-completed"`
		TestRemaing   []testInfo `json:"test-remaining"`
	}

	deployResponse struct {
		DeployId int    `json:"deploy_id"`
		SiteId   int    `json:"site_id"`
		Status   string `json:"status"`
		Message  string `json:"message"`
	}
)

var (
	apiToken string
	client   *http.Client
)

func init() {
	t, ok := os.LookupEnv("SPD_API_TOKEN")
	if ok != true {
		log.Fatalln("Unable to find Speedcurve API token.")
	}
	apiToken = t
	client = &http.Client{}
}

func Add(site, note, details string) (deployResponse, error) {

	payload := url.Values{}
	payload.Add("site_id", site)
	payload.Add("note", note)
	payload.Add("details", details)

	req, _ := http.NewRequest("POST", uri, bytes.NewBufferString(payload.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(apiToken, "x")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer resp.Body.Close()

	var r deployResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
		return deployResponse{}, err
	}

	return r, nil
}

func Get(id string) (deployInfo, error) {

	rs := "/" + id
	req, _ := http.NewRequest("GET", uri+rs, nil)
	req.SetBasicAuth(apiToken, "x")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer resp.Body.Close()

	var d deployInfo
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Println("ERROR:", "Failed to decode JSON deploy response")
		return deployInfo{}, err
	}

	return d, nil
}

func Getlatest() (deployInfo, error) {
	return Get("latest")
}
