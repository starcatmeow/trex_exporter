package trex

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestParsesTrexSummary_Api3_3(t *testing.T) {
	apiResponse, err := os.Open("../trex-api-responses/summary-api_3-3.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer apiResponse.Close()

	var parsedSummary summary

	err = json.NewDecoder(apiResponse).Decode(&parsedSummary)
	if err != nil {
		t.Error(err)
		return
	}

	expectedSummary := summary{
		Hashrate:      63829882,
		AcceptedCount: 583,
		RejectedCount: 0,
		Uptime:        36267,
		GpuTotal:      1,
		SolvedCount:   0,
		Gpus: []gpuSummary{
			{Hashrate: 63829882, FanSpeed: 75, Temperature: 44, Power: 202, DeviceId: 0, MemoryTemperature: 82, LHRTune: 73.90000000000003, Name: "RTX 3070 Ti"},
		},
	}

	if !cmp.Equal(parsedSummary, expectedSummary) {
		t.Error("Summary api not correctly parsed")
		t.Logf("Parsed   %+v", parsedSummary)
		t.Logf("Expected %+v", expectedSummary)
	}
}
