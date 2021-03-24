// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/gincoat/gincoat/core"
)

// Features to control Gincoat core features
var Features *core.Features = &core.Features{
	Database: false, // Database is flag enable or disable the database
	Cache:    false, // Cache is a flag to enable or disable the cache
	GRPC:     false, // GRPC is a flag to enable or disable gRPC
}
