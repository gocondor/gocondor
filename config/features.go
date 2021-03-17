package config

import (
	"github.com/gincoat/gincoat/core"
)

// Features to control Gincoat core features
var Features *core.Features = &core.Features{
	Database: false, // Database is flag enable or disable the database
	Cache:    true,  // Cache is a flag to enable or disable the cache
	GRPC:     true,  // GRPC is a flag to enable or disable gRPC
}
