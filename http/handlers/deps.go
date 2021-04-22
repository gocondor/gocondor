package handlers

import (
	"github.com/gocondor/core/cache"
	"github.com/gocondor/core/database"
	"github.com/gocondor/core/jwt"
	"gorm.io/gorm"
)

var (
	// DB for database manipulation
	DB *gorm.DB
	// Cache for cache manipulation
	Cache *cache.CacheEngine
	// JWT used for jwt tokens creation and validation
	JWT *jwt.JWTUtil
)

// InitiateHandlersDependencies to initiate the any dependency of the handlers
func InitiateHandlersDependencies() {
	DB = database.Resolve()
	Cache = cache.Resolve()
	JWT = jwt.Resolve()
}
