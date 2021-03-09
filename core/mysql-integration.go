// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Mysql returns a gin handler func with db set in gin.Context
func Mysql(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
