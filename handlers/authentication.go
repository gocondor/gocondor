// // Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// // Use of this source code is governed by MIT-style
// // license that can be found in the LICENSE file.

package handlers

import (
	"errors"
	"net/http"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/events"
	"github.com/gocondor/gocondor/models"
	"gorm.io/gorm"
)

func Signup(c *core.Context) *core.Response {
	name := c.GetRequestParam("name")
	email := c.GetRequestParam("email")
	password := c.GetRequestParam("password")
	// check if email exists
	var user models.User
	res := c.GetGorm().Where("email = ?", c.CastToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.LogError(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal error",
		}))
	}
	if (user != models.User{}) {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "email already exist in the database",
		}))
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
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(v.GetErrorMessagesJson())
	}

	//hash the password
	passwordHashed, err := c.GetHashing().HashPassword(c.CastToString(password))
	if err != nil {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		}))
	}
	user = models.User{
		Name:     c.CastToString(name),
		Email:    c.CastToString(email),
		Password: passwordHashed,
	}
	res = c.GetGorm().Create(&user)
	if res.Error != nil {
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": res.Error.Error(),
		}))
	}

	err = c.GetEventsManager().Fire(&core.Event{Name: events.USER_REGISTERED, Payload: map[string]interface{}{
		"userStruct": user,
	}})
	if err != nil {
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal error",
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"status":  "success",
		"message": "user created successfully",
	}))
}

func Signin(c *core.Context) *core.Response {
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
		return c.Response.SetStatusCode(http.StatusOK).Json(v.GetErrorMessagesJson())
	}

	var user models.User
	res := c.GetGorm().Where("email = ?", c.CastToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.LogError(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal server error",
		}))
	}

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "wrong email or password",
		}))
	}

	ok, err := c.GetHashing().CheckPasswordHash(user.Password, c.CastToString(password))
	if err != nil {
		c.LogError(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": err.Error(),
		}))
	}

	if !ok {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "wrong email or password",
		}))
	}

	token, err := c.GetJWT().GenerateToken(map[string]interface{}{
		"id": user.ID,
	})

	if err != nil {
		c.LogError(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"status":  "error",
			"message": "internal server error",
		}))
	}
	return c.Response.Json(c.MapToJson(map[string]string{
		"status": "success",
		"token":  token,
	}))
}

// TODO implement
func ResetPassword(c *core.Context) *core.Response {
	return nil
}
