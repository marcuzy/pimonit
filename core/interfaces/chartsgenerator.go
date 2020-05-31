package interfaces

import (
	"bytes"
	"time"
)

type (
	ChartsGenerator interface {
		Timeseries(x []time.Time, y []float64) (*bytes.Buffer, error)
	}
)
