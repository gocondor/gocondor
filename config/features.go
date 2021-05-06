// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/gocondor/core"
)

// Features helps you decide what features of the framework to use
var Features *core.Features = &core.Features{
	Database:       true,
	Cache:          false,
	GRPC:           false,
	Sessions:       true,
	Authentication: true,
}
