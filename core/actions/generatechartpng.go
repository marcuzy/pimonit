package actions

import (
	"bytes"
	"github.com/marcuzy/pimonit/core/interfaces"
	"time"
)

type GenerateChartPNG func(from, to time.Time) (*bytes.Buffer, error)

func NewGenerateChartPNG(ts interfaces.TimeSeries, charts interfaces.ChartsGenerator) GenerateChartPNG {
	return func(from, to time.Time) (*bytes.Buffer, error) {
		items := ts.GetRange(from, to)
		x := make([]time.Time, len(items))
		y := make([]float64, len(items))
		for i, item := range items {
			x[i] = item.Date
			y[i] = item.Value
		}
		return charts.Timeseries(x, y)
	}
}
