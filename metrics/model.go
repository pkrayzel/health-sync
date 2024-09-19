package metrics

import "time"

// Root structure
type HealthData struct {
	Data MetricsContainer `json:"data"`
}

// Container for all metrics
type MetricsContainer struct {
	Metrics []Metric `json:"metrics"`
}

// Base Metric interface (optional, useful if you want to handle all metrics generically)
type Metric interface {
	GetName() string
}

// Specific Metric Types

// StandTimeMetric struct
type StandTimeMetric struct {
	Name  string           `json:"name"`
	Units string           `json:"units"`
	Data  []StandTimeEntry `json:"data"`
}

type StandTimeEntry struct {
	Date   time.Time `json:"date"`
	Qty    int       `json:"qty"`
	Source string    `json:"source,omitempty"`
}

func (s StandTimeMetric) GetName() string {
	return s.Name
}

// HeartRateMetric structure
type HeartRateMetric struct {
	Name  string          `json:"name"`
	Units string          `json:"units"`
	Data  []HeartRateData `json:"data"`
}

func (h HeartRateMetric) GetName() string {
	return h.Name
}

// RestingHeartRateMetric struct
type RestingHeartRateMetric struct {
	Name  string                 `json:"name"`
	Units string                 `json:"units"`
	Data  []RestingHeartRateData `json:"data"`
}

type RestingHeartRateData struct {
	Date time.Time `json:"date"`
	Qty  int       `json:"qty"`
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
	Date   time.Time `json:"date"`
	Qty    float64   `json:"qty"`
	Source string    `json:"source,omitempty"`
}

type HeartRateData struct {
	Date   time.Time `json:"date"`
	Avg    float64   `json:"Avg"`
	Max    int       `json:"Max"`
	Min    int       `json:"Min"`
	Source string    `json:"source,omitempty"`
}

// SleepMetric struct
type SleepMetric struct {
	Name  string         `json:"name"`
	Units string         `json:"units"`
	Data  []SleepSession `json:"data"`
}

type SleepSession struct {
	Date       time.Time `json:"date"`
	Asleep     float64   `json:"asleep"`
	Awake      float64   `json:"awake"`
	Rem        float64   `json:"rem"`
	Deep       float64   `json:"deep"`
	Core       float64   `json:"core"`
	InBed      float64   `json:"inBed"`
	InBedStart time.Time `json:"inBedStart"`
	InBedEnd   time.Time `json:"inBedEnd"`
	SleepStart time.Time `json:"sleepStart"`
	SleepEnd   time.Time `json:"sleepEnd"`
	Source     string    `json:"source"`
}

// StepCountMetric struct
type StepCountMetric struct {
	Name  string          `json:"name"`
	Units string          `json:"units"`
	Data  []StepCountData `json:"data"`
}

type StepCountData struct {
	Date   time.Time `json:"date"`
	Qty    int       `json:"qty"`
	Source string    `json:"source,omitempty"`
}

// Define additional specific metrics here...

// Example usage
func main() {
	// Load and process your data
}
