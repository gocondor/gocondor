package middlewares

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

// InitiateMiddlewaresDependencies to initiate the any dependency of the handlers
func InitiateMiddlewaresDependencies() {
	DB = database.Resolve()
	Cache = cache.Resolve()
	JWT = jwt.Resolve()
}
