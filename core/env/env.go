// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package env

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var env map[string]string
var err error

// Load reads the environment vars
func Load() {
	// load env vars
	env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//set to os
	for key, val := range env {
		os.Setenv(strings.TrimSpace(key), strings.TrimSpace(val))
	}
}

//Get get env var
func Get(key string) string {
	return strings.TrimSpace(env[key])
}
