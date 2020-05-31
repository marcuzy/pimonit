package core

import "time"

type (
	RangeItem struct {
		Value float64
		Date  time.Time
	}
	TimeSeries interface {
		Add(value float64)
		GetRange(from, to time.Time) []*RangeItem
		Avg(from, to time.Time) float64
		// Sum
	}
)
