// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import "github.com/gocondor/core"

// Retrieve the main config for controlling the .env file
func GetEnvFileConfig() core.EnvFileConfig {
	//#########################################################
	//# Main configuration for controlling the .env file  #####
	//#########################################################

	return core.EnvFileConfig{
		// Set to true to read the environment variables from the .env file and then
		// inject them into the os environment, please keep in mind this will override any
		// variables previously set in the os envrionment
		// set to false to ignore the .env file
		UseDotEnvFile: true,
	}
}
