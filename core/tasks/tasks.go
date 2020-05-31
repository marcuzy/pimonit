package tasks

import (
	"github.com/marcuzy/pimonit/core/interfaces"
)

type Tasks struct {
	*TempSensorChecker
}

func Init(s interfaces.TempSensor, ts interfaces.TimeSeries) *Tasks {
	t := &Tasks{
		NewTempSensorChecker(s, ts),
	}

	t.TempSensorChecker.Run()

	return t
}
