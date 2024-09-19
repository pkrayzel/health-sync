package domain

import (
	"testing"
	"time"

	"github.com/pkrayzel/health-sync-api/metrics"
	"github.com/stretchr/testify/assert"
)

// TestCalculateAverageCalories tests the average calorie calculation over different periods.
func TestCalculateAverageCalories(t *testing.T) {
	// Prepare fake data for active energy and basal energy using the new EnergyMetric structure
	fakeMetrics := []metrics.Metric{
		metrics.EnergyMetric{
			Name:  "active_energy",
			Units: "kcal",
			Date:  metrics.CustomTime{Time: time.Now().AddDate(0, 0, -1)},
			Qty:   2000,
		},
		metrics.EnergyMetric{
			Name:  "basal_energy_burned",
			Units: "kcal",
			Date:  metrics.CustomTime{Time: time.Now().AddDate(0, 0, -1)},
			Qty:   2400, // Convert from kJ to kcal
		},
	}

	// Test for last week
	avgTotalWeek, avgActiveWeek, avgBasalWeek := CalculateAverageLastWeek(fakeMetrics)
	assert.InDelta(t, 4400, avgTotalWeek, 1, "Average total calories per day (last week) should be correct")
	assert.InDelta(t, 2000, avgActiveWeek, 1, "Average active calories per day (last week) should be correct")
	assert.InDelta(t, 2400, avgBasalWeek, 1, "Average basal calories per day (last week) should be correct")

	// Test for last 2 weeks
	avgTotalTwoWeeks, avgActiveTwoWeeks, avgBasalTwoWeeks := CalculateAverageLastTwoWeeks(fakeMetrics)
	assert.InDelta(t, 4400, avgTotalTwoWeeks, 1, "Average total calories per day (last 2 weeks) should be correct")
	assert.InDelta(t, 2000, avgActiveTwoWeeks, 1, "Average active calories per day (last 2 weeks) should be correct")
	assert.InDelta(t, 2400, avgBasalTwoWeeks, 1, "Average basal calories per day (last 2 weeks) should be correct")

	// Test for last month
	avgTotalMonth, avgActiveMonth, avgBasalMonth := CalculateAverageLastMonth(fakeMetrics)
	assert.InDelta(t, 4400, avgTotalMonth, 1, "Average total calories per day (last month) should be correct")
	assert.InDelta(t, 2000, avgActiveMonth, 1, "Average active calories per day (last month) should be correct")
	assert.InDelta(t, 2400, avgBasalMonth, 1, "Average basal calories per day (last month) should be correct")
}
