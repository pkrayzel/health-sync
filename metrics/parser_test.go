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

	// Verify that both active_energy and basal_energy_burned are present
	var foundActiveEnergy, foundBasalEnergy bool
	for _, metric := range metricsData {
		switch metric.GetName() {
		case "active_energy":
			foundActiveEnergy = true
			assert.NotNil(t, metric, "Active energy metric should be valid")
		case "basal_energy_burned":
			foundBasalEnergy = true
			assert.NotNil(t, metric, "Basal energy metric should be valid")
			// Check that conversion from kJ to kcal happened
			assert.Less(t, metric.(EnergyMetric).Qty, 4000.0, "Basal energy should be less than 4000 kcal after conversion")
		default:
			t.Errorf("Unexpected metric found: %s", metric.GetName())
		}
	}

	assert.True(t, foundActiveEnergy, "Active energy metric should be found in parsed data")
	assert.True(t, foundBasalEnergy, "Basal energy metric should be found in parsed data")
}
