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
	timeString := string(b[1 : len(b)-1])
	parsedTime, err := time.Parse(customTimeFormat, timeString)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
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
type EnergyMetric struct {
	Name   string     `json:"name"`
	Units  string     `json:"units"`
	Date   CustomTime `json:"date"`
	Qty    float64    `json:"qty"`
	Source string     `json:"source,omitempty"`
}

// GetName method to satisfy the Metric interface
func (em EnergyMetric) GetName() string {
	return em.Name
}
