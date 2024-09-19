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

		var energyMetricHolder EnergyMetricHolder
		if err := json.Unmarshal(metricJson, &energyMetricHolder); err != nil {
			return nil, err
		}

		// Only keep the metrics we are interested in
		switch energyMetricHolder.Name {
		case "active_energy", "basal_energy_burned":
			// Iterate over the Data slice and check the Date field
			for _, data := range energyMetricHolder.Data {
				if data.Date == nil {
					log.Printf("Skipping metric %s due to missing or invalid date", energyMetricHolder.Name)
					continue
				}

				// Convert kJ to kcal
				convertedQty := data.Qty * 0.239

				metrics = append(
					metrics,
					NewEnergyMetric(energyMetricHolder.Name, energyMetricHolder.Units, data.Date.Time, convertedQty, data.Source),
				)
			}
		default:
			log.Printf("Skipping unsupported metric: %s", energyMetricHolder.Name)
		}
	}
	return metrics, nil
}
