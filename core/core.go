// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gincoat/gincoat/core/cache"
	"github.com/gincoat/gincoat/core/database"
	"github.com/gincoat/gincoat/core/middlewaresengine"
	"github.com/gincoat/gincoat/core/pkgintegrator"
	"github.com/gincoat/gincoat/core/routing"
	"github.com/unrolled/secure"
)

// App struct
type App struct {
	Features *Features
}

// GORM is a const represents gorm variable name
const GORM = "gorm"

// CACHE a cache engine variable
const CACHE = "cache"

// logs file path
const logsFilePath = "logs/app.log"

// logs file
var logsFile *os.File

// New initiates the app struct
func New() *App {
	return &App{
		Features: &Features{},
	}
}

// SetEnv sets environment varialbes
func (app *App) SetEnv(env map[string]string) {
	for key, val := range env {
		os.Setenv(strings.TrimSpace(key), strings.TrimSpace(val))
	}
}

//Bootstrap initiate app
func (app *App) Bootstrap() {
	//initiate package integrator variable
	pkgintegrator.New()

	//initiate middlewares engine varialbe
	middlewaresengine.New()

	//initiate routing engine varialbe
	routing.New()

	//initiate data base varialb
	if app.Features.Database == true {
		database.New()
	}

	// initiate the cache varialbe
	if app.Features.Cache == true {
		cache.New()
	}
}

// Run execute the app
func (app *App) Run(portNumber string) {
	// fallback to port number to 80 if not set
	if portNumber == "" {
		portNumber = "80"
	}

	logsFile, err := os.OpenFile(logsFilePath, os.O_CREATE|os.O_APPEND|os.O_CREATE, 644)
	if err != nil {
		panic(err)
	}
	defer logsFile.Close()
	gin.DefaultWriter = io.MultiWriter(logsFile, os.Stdout)

	//initiate gin engines
	httpGinEngine := gin.Default()
	httpsGinEngine := gin.Default()

	httpsOn, _ := strconv.ParseBool(os.Getenv("APP_HTTPS_ON"))
	redirectToHTTPS, _ := strconv.ParseBool(os.Getenv("APP_REDIRECT_HTTP_TO_HTTPS"))

	if httpsOn {
		//serve the https
		httpsGinEngine = app.IntegratePackages(httpsGinEngine, pkgintegrator.Resolve().GetIntegrations())
		router := routing.Resolve()
		httpsGinEngine = app.RegisterRoutes(router.GetRoutes(), httpsGinEngine)
		certFile := os.Getenv("APP_HTTPS_CERT_FILE_PATH")
		keyFile := os.Getenv("APP_HTTPS_KEY_FILE_PATH")
		host := app.getHTTPSHost() + ":443"
		go httpsGinEngine.RunTLS(host, certFile, keyFile)
	}

	//redirect http to https
	if httpsOn && redirectToHTTPS {
		secureFunc := func() gin.HandlerFunc {
			return func(c *gin.Context) {
				secureMiddleware := secure.New(secure.Options{
					SSLRedirect: true,
					SSLHost:     app.getHTTPSHost() + ":443",
				})
				err := secureMiddleware.Process(c.Writer, c.Request)
				if err != nil {
					return
				}
				c.Next()
			}
		}()
		redirectEngine := gin.New()
		redirectEngine.Use(secureFunc)
		host := fmt.Sprintf("%s:%s", app.getHTTPHost(), portNumber)
		redirectEngine.Run(host)
	}

	//serve the http version
	httpGinEngine = app.IntegratePackages(httpGinEngine, pkgintegrator.Resolve().GetIntegrations())
	router := routing.Resolve()
	httpGinEngine = app.RegisterRoutes(router.GetRoutes(), httpGinEngine)
	host := fmt.Sprintf("%s:%s", app.getHTTPHost(), portNumber)
	httpGinEngine.Run(host)
}

func (app *App) SetAppMode(mode string) {
	if mode == gin.ReleaseMode || mode == gin.TestMode || mode == gin.DebugMode {
		gin.SetMode(mode)
	} else {
		gin.SetMode(gin.TestMode)
	}
}

func (app *App) IntegratePackages(engine *gin.Engine, handlerFuncs []gin.HandlerFunc) *gin.Engine {
	for _, pkgIntegration := range handlerFuncs {
		engine.Use(pkgIntegration)
	}

	return engine
}

//FeaturesControl to control what features to turn on or off
func (app *App) SetEnabledFeatures(features *Features) {
	app.Features = features
}

func (app *App) UseMiddlewares(middlewares []gin.HandlerFunc, engine *gin.Engine) *gin.Engine {
	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	return engine
}

func (app *App) RegisterRoutes(routers []routing.Route, engine *gin.Engine) *gin.Engine {
	for _, route := range routers {
		switch route.Method {
		case "get":
			engine.GET(route.Path, route.Handlers...)
		case "post":
			engine.POST(route.Path, route.Handlers...)
		case "delete":
			engine.DELETE(route.Path, route.Handlers...)
		case "patch":
			engine.PATCH(route.Path, route.Handlers...)
		case "put":
			engine.PUT(route.Path, route.Handlers...)
		case "options":
			engine.OPTIONS(route.Path, route.Handlers...)
		case "head":
			engine.HEAD(route.Path, route.Handlers...)
		}
	}

	return engine
}

func (app *App) getHTTPSHost() string {
	host := os.Getenv("APP_HTTPS_HOST")
	//if not set get http instead
	if host == "" {
		host = os.Getenv("APP_HTTP_HOST")
	}
	//if both not set use local host
	if host == "" {
		host = "localhost"
	}
	return host
}

func (app *App) getHTTPHost() string {
	host := os.Getenv("APP_HTTP_HOST")
	//if both not set use local host
	if host == "" {
		host = "localhost"
	}
	return host
}
