// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"errors"
	"net/http"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/models"
	"gorm.io/gorm"
)

func Signup(c *core.Context) {
	name := c.GetRequestParam("name")
	email := c.GetRequestParam("email")
	password := c.GetRequestParam("password")
	// check if email exists
	var user models.User
	res := c.GetGorm().Where("email = ?", c.InterfaceToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.LogError(res.Error.Error())
		c.Response.SetStatusCode(http.StatusInternalServerError).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal error",
		})))
		return
	}
	if (user != models.User{}) {
		c.Response.SetStatusCode(http.StatusOK).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "success",
			"message": "email already exist in the database",
		})))
		return
	}

	// validate the input
	v := c.GetValidator().Validate(map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": password,
	}, map[string]interface{}{
		"name":     "required|alphaNumeric",
		"email":    "required|email",
		"password": "required|length:6,10",
	})

	if v.Failed() {
		c.Response.WriteJson([]byte(v.GetErrorMessagesJson()))
		return
	}

	//hash the password
	passwordHashed, err := c.GetHashing().HashPassword(c.InterfaceToString(password))
	if err != nil {
		c.Response.WriteJson([]byte(c.MapToJson(map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})))
	}

	c.GetGorm().Create(&models.User{
		Name:     c.InterfaceToString(name),
		Email:    c.InterfaceToString(email),
		Password: passwordHashed,
	})

	c.Response.WriteJson([]byte(c.MapToJson(map[string]string{
		"status":  "success",
		"message": "user created successfully",
	})))
}

// TODO implement
func Signin(c *core.Context) {

}

// TODO implement
func ResetPassword(c *core.Context) {

}
