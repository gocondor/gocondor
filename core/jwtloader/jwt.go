// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package jwtloader

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/harranali/gincoat/core/env"
)

// JwtLoader pakcage struct
type JwtLoader struct{}

var defaultLifeSpan time.Duration = 15             //15 minutes
var defaultRefreshTokenLifeSpan time.Duration = 24 //24 hours

var jwtLoader *JwtLoader

// New initiates Jwt struct
func New() {
	jwtLoader = &JwtLoader{}
}

//Resolve returns initiated jwt token
func Resolve() *JwtLoader {
	return jwtLoader
}

// CreateToken generates new jwt token with the given payload
func (jl *JwtLoader) CreateToken(payload map[string]string) (string, error) {

	claims := jwt.MapClaims{}

	var duration time.Duration
	durationStr := env.Get("JWT_LIFESPAN_MINUTES")
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
	secret := env.Get("JWT_SECRET")
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
func (jl *JwtLoader) CreateRefreshToken(payload map[string]string) (string, error) {
	claims := jwt.MapClaims{}

	var duration time.Duration
	durationStr := env.Get("JWT_REFRESH_TOKEN_LIFESPAN_HOURS")
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
	secret := env.Get("JWT_REFRESH_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("missing jwt token refresh secret")
	}
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

// DecodeToken decodes a given token and returns the payload
// TODO: implement
func DecodeToken(token string) (payload map[string]string, err error) {
	return
}

// ValidateToken makes sure the given token is valid
// TODO: implement
func ValidateToken(token string) (isValid bool, err error) {
	return
}

// RefreshToken generates a new token based on the refresh token
// TODO: implement
func RefreshToken(token string, refreshToken string) (newToken string, err error) {
	return
}
