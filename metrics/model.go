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

// Base Metric interface (optional, useful for handling generically)
type Metric interface {
	GetName() string
}

// ActiveEnergyMetric struct
type ActiveEnergyMetric struct {
	Name  string             `json:"name"`
	Units string             `json:"units"`
	Data  []ActiveEnergyData `json:"data"`
}

func (aem ActiveEnergyMetric) GetName() string {
	return aem.Name
}

type ActiveEnergyData struct {
	Date   CustomTime `json:"date"`
	Qty    float64    `json:"qty"`
	Source string     `json:"source,omitempty"`
}

// BasalEnergyBurnedMetric struct
type BasalEnergyBurnedMetric struct {
	Name  string            `json:"name"`
	Units string            `json:"units"`
	Data  []BasalEnergyData `json:"data"`
}

func (beb BasalEnergyBurnedMetric) GetName() string {
	return beb.Name
}

type BasalEnergyData struct {
	Date   CustomTime `json:"date"`
	Qty    float64    `json:"qty"`
	Source string     `json:"source,omitempty"`
}
