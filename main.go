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
	c := core.Init(ts, ch, s)

	if err := api.StartSever(8081, c); err != nil {
		panic(err)
	}
}
