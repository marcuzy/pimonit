package infrastructure

import (
	"bytes"
	"github.com/marcuzy/pimonit/core/interfaces"
	"github.com/wcharczuk/go-chart"
	"time"
)

type (
	chartGen struct {
	}
)

func NewChartGen() interfaces.ChartsGenerator {
	return &chartGen{}
}

func (c chartGen) Timeseries(x []time.Time, y []float64) (*bytes.Buffer, error) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)

	return buffer, err
}
