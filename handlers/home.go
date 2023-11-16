// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gocondor/core"
)

// Show home page
func WelcomeHome(c *core.Context) *core.Response {
	message := "{\"message\": \"Welcome to GoCondor\"}"
	return c.Response.Json(message)
}

// Show dashboard
func WelcomeToDashboard(c *core.Context) *core.Response {
	message := "{\"message\": \"Welcome to Dashboard\"}"
	return c.Response.Json(message)
}
