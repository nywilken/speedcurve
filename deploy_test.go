package speedcurve

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Log("Given the need to get a deploy from Speedcurve.")
	t.Log("\tWhen requesting the latest deploy.")
	resp, _ := Get()
	if resp.Status != "completed" {
		t.Errorf("\t\tShould have gotten status completed but got %v", resp.Status)
		return
	}
	t.Logf("\t\tShould respond with status %s", resp.Status)
}
