// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
)

// HomeShow to show home page
func HomeShow(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello from home show!",
	})
}

// HomeUpdate to update home page
func HomeUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello from home update",
	})
}
