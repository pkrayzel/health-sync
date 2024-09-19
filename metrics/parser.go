package metrics

import (
	"encoding/json"
	"errors"
	"log"
)

func ParsePayload(payload map[string]interface{}) ([]Metric, error) {
	var metrics []Metric
	for _, metricData := range payload["data"].(map[string]interface{})["metrics"].([]interface{}) {
		metricJson, err := json.Marshal(metricData)
		if err != nil {
			return nil, errors.New("failed to marshal metric data")
		}

		// Parse each metric based on its "name" field
		var metric Metric
		switch metricData.(map[string]interface{})["name"] {
		case "active_energy":
			var activeEnergyMetric ActiveEnergyMetric
			if err := json.Unmarshal(metricJson, &activeEnergyMetric); err != nil {
				return nil, err
			}
			metric = activeEnergyMetric

		case "basal_energy_burned":
			var basalEnergyMetric BasalEnergyBurnedMetric
			if err := json.Unmarshal(metricJson, &basalEnergyMetric); err != nil {
				return nil, err
			}
			metric = basalEnergyMetric

		default:
			// Skip unsupported metrics but log them
			log.Printf("Skipping unknown metric: %s", metricData.(map[string]interface{})["name"])
			continue
		}

		metrics = append(metrics, metric)
	}
	return metrics, nil
}
