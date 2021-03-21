// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package routing

import (
	"github.com/gin-gonic/gin"
)

// Route ais struct describes a specific route
type Route struct {
	Method   string
	Path     string
	Handlers []gin.HandlerFunc
}

// Router handles routing
type Router struct {
	routes []Route
}

var router *Router

// New initiates new router
func New() *Router {
	router = &Router{
		[]Route{},
	}
	return router
}

// ResolveRouter resolves an already initiated router
func ResolveRouter() *Router {
	return router
}

// Get is a definition for get requests
func (r *Router) Get(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "get",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

//GetRoutes returns all Defined routes
func (r *Router) GetRoutes() []Route {
	return r.routes
}
