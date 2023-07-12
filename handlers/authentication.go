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
		c.Response.SetStatusCode(http.StatusUnprocessableEntity).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
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
		c.Response.SetStatusCode(http.StatusUnprocessableEntity).WriteJson([]byte(v.GetErrorMessagesJson()))
		return
	}

	//hash the password
	passwordHashed, err := c.GetHashing().HashPassword(c.InterfaceToString(password))
	if err != nil {
		c.Response.SetStatusCode(http.StatusUnprocessableEntity).WriteJson([]byte(c.MapToJson(map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})))
		return
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

func Signin(c *core.Context) {
	email := c.GetRequestParam("email")
	password := c.GetRequestParam("password")

	v := c.GetValidator().Validate(map[string]interface{}{
		"email":    email,
		"password": password,
	}, map[string]interface{}{
		"email":    "required|email",
		"password": "required",
	})

	if v.Failed() {
		c.Response.SetStatusCode(http.StatusOK).WriteJson([]byte(v.GetErrorMessagesJson()))
		return
	}

	var user models.User
	res := c.GetGorm().Where("email = ?", c.InterfaceToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.LogError(res.Error.Error())
		c.Response.SetStatusCode(http.StatusInternalServerError).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal server error",
		})))
		return
	}

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.Response.SetStatusCode(http.StatusUnprocessableEntity).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "wrong email or password",
		})))
		return
	}

	ok, err := c.GetHashing().CheckPasswordHash(user.Password, c.InterfaceToString(password))
	if err != nil {
		c.LogError(err.Error())
		c.Response.SetStatusCode(http.StatusInternalServerError).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})))
		return
	}

	if !ok {
		c.Response.SetStatusCode(http.StatusUnprocessableEntity).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "wrong email or password",
		})))
		return
	}

	token, err := c.GetJWT().GenerateToken(map[string]interface{}{
		"id": user.ID,
	})

	if err != nil {
		c.LogError(err.Error())
		c.Response.SetStatusCode(http.StatusInternalServerError).WriteJson([]byte(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal server error",
		})))
		return
	}
	c.Response.WriteJson([]byte(c.MapToJson(map[string]string{
		"status": "success",
		"token":  token,
	})))
}

// TODO implement
func ResetPassword(c *core.Context) {

}
