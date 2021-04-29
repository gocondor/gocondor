module github.com/gocondor/gocondor

replace (
	github.com/gocondor/gocondor/config => ./config
	github.com/gocondor/gocondor/http => ./http
	github.com/gocondor/gocondor/http/handlers => ./http/handlers
	github.com/gocondor/gocondor/http/middlewares => ./http/middlewares
	github.com/gocondor/gocondor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/gocondor/core v0.0.0-20210428170922-d3b9fcce6905
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
