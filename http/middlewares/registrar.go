// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gocondor/core/middlewaresengine"
)

// RegisterMiddlewares helps you attach middlwares globally
func RegisterMiddlewares() {
	mwEngine := middlewaresengine.Resolve()

	// Register your middlewares here
	mwEngine.Attach(MiddlewareExample)
}
