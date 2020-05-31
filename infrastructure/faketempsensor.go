package infrastructure

import (
	"github.com/marcuzy/pimonit/core/interfaces"
	"math/rand"
)

type (
	fakeTempSensor struct {
	}
)

func NewFakeTempSensor() interfaces.TempSensor {
	return &fakeTempSensor{}
}

func (_ fakeTempSensor) CurrentTemperature(units interfaces.TemperatureUnits) (float64, error) {
	return rand.Float64(), nil
}
