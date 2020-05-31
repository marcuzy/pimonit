package main

import (
	"github.com/marcuzy/pimonit/api"
	"github.com/marcuzy/pimonit/core"
	"github.com/marcuzy/pimonit/infrastructure"
)

//
//
func main() {
	ts := infrastructure.NewInmemoryTimeSerias()
	ch := infrastructure.NewChartGen()
	//s := infrastructure.NewPiTempSensor()
	s := infrastructure.NewFakeTempSensor()
	checker := core.NewTempSensorChecker(s, ts)
	checker.Run()
	if err := api.StartSever(8081, ts, ch); err != nil {
		panic(err)
	}
}
