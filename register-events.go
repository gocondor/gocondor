// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/events"
	eventjobs "github.com/gocondor/gocondor/events/jobs"
)

// Register events
func registerEvents() {
	eventsManager := core.ResolveEventsManager()
	//########################################
	//#      events registration         #####
	//########################################

	// register your event here here ...
	eventsManager.Register(events.USER_REGISTERED, eventjobs.SendEmail)
	eventsManager.Register(events.USER_REGISTERED, eventjobs.TestEvent)
}
