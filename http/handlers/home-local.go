// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HomeShow to show home page
func HomeDisplay(c *gin.Context) {
	// jwt := c.MustGet(core.JWT).(*jwtloader.JwtLoader)
	// token, err := jwt.CreateRefreshToken(map[string]string{"userId": "123"})

	// cache := c.MustGet(core.CACHE).(*cache.CacheEngine)

	// _, err := cache.Set("name", "harran")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// res, err := cache.Get("name")

	// fmt.Println(res)
	// jwtloader.New()
	// tokenCreator := jwtloader.Resolve()

	// payload := map[string]string{
	// 	"userId": "3333",
	// }
	// token, err := tokenCreator.CreateToken(payload)
	//refreshToken, err := tokenCreator.CreateRefreshToken(payload)

	//tokenString, err := tokenCreator.ExtractToken(c)
	// _, err = tokenCreator.ValidateToken(tokenString)
	// payload, err := tokenCreator.DecodeToken(tokenString)

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	fmt.Println("payload from home is")
	// fmt.Println(payload)
	//fmt.Println("the token is:::::::" + token)
	//fmt.Println("the refresh token is:::::::" + refreshToken)

	//--------------------------------------------------
	// db := c.MustGet(core.GORM).(*gorm.DB)
	// fmt.Println("+++++++++db ", db)
	// db.First(&models.Product{})

	c.JSON(200, gin.H{
		"token": "tokenStringffff",
	})
}

// HomeUpdate to update home page
func HomeEdit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update home page",
	})
}
