package core

import (
	"log"
	"time"
)

type (
	TempSensorChecker struct {
		sensor TempSensor
		ts     TimeSeries
	}
)

func NewTempSensorChecker(s TempSensor, ts TimeSeries) *TempSensorChecker {
	return &TempSensorChecker{s, ts}
}

func (t *TempSensorChecker) Run() {
	go func() {
		for {
			temp, err := t.sensor.CurrentTemperature(TemperatureUnitsCelsius)
			if err == nil {
				t.ts.Add(temp)
			} else {
				log.Print(err)
			}
			time.Sleep(time.Second * 1)
		}
	}()
}
