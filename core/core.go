package core

import (
	"github.com/marcuzy/pimonit/core/actions"
	"github.com/marcuzy/pimonit/core/interfaces"
	"github.com/marcuzy/pimonit/core/tasks"
)

type Core struct {
	*actions.Actions
	*tasks.Tasks
}

func Init(ts interfaces.TimeSeries, charts interfaces.ChartsGenerator, s interfaces.TempSensor) *Core {
	return &Core{
		actions.Init(ts, charts),
		tasks.Init(s, ts),
	}
}
