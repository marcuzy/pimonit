package core

const (
	TemperatureUnitsCelsius TemperatureUnits = iota
	TemperatureUnitsFahrenheit
)

type (
	TemperatureUnits int
	TempSensor       interface {
		CurrentTemperature(units TemperatureUnits) (float64, error)
	}
)
