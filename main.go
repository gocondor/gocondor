// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/harranali/gincoat/config"
	"github.com/harranali/gincoat/core"
	"github.com/harranali/gincoat/core/env"
	"github.com/harranali/gincoat/httpd"
	"github.com/harranali/gincoat/httpd/middlewares"
	"github.com/harranali/gincoat/integrations"
	"github.com/harranali/gincoat/models"
)

func main() {
	// Initiate app
	app := core.New()

	// What features to turn on or off
	app.FeaturesControl(config.Features)

	// Bootstrap dependencies
	app.Bootstrap()

	// Register packages integrations
	integrations.RegisterPKGIntegrations()

	// Register global middlewares
	middlewares.RegisterMiddlewares()

	// Register routes
	httpd.RegisterRoutes()

	if config.Features.Database == true {
		//auto migrate tables
		models.MigrateDB()
	}

	// Run App
	app.Run(env.Get("APP_HTTP_PORT"))
}
