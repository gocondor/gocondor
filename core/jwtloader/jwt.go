// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package jwtloader

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtLoader pakcage struct
type JwtLoader struct{}

var defaultLifeSpan time.Duration = 15             //15 minutes
var defaultRefreshTokenLifeSpan time.Duration = 24 //24 hours

var jwtLoader *JwtLoader

// New initiates Jwt struct
func New() *JwtLoader {
	jwtLoader = &JwtLoader{}
	return jwtLoader
}

//Resolve returns initiated jwt token
func Resolve() *JwtLoader {
	return jwtLoader
}

// CreateToken generates new jwt token with the given payload
func (j *JwtLoader) CreateToken(payload map[string]string) (string, error) {

	claims := jwt.MapClaims{}

	var duration time.Duration
	durationStr := os.Getenv("JWT_LIFESPAN_MINUTES")
	if durationStr == "" {
		duration = defaultLifeSpan
	} else {
		d, _ := strconv.ParseInt(durationStr, 10, 64)
		duration = time.Duration(d)
	}

	for key, val := range payload {
		claims[key] = val
	}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * duration).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("missing jwt token secret")
	}
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateRefreshToken generates new jwt refresh token with the given payload
func (j *JwtLoader) CreateRefreshToken(payload map[string]string) (string, error) {
	claims := jwt.MapClaims{}

	var duration time.Duration
	durationStr := os.Getenv("JWT_REFRESH_TOKEN_LIFESPAN_HOURS")
	if durationStr == "" {
		duration = defaultRefreshTokenLifeSpan
	} else {
		d, _ := strconv.ParseInt(durationStr, 10, 64)
		duration = time.Duration(d)
	}

	for key, val := range payload {
		claims[key] = val
	}
	claims["refresh"] = true
	claims["exp"] = time.Now().Add(time.Hour * duration).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("missing jwt token refresh secret")
	}
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

//ExtractToken extracts the token from the request header
func (j *JwtLoader) ExtractToken(c *gin.Context) (token string, err error) {
	sentTokenSlice := c.Request.Header["Authorization"]
	if len(sentTokenSlice) == 0 {
		return "", errors.New("Missing authorization token")
	}
	sentTokenSlice = strings.Split(sentTokenSlice[0], " ")
	if len(sentTokenSlice) != 2 {
		return "", errors.New("Something wrong with the token")
	}

	return sentTokenSlice[1], nil
}

// DecodeToken decodes a given token and returns the payload
func (j *JwtLoader) DecodeToken(tokenString string) (payload map[string]interface{}, err error) {
	// validate the token
	_, err = j.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	//extract claims
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	claims := token.Claims.(jwt.MapClaims)
	delete(claims, "authorized")
	delete(claims, "exp")

	return claims, nil
}

// ValidateToken makes sure the given token is valid
func (j *JwtLoader) ValidateToken(tokenString string) (bool, error) {
	// parse the token string
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %s", token.Method.Alg())
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

// RefreshToken generates a new token based on the refresh token
// TODO: implement
func RefreshToken(token string, refreshToken string) (newToken string, err error) {
	return
}
