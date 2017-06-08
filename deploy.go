package speedcurve

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func Get() (deployInfo, error) {
	u := fmt.Sprintf("%s/latest", uri)
	rx, _ := http.NewRequest("GET", u, nil)
	rx.SetBasicAuth(apiToken, "x")
	resp, err := client.Do(rx)
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
