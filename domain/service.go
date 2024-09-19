package domain

import (
	"time"

	"github.com/pkrayzel/health-sync-api/metrics"
)

// CalculateAverageCalories calculates average daily calories for a given time period.
func CalculateAverageCalories(metricsData []metrics.Metric) float64 {
	totalCalories := 0.0
	days := 0

	for _, metric := range metricsData {
		if activeEnergyMetric, ok := metric.(metrics.ActiveEnergyMetric); ok {
			for _, entry := range activeEnergyMetric.Data {
				if entry.Date.After(time.Now().AddDate(0, -1, 0)) { // Last 30 days
					totalCalories += entry.Qty
					days++
				}
			}
		}
	}

	if days == 0 {
		return 0
	}
	return totalCalories / float64(days)
}
