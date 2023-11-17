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

	// register your event here ...
	eventsManager.Register(events.USER_REGISTERED, eventjobs.SendWelcomeEmail)
	eventsManager.Register(events.USER_REGISTERED, eventjobs.TestEvent)
	eventsManager.Register(events.USER_PASSWORD_RESET_REQUESTED, eventjobs.SendResetPasswordEmail)
	eventsManager.Register(events.PASSWORD_CHANGED, eventjobs.SendPasswordChangedEmail)
}
