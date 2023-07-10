// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/handlers"
	"github.com/gocondor/gocondor/middlewares"
)

// Register the app routes
func registerRoutes() {
	router := core.ResolveRouter()
	//#############################
	//# App Routes            #####
	//#############################

	// Define your routes here...
	router.Get("/", handlers.HomeGet, middlewares.ExampleMiddleware)
}
