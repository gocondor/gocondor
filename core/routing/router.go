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

// Resolve resolves an already initiated router
func (r *Router) Resolve() *Router {
	return router
}

// Get is a definition for get request
func (r *Router) Get(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "get",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

// Post is a definition for post request
func (r *Router) Post(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "post",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

// Delete is a definition for delete request
func (r *Router) Delete(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "delete",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

// Put is a definition for put request
func (r *Router) Put(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "put",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

// Options is a definition for options request
func (r *Router) Options(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "options",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

// Head is a definition for head request
func (r *Router) Head(path string, handlers ...gin.HandlerFunc) *Router {
	r.routes = append(r.routes, Route{
		Method:   "head",
		Path:     path,
		Handlers: handlers,
	})

	return r
}

//GetRoutes returns all Defined routes
func (r *Router) GetRoutes() []Route {
	return r.routes
}
