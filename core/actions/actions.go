package actions

import (
	"github.com/marcuzy/pimonit/core/interfaces"
)

type Actions struct {
	GenerateChartPNG
	GetRange
}

func Init(ts interfaces.TimeSeries, charts interfaces.ChartsGenerator) *Actions {
	return &Actions{
		NewGenerateChartPNG(ts, charts),
		NewGetRange(ts),
	}
}
