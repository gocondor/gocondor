// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/gocondor/core"
	"github.com/gocondor/core/cache"
	"github.com/gocondor/core/database"
	"github.com/gocondor/core/pkgintegrator"
	"github.com/gocondor/gocondor/config"
	"github.com/gocondor/gocondor/http"
	"github.com/gocondor/gocondor/http/handlers"
	"github.com/gocondor/gocondor/http/middlewares"
	"github.com/gocondor/gocondor/integrations"
	"github.com/gocondor/gocondor/models"
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
	app.SetEnabledFeatures(config.Features)

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

	// initiate handlers dependancies
	handlers.InitiateHandlersDependencies()

	// Register routes
	http.RegisterRoutes()

	//auto migrate tables
	if config.Features.Database == true {
		models.MigrateDB()
	}

	// Run App
	app.Run(os.Getenv("APP_HTTP_PORT"))
}
