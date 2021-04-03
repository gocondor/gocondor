// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"github.com/gin-gonic/gin"
)

// PKGIntegratorExample to integrate a package
var PKGIntegratorExample gin.HandlerFunc = func(c *gin.Context) {
	//initiate your package variable here

	//set the package varialbe in gin context using c.Set(key, pkg-variable)
	c.Set("key", "package-var")

	//continute execution
	c.Next()
}
