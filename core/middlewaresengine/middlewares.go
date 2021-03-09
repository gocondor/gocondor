// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewaresengine

import (
	"github.com/gin-gonic/gin"
)

// MiddlewaresEngine handles middlewares registration
type MiddlewaresEngine struct {
	middlewares []gin.HandlerFunc
}

//Middleware a function defines a middleware
type Middleware func(c *gin.Context)

var mwEngine *MiddlewaresEngine

//NewMiddlewareEngine initiate a new middlware engine
func New() *MiddlewaresEngine {
	mwEngine = &MiddlewaresEngine{}
	return mwEngine
}

//ResolveMiddlewaresEngine resolve an already initated middleware engine
func Resolve() *MiddlewaresEngine {
	return mwEngine
}

//Attach attach a middleware globally to the app
func (m *MiddlewaresEngine) Attach(mw gin.HandlerFunc) *MiddlewaresEngine {
	m.middlewares = append(m.middlewares, mw)

	return mwEngine
}

//GetMiddlewares get all attached middlewares
func (m *MiddlewaresEngine) GetMiddlewares() []gin.HandlerFunc {
	return m.middlewares
}
