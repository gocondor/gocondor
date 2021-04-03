// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/gincoat/core"
)

// Features decide what features of the framework to use
var Features *core.Features = &core.Features{
	Database: false,
	Cache:    false,
	GRPC:     false,
}
