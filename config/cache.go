// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import "github.com/gocondor/core"

// Retrieve the main config for the cache
func GetCacheConfig() core.CacheConfig {
	//#####################################
	//# Main configuration for cache  #####
	//#####################################

	return core.CacheConfig{
		// For enabling and disabling the cache
		// set to true to enable it, set to false to disable
		EnableCache: true,
	}
}
