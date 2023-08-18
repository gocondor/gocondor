package eventjobs

import (
	"github.com/gocondor/core"
)

var TestEvent core.EventJob = func(event *core.Event, c *core.Context) {
	c.LogInfo("hello from event test job")
}
