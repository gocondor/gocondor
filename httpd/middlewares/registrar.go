// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/harranali/gincoat/core/middlewaresengine"
)

// RegisterMiddlewares middlwares
func RegisterMiddlewares() {
	mwEngine := middlewaresengine.Resolve()

	//Register your middlewares here
	mwEngine.Attach(MiddlewareExample) // Attach the middleware `MiddlewareExample` globally
}
