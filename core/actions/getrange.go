package actions

import (
	"github.com/marcuzy/pimonit/core/interfaces"
	"time"
)

type GetRange func(from, to time.Time) ([]*interfaces.RangeItem, error)

func NewGetRange(ts interfaces.TimeSeries) GetRange {
	return func(from, to time.Time) ([]*interfaces.RangeItem, error) {
		items := ts.GetRange(from, to)
		return items, nil
	}
}
