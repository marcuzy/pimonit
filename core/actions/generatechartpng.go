package actions

import (
	"bytes"
	"github.com/marcuzy/pimonit/core"
	"time"
)

func GenerateChartPNG(ts core.TimeSeries, charts core.ChartsGenerator, from, to time.Time) (*bytes.Buffer, error) {
	items := ts.GetRange(from, to)
	x := make([]time.Time, len(items))
	y := make([]float64, len(items))
	for i, item := range items {
		x[i] = item.Date
		y[i] = item.Value
	}
	return charts.Timeseries(x, y)
}
