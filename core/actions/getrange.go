package actions

import (
	"github.com/marcuzy/pimonit/core"
	"time"
)

func GetRange(ts core.TimeSeries, from, to time.Time) ([]*core.RangeItem, error) {
	items := ts.GetRange(from, to)
	return items, nil
}
