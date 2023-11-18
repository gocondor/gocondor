package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gocondor/core"
	"github.com/gocondor/gocondor/models"
	"github.com/gocondor/gocondor/utils"
	"gorm.io/gorm"
)

var AuthCheck core.Middleware = func(c *core.Context) {
	tokenRaw := c.GetHeader("Authorization")
	token := strings.TrimSpace(strings.Replace(tokenRaw, "Bearer", "", 1))
	if token == "" {
		c.Response.SetStatusCode(http.StatusUnauthorized).Json(c.MapToJson(map[string]interface{}{
			"message": "unauthorized",
		})).ForceSendResponse()
		return
	}
	payload, err := c.GetJWT().DecodeToken(token)
	if err != nil {
		c.Response.SetStatusCode(http.StatusUnauthorized).Json(c.MapToJson(map[string]interface{}{
			"message": "unauthorized",
		})).ForceSendResponse()
		return
	}
	userAgent := c.GetUserAgent()
	hashedCacheKey := utils.CreateAuthTokenHashedCacheKey(uint(c.CastToInt(payload["userID"])), userAgent)

	cachedToken, err := c.GetCache().Get(hashedCacheKey)
	if err != nil {
		// user signed out
		c.Response.SetStatusCode(http.StatusUnauthorized).Json(c.MapToJson(map[string]interface{}{
			"message": "unauthorized",
		})).ForceSendResponse()
		return
	}
	if cachedToken != token {
		// using old token replaced with new one after recent signin
		c.Response.SetStatusCode(http.StatusUnauthorized).Json(c.MapToJson(map[string]interface{}{
			"message": "unauthorized",
		})).ForceSendResponse()
		return
	}

	var user models.User
	res := c.GetGorm().Where("id = ?", payload["userID"]).First(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// error with the database
		c.GetLogger().Error(res.Error.Error())
		c.Response.SetStatusCode(http.StatusInternalServerError).Json(c.MapToJson(map[string]interface{}{
			"message": "internal error",
		})).ForceSendResponse()
		return
	}

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// user record is not found (deleted)
		c.Response.SetStatusCode(http.StatusUnauthorized).Json(c.MapToJson(map[string]interface{}{
			"message": "unauthorized",
		})).ForceSendResponse()
		return
	}

	c.Next()
}
