// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gincoat/gincoat/core/cache"
)

// Cache returns a gin handler func with cahce variable set in gin context
func Cache(cache *cache.CacheEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cache", cache)
		c.Next()
	}
}
