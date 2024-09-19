package metrics

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParsePayloadFromFile tests the parsing of a real production JSON payload.
func TestParsePayloadFromFile(t *testing.T) {
	// Load the JSON file
	file, err := os.Open("../tests/metrics.json")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the file content into a generic map
	var payload map[string]interface{}
	err = json.NewDecoder(file).Decode(&payload)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Call the parser with the loaded payload
	metricsData, err := ParsePayload(payload)
	if err != nil {
		t.Fatalf("Failed to parse payload: %v", err)
	}

	// Assert that the parsed data contains expected metrics
	assert.NotEmpty(t, metricsData, "Parsed metrics data should not be empty")

	// Verify that only active_energy and basal_energy_burned are present
	var foundActiveEnergy, foundBasalEnergy bool
	for _, metric := range metricsData {
		switch metric.GetName() {
		case "active_energy":
			foundActiveEnergy = true
			activeEnergyMetric := metric.(ActiveEnergyMetric)
			assert.Greater(t, len(activeEnergyMetric.Data), 0, "Active energy data should not be empty")
			assert.NotNil(t, activeEnergyMetric.Data[0].Qty, "Active energy should have a valid quantity")

		case "basal_energy_burned":
			foundBasalEnergy = true
			basalEnergyMetric := metric.(BasalEnergyBurnedMetric)
			assert.Greater(t, len(basalEnergyMetric.Data), 0, "Basal energy data should not be empty")
			assert.NotNil(t, basalEnergyMetric.Data[0].Qty, "Basal energy should have a valid quantity")

		default:
			t.Errorf("Unexpected metric found: %s", metric.GetName())
		}
	}

	assert.True(t, foundActiveEnergy, "Active energy metric should be found in parsed data")
	assert.True(t, foundBasalEnergy, "Basal energy metric should be found in parsed data")
}
