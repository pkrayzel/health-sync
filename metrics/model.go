package metrics

import (
	"log"
	"time"
)

// CustomTime for handling time in the format "2024-09-12 00:00:00 +0200"
type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05 -0700"

// UnmarshalJSON for CustomTime to parse custom time format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	timeString := string(b[1 : len(b)-1]) // Remove surrounding quotes
	parsedTime, err := time.Parse(customTimeFormat, timeString)
	if err != nil {
		log.Printf("Error parsing time: %s, error: %v", timeString, err) // Log the error
		return err
	}
	ct.Time = parsedTime
	return nil
}

// Metric interface for generic handling of different metrics
type Metric interface {
	GetName() string
}

// Flattened structure for ActiveEnergy and BasalEnergyBurned metrics
type EnergyMetricHolder struct {
	Name  string `json:"name"`
	Units string `json:"units"`
	Data  []struct {
		Date   *CustomTime `json:"date"`
		Qty    float64     `json:"qty"`
		Source string      `json:"source,omitempty"`
	} `json:"data"`
}

// GetName method to satisfy the Metric interface
func (em EnergyMetricHolder) GetName() string {
	return em.Name
}

// Flattened structure for a single EnergyMetric entry
type EnergyMetric struct {
	Name   string
	Units  string
	Date   time.Time
	Qty    float64
	Source string
}

// Constructor to create a flattened EnergyMetric from EnergyMetricHolder
func NewEnergyMetric(name string, units string, date time.Time, qty float64, source string) EnergyMetric {
	return EnergyMetric{
		Name:   name,
		Units:  units,
		Date:   date,
		Qty:    qty,
		Source: source,
	}
}

// GetName method to satisfy the Metric interface
func (em EnergyMetric) GetName() string {
	return em.Name
}
