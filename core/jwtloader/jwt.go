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

var defaultLifeSpan time.Duration = 15 //15 minutes

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
		return "", errors.New("missing jwt secret")
	}
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
