// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authpkg "github.com/gocondor/core/auth"
)

// Auth checks if the request is authenticated
var Auth gin.HandlerFunc = func(c *gin.Context) {
	auth := authpkg.Resolve()
	ok, err := auth.Check(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong while performing auth check",
		})
		return
	}

	if !ok {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
		return
	}

	// Pass on to the next-in-chain
	c.Next()
}
