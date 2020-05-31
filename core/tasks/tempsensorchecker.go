package tasks

import (
	"github.com/marcuzy/pimonit/core/interfaces"
	"log"
	"time"
)

type (
	TempSensorChecker struct {
		sensor interfaces.TempSensor
		ts     interfaces.TimeSeries
	}
)

func NewTempSensorChecker(s interfaces.TempSensor, ts interfaces.TimeSeries) *TempSensorChecker {
	return &TempSensorChecker{s, ts}
}

func (t *TempSensorChecker) Run() {
	go func() {
		for {
			temp, err := t.sensor.CurrentTemperature(interfaces.TemperatureUnitsCelsius)
			if err == nil {
				t.ts.Add(temp)
			} else {
				log.Print(err)
			}
			time.Sleep(time.Second * 1)
		}
	}()
}
