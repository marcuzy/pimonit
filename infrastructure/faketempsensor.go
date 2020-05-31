package infrastructure

import (
	"github.com/marcuzy/pimonit/core"
	"math/rand"
)

type (
	fakeTempSensor struct {
	}
)

func NewFakeTempSensor() core.TempSensor {
	return &fakeTempSensor{}
}

func (_ fakeTempSensor) CurrentTemperature(units core.TemperatureUnits) (float64, error) {
	return rand.Float64(), nil
}
