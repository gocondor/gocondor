// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MiddlewareExample is an example of a middleware gets executed before the request handler
var MiddlewareExample gin.HandlerFunc = func(c *gin.Context) {
	fmt.Println("I'm an example middleware!")
	// Pass on to the next-in-chain
	c.Next()
}
