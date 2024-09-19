package metrics

import (
	"encoding/json"
	"errors"
	"log"
)

// ParsePayload takes a payload in a map format and parses it into a slice of metrics.
func ParsePayload(payload map[string]interface{}) ([]Metric, error) {
	var metrics []Metric

	for _, metricData := range payload["data"].(map[string]interface{})["metrics"].([]interface{}) {
		metricJson, err := json.Marshal(metricData)
		if err != nil {
			return nil, errors.New("failed to marshal metric data")
		}

		var energyMetric EnergyMetric
		if err := json.Unmarshal(metricJson, &energyMetric); err != nil {
			return nil, err
		}

		// Only keep the metrics we are interested in (active_energy, basal_energy_burned)
		switch energyMetric.Name {
		case "active_energy", "basal_energy_burned":
			// For basal energy, convert kJ to kcal
			if energyMetric.Name == "basal_energy_burned" {
				energyMetric.Qty = energyMetric.Qty / 4.184
			}
			metrics = append(metrics, energyMetric)
		default:
			// Log unsupported metrics and skip them
			log.Printf("Skipping unsupported metric: %s", energyMetric.Name)
			continue
		}
	}

	return metrics, nil
}
