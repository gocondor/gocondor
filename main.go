// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/gincoat/gincoat/config"
	"github.com/gincoat/gincoat/core"
	"github.com/gincoat/gincoat/core/cache"
	"github.com/gincoat/gincoat/core/database"
	"github.com/gincoat/gincoat/core/pkgintegrator"
	"github.com/gincoat/gincoat/httpd"
	"github.com/gincoat/gincoat/httpd/middlewares"
	"github.com/gincoat/gincoat/integrations"
	"github.com/gincoat/gincoat/models"
	"github.com/joho/godotenv"
)

func main() {
	// New initializes new App variable
	app := core.New()

	// set env
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.SetEnv(env)

	// set the app mode
	app.SetAppMode(os.Getenv("APP_MODE"))

	// What features to turn on or off
	app.FeaturesControl(config.Features)

	// initialize core packages
	app.Bootstrap()

	//register database driver
	if app.Features.Database == true {
		pkgintegrator.Resolve().Integrate(core.GORMIntegrator(database.Resolve()))
	}

	//register the cache
	if app.Features.Cache == true {
		pkgintegrator.Resolve().Integrate(core.Cache(cache.Resolve()))
	}

	// Register packages integrations
	integrations.RegisterPKGIntegrations()

	// Register global middlewares
	middlewares.RegisterMiddlewares()

	// Register routes
	httpd.RegisterRoutes()

	//auto migrate tables
	if config.Features.Database == true {
		models.MigrateDB()
	}

	// Run App
	app.Run(os.Getenv("APP_HTTP_PORT"))
}
