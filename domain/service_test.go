package domain

import (
	"testing"
	"time"

	"github.com/pkrayzel/health-sync-api/metrics"
	"github.com/stretchr/testify/assert"
)

// TestCalculateAverageCalories tests the average calorie calculation over different periods.
func TestCalculateAverageCalories(t *testing.T) {
	// Prepare fake data for active energy and basal energy
	fakeMetrics := []metrics.Metric{
		metrics.ActiveEnergyMetric{
			Name:  "active_energy",
			Units: "kcal",
			Data: []metrics.ActiveEnergyData{
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -1)}, Qty: 2000},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -2)}, Qty: 2200},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -3)}, Qty: 1800},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -10)}, Qty: 2500},
			},
		},
		metrics.BasalEnergyBurnedMetric{
			Name:  "basal_energy_burned",
			Units: "kJ",
			Data: []metrics.BasalEnergyData{
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -1)}, Qty: 8000},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -2)}, Qty: 8500},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -3)}, Qty: 7500},
				{Date: metrics.CustomTime{time.Now().AddDate(0, 0, -10)}, Qty: 9000},
			},
		},
	}

	// Test for last week
	avgTotalWeek, avgActiveWeek, avgBasalWeek := CalculateAverageLastWeek(fakeMetrics)
	assert.InDelta(t, 2616, avgTotalWeek, 1, "Average total calories per day (last week) should be correct")
	assert.InDelta(t, 2000, avgActiveWeek, 1, "Average active calories per day (last week) should be correct")
	assert.InDelta(t, 616, avgBasalWeek, 1, "Average basal calories per day (last week) should be correct")

	// Test for last 2 weeks
	avgTotalTwoWeeks, avgActiveTwoWeeks, avgBasalTwoWeeks := CalculateAverageLastTwoWeeks(fakeMetrics)
	assert.InDelta(t, 2704, avgTotalTwoWeeks, 1, "Average total calories per day (last 2 weeks) should be correct")
	assert.InDelta(t, 2125, avgActiveTwoWeeks, 1, "Average active calories per day (last 2 weeks) should be correct")
	assert.InDelta(t, 579, avgBasalTwoWeeks, 1, "Average basal calories per day (last 2 weeks) should be correct")

	// Test for last month
	avgTotalMonth, avgActiveMonth, avgBasalMonth := CalculateAverageLastMonth(fakeMetrics)
	assert.InDelta(t, 2704, avgTotalMonth, 1, "Average total calories per day (last month) should be correct")
	assert.InDelta(t, 2125, avgActiveMonth, 1, "Average active calories per day (last month) should be correct")
	assert.InDelta(t, 579, avgBasalMonth, 1, "Average basal calories per day (last month) should be correct")
}
