// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"fmt"

	"github.com/gocondor/core"
)

// Another example middleware
var AnotherExampleMiddleware core.Middleware = func(c *core.Context) {
	fmt.Println("another example middleware!")
	c.Next()
}
