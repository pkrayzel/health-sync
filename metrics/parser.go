package metrics

import (
	"encoding/json"
	"errors"
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
		case "apple_stand_time":
			var standTimeMetric StandTimeMetric
			if err := json.Unmarshal(metricJson, &standTimeMetric); err != nil {
				return nil, err
			}
			metric = standTimeMetric
		case "heart_rate":
			var heartRateMetric HeartRateMetric
			if err := json.Unmarshal(metricJson, &heartRateMetric); err != nil {
				return nil, err
			}
			metric = heartRateMetric
		// Add more cases for other metric types
		default:
			return nil, errors.New("unknown metric type")
		}

		metrics = append(metrics, metric)
	}
	return metrics, nil
}
