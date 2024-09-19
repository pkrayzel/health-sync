package domain

import (
	"time"

	"github.com/pkrayzel/health-sync-api/metrics"
)

// Calculate the total calories per day: active + basal (converted to kcal)
func calculateTotalCaloriesPerDay(activeEnergy float64, basalEnergy float64) float64 {
	// Convert basal_energy from kJ to kcal
	basalEnergyInKcal := basalEnergy / 4.184
	// Total calories is active_energy + basal_energy (converted to kcal)
	return activeEnergy + basalEnergyInKcal
}

// Calculate the averages for a given time range (1 week, 2 weeks, or 1 month)
func CalculateAverageCalories(metricsData []metrics.Metric, startDate time.Time) (float64, float64, float64) {
	var totalActiveCalories, totalBasalCalories float64
	var dayCount int

	// Iterate over metrics and sum calories for active and basal
	for _, metric := range metricsData {
		switch m := metric.(type) {
		case metrics.ActiveEnergyMetric:
			for _, data := range m.Data {
				if data.Date.After(startDate) {
					totalActiveCalories += data.Qty
					dayCount++
				}
			}
		case metrics.BasalEnergyBurnedMetric:
			for _, data := range m.Data {
				if data.Date.After(startDate) {
					totalBasalCalories += data.Qty
				}
			}
		}
	}

	// Calculate total calories
	totalCalories := calculateTotalCaloriesPerDay(totalActiveCalories, totalBasalCalories)

	// Average calories per day
	averageTotalCalories := totalCalories / float64(dayCount)
	averageActiveCalories := totalActiveCalories / float64(dayCount)
	averageBasalCalories := (totalBasalCalories / 4.184) / float64(dayCount)

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
