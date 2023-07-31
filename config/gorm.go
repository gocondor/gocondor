// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import "github.com/gocondor/core"

// Retrieve the main config for the GORM
func GetGormConfig() core.GormConfig {
	//#####################################
	//# Main configuration for GORM   #####
	//#####################################

	return core.GormConfig{
		// For enabling and disabling the GORM
		// set to true to enable it, set to false to disable
		EnableGorm: false,
	}
}
