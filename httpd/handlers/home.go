// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/harranali/gincoat/core"
	"github.com/harranali/gincoat/models"
	"gorm.io/gorm"
)

// HomeShow to show home page
func HomeShow(c *gin.Context) {
	db := c.MustGet(core.DB).(*gorm.DB)

	db.First(&models.Product{})

	c.JSON(200, gin.H{
		"message": "show home page!",
	})
}

// HomeUpdate to update home page
func HomeUpdate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update home page",
	})
}
