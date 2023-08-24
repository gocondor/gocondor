package eventjobs

import (
	"github.com/gocondor/core"
)

var TestEvent core.EventJob = func(event *core.Event, c *core.Context) {
	c.GetLogger().Info("hello from event test job")
}
