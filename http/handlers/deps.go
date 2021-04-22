package handlers

import (
	"github.com/gocondor/core/cache"
	"github.com/gocondor/core/database"
	"github.com/gocondor/core/jwt"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB           // DB database variable (gorm)
	Cache *cache.CacheEngine // Cache variable
	JWT   *jwt.JWTUtil       // JWT represents managment variable
)

// InitiateHandlersDependencies to initiate the any dependency of the handlers
func InitiateHandlersDependencies() {
	DB = database.Resolve()
	Cache = cache.Resolve()
	JWT = jwt.Resolve()
}
