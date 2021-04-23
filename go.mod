module github.com/gocondor/gocondor

replace (
	github.com/gocondor/gocondor/config => ./config
	github.com/gocondor/gocondor/http => ./http
	github.com/gocondor/gocondor/http/middlewares => ./http/middlewares
	github.com/gocondor/gocondor/http/handlers => ./http/handlers
	github.com/gocondor/gocondor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210422202316-a6c1295d69b4
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
