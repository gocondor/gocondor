// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/handlers"
)

// Register the app routes
func registerRoutes() {
	router := core.ResolveRouter()
	//#############################
	//# App Routes            #####
	//#############################

	// Define your routes here...
	router.Get("/", handlers.WelcomeHome)
	// Uncomment the lines below to enable authentication
	// router.Post("/signup", handlers.Signup)
	// router.Post("/signin", handlers.Signin)
	// router.Post("/signout", handlers.Signout)
	// router.Post("/reset-password", handlers.ResetPasswordRequest)
	// router.Post("/reset-password/code/:code", handlers.SetNewPassword)
	// router.Get("/dashboard", handlers.WelcomeToDashboard, middlewares.AuthCheck)
}
