package infrastructure

import (
	"github.com/marcuzy/pimonit/core"
	"time"
)

//
//
type inmemoryTimeSerias struct {
	data []*core.RangeItem
}

func NewInmemoryTimeSerias() core.TimeSeries {
	return &inmemoryTimeSerias{}
}

func (t *inmemoryTimeSerias) Add(value float64) {
	t.data = append(t.data, &core.RangeItem{
		Date:  time.Now(),
		Value: value,
	})
}

func (t *inmemoryTimeSerias) GetRange(from, to time.Time) []*core.RangeItem {
	var res []*core.RangeItem
	for _, item := range t.data {
		if item.Date.After(from) && item.Date.Before(to) {
			res = append(res, item)
		}
	}

	return res
}

func (t *inmemoryTimeSerias) Avg(from, to time.Time) float64 {
	rang := t.GetRange(from, to)
	sum := 0.0
	for _, r := range rang {
		sum = sum + r.Value
	}

	return sum / float64(len(rang))
}
