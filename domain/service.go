package domain

import (
	"time"

	"github.com/pkrayzel/health-sync-api/metrics"
)

// Calculate the averages for active, basal, and total calories for a given time range
func CalculateAverageCalories(metricsData []metrics.Metric, startDate time.Time) (float64, float64, float64) {
	var totalActiveCalories, totalBasalCalories float64
	dayCount := make(map[string]struct{}) // to track unique days

	// Iterate over metrics and sum active and basal calories
	for _, metric := range metricsData {
		m := metric.(metrics.EnergyMetric)
		if m.Date.After(startDate) {
			dayCount[m.Date.Format("2006-01-02")] = struct{}{} // Track unique day

			if m.Name == "active_energy" {
				totalActiveCalories += m.Qty
			} else if m.Name == "basal_energy_burned" {
				totalBasalCalories += m.Qty
			}
		}
	}

	// Determine the number of unique days
	numDays := float64(len(dayCount))

	// Return 0 if no days were processed
	if numDays == 0 {
		return 0, 0, 0
	}

	// Calculate average calories per day
	averageTotalCalories := (totalActiveCalories + totalBasalCalories) / numDays
	averageActiveCalories := totalActiveCalories / numDays
	averageBasalCalories := totalBasalCalories / numDays

	return averageTotalCalories, averageActiveCalories, averageBasalCalories
}

// Wrapper functions for different time ranges
func CalculateAverageLastMonth(metricsData []metrics.Metric) (float64, float64, float64) {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return CalculateAverageCalories(metricsData, oneMonthAgo)
}

func CalculateAverageLastTwoWeeks(metricsData []metrics.Metric) (float64, float64, float64) {
	twoWeeksAgo := time.Now().AddDate(0, 0, -14)
	return CalculateAverageCalories(metricsData, twoWeeksAgo)
}

func CalculateAverageLastWeek(metricsData []metrics.Metric) (float64, float64, float64) {
	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	return CalculateAverageCalories(metricsData, oneWeekAgo)
}
