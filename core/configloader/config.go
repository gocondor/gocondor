// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package configloader

// ConfigLoader to load configuration
type ConfigLoader struct{}

var configLoader *ConfigLoader

// New initiate configuration loader variable
func New() *ConfigLoader {
	return &ConfigLoader{}
}

// Resolve gets the initiated configuration variable
func Resolve() *ConfigLoader {
	return configLoader
}
