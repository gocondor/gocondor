// // Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// // Use of this source code is governed by MIT-style
// // license that can be found in the LICENSE file.

package handlers

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/events"
	"github.com/gocondor/gocondor/models"
	"github.com/google/uuid"
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
		c.GetLogger().Error(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal error",
		}))
	}
	if res.Error == nil {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "email already exists in the database",
		}))
	}

	// validation data
	data := map[string]interface{}{
		"name":     name,
		"email":    email,
		"password": password,
	}
	// validation rules
	rules := map[string]interface{}{
		"name":     "required|alphaNumeric",
		"email":    "required|email",
		"password": "required|length:6,10",
	}
	// validate
	v := c.GetValidator().Validate(data, rules)
	if v.Failed() {
		c.GetLogger().Error(v.GetErrorMessagesJson())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(v.GetErrorMessagesJson())
	}

	//hash the password
	passwordHashed, err := c.GetHashing().HashPassword(c.CastToString(password))
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]interface{}{
			"message": err.Error(),
		}))
	}
	// store the record in db
	user = models.User{
		Name:     c.CastToString(name),
		Email:    c.CastToString(email),
		Password: passwordHashed,
	}
	res = c.GetGorm().Create(&user)
	if res.Error != nil {
		c.GetLogger().Error(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": res.Error.Error(),
		}))
	}

	token, err := c.GetJWT().GenerateToken(map[string]interface{}{
		"userID": user.ID,
	})

	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}
	// cache the token
	userAgent := c.Request.HttpRequest.UserAgent()
	cacheKey := fmt.Sprintf("userid:_%v_useragent:_%v_jwt_token", user.ID, userAgent)
	hashedCacheKey := c.CastToString(fmt.Sprintf("%x", md5.Sum([]byte(cacheKey))))
	err = c.GetCache().Set(hashedCacheKey, token)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]interface{}{
			"message": "internal server error",
		}))
	}

	// fire user registered event
	err = c.GetEventsManager().Fire(&core.Event{Name: events.USER_REGISTERED, Payload: map[string]interface{}{
		"user": user,
	}})
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"token": token,
	}))
}

func Signin(c *core.Context) *core.Response {
	email := c.GetRequestParam("email")
	password := c.GetRequestParam("password")

	data := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	rules := map[string]interface{}{
		"email":    "required|email",
		"password": "required",
	}
	v := c.GetValidator().Validate(data, rules)

	if v.Failed() {
		c.GetLogger().Error(v.GetErrorMessagesJson())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(v.GetErrorMessagesJson())
	}

	var user models.User
	res := c.GetGorm().Where("email = ?", c.CastToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.GetLogger().Error(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid email or password",
		}))
	}

	ok, err := c.GetHashing().CheckPasswordHash(user.Password, c.CastToString(password))
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": err.Error(),
		}))
	}

	if !ok {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid email or password",
		}))
	}

	token, err := c.GetJWT().GenerateToken(map[string]interface{}{
		"userID": user.ID,
	})

	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}
	// cache the token
	userAgent := c.Request.HttpRequest.UserAgent()
	cacheKey := fmt.Sprintf("userid:_%v_useragent:_%v_jwt_token", user.ID, userAgent)
	hashedCacheKey := c.CastToString(fmt.Sprintf("%x", md5.Sum([]byte(cacheKey))))
	err = c.GetCache().Set(hashedCacheKey, token)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]interface{}{
			"message": "internal server error",
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"token": token,
	}))
}

func ResetPasswordRequest(c *core.Context) *core.Response {
	email := c.GetRequestParam("email")

	// validation data
	data := map[string]interface{}{
		"email": email,
	}
	// validation rules
	rules := map[string]interface{}{
		"email": "required|email",
	}
	// validate
	v := c.GetValidator().Validate(data, rules)
	if v.Failed() {
		c.GetLogger().Error(v.GetErrorMessagesJson())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(v.GetErrorMessagesJson())
	}

	// check email in the database
	var user models.User
	res := c.GetGorm().Where("email = ?", c.CastToString(email)).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.GetLogger().Error(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "email not found in our database",
		}))
	}

	// generate the link
	expiresAt := time.Now().Add(time.Hour * 3).Unix()
	linkCodeData := map[string]string{
		"userID":    c.CastToString(user.ID),
		"expiresAt": c.CastToString(expiresAt),
	}
	code := uuid.NewString()
	c.GetCache().SetWithExpiration(code, c.MapToJson(linkCodeData), time.Hour*3)
	err := c.GetEventsManager().Fire(&core.Event{Name: events.USER_PASSWORD_RESET_REQUESTED, Payload: map[string]interface{}{
		"user": user,
		"code": code,
	}})
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"message": "reset password email sent successfully",
	}))
}

func SetNewPassword(c *core.Context) *core.Response {
	urlCode := c.CastToString(c.GetPathParam("code"))
	linkCodeDataStr, err := c.GetCache().Get(urlCode)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}
	var linkCode map[string]string
	json.Unmarshal([]byte(linkCodeDataStr), &linkCode)
	expiresAtUnix, err := strconv.ParseInt(linkCode["expiresAt"], 10, 64)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}
	expiresAt := time.Unix(expiresAtUnix, 0)
	if time.Now().After(expiresAt) {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}
	userID, err := strconv.ParseUint(linkCode["userID"], 10, 64)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}

	oldPassword := c.CastToString(c.GetRequestParam("old_password"))
	newPassword := c.CastToString(c.GetRequestParam("new_password"))
	newPasswordConfirm := c.CastToString(c.GetRequestParam("new_password_confirm"))

	// validation data
	data := map[string]interface{}{
		"old_password":         oldPassword,
		"new_password":         newPassword,
		"new_password_confirm": newPasswordConfirm,
	}
	// validation rules
	rules := map[string]interface{}{
		"old_password":         "required|length:6,10",
		"new_password":         "required|length:6,10",
		"new_password_confirm": "required|length:6,10",
	}
	// validate
	v := c.GetValidator().Validate(data, rules)
	if v.Failed() {
		c.GetLogger().Error(v.GetErrorMessagesJson())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(v.GetErrorMessagesJson())
	}

	var user models.User
	res := c.GetGorm().Where("id = ?", userID).First(&user)
	if res.Error != nil {
		c.GetLogger().Error(res.Error.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}

	SamePassword, err := c.GetHashing().CheckPasswordHash(user.Password, oldPassword)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid link",
		}))
	}

	if !SamePassword {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "the old password is incorrect",
		}))
	}

	if newPassword != newPasswordConfirm {
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "new password does not match new password confirmation",
		}))
	}

	hashedNewPassword, err := c.GetHashing().HashPassword(newPassword)
	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusUnprocessableEntity).Json(c.MapToJson(map[string]string{
			"message": "invalid password",
		}))
	}

	user.Password = hashedNewPassword
	c.GetGorm().Save(&user)

	err = c.GetEventsManager().Fire(&core.Event{Name: events.PASSWORD_CHANGED, Payload: map[string]interface{}{
		"user": user,
	}})

	if err != nil {
		c.GetLogger().Error(err.Error())
		return c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]string{
			"message": "internal server error",
		}))
	}

	return c.Response.Json(c.MapToJson(map[string]string{
		"message": "password changed successfully",
	}))
}
