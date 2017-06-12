package speedcurve

import (
	"testing"
)

func TestGet(t *testing.T) {
	t.Log("Given the need to get a deploy from Speedcurve.")
	t.Log("\tWhen requesting the latest deploy.")
	resp, _ := Getlatest()
	if resp.Status != "completed" {
		t.Errorf("\t\tShould have gotten status completed but got %s", resp.Status)
		return
	}
	t.Logf("\t\tShould respond with status %s", resp.Status)
}

func TestAdd(t *testing.T) {
	t.Log("Given the need to add a new deploy.")
	t.Log("\tWhen issuing a POST for site 11794")
	resp, _ := Add("11794", "Custom deploy via go cli", "")
	if resp.Status != "success" {
		t.Errorf("\t\tShould have gotten status success, but got %s", resp.Status)
		return
	}
	t.Log("\t\tShould respond successfully")
}
