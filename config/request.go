// Copyright 2023 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import "github.com/gocondor/core"

// Retrieve the main config for the HTTP request
func GetRequestConfig() core.RequestConfig {
	//#####################################
	//# Main configuration for gorm   #####
	//#####################################

	return core.RequestConfig{
		// Set the max file upload size
		MaxUploadFileSize: 20000000, // ~20MB
	}
}
