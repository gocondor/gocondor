// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"path/filepath"

	"github.com/gocondor/core"
	"github.com/gocondor/core/env"
	"github.com/gocondor/core/logger"
	"github.com/gocondor/gocondor/config"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// The main function
func main() {
	app := core.New()
	// Handle the reading of the .env file
	if config.GetEnvFileConfig().UseDotEnvFile {
		envVars, err := godotenv.Read(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		env.SetEnvVars(envVars)
	}
	// Handle the logs
	path, err := filepath.Abs("./logs/app.log")
	os.MkdirAll(filepath.Dir(path), 644)
	if err != nil {
		panic(err)
	}
	app.SetLogsDriver(&logger.LogFileDriver{
		FilePath: path,
	})
	app.SetRequestConfig(config.GetRequestConfig())
	app.SetGormConfig(config.GetGormConfig())
	app.SetCacheConfig(config.GetCacheConfig())
	app.Bootstrap()
	registerGlobalMiddlewares()
	registerRoutes()
	runAutoMigrations()
	app.Run(httprouter.New())
}
