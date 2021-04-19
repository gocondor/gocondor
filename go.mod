module github.com/gocondor/gocondor

replace (
	github.com/gocondor/gocondor/config => ./config
	github.com/gocondor/gocondor/http => ./http
	github.com/gocondor/gocondor/http/middlewares => ./http/middlewares
	github.com/gocondor/gocondor/integrations => ./integrations
	github.com/gocondor/gocondor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/autotls v0.0.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210419205232-3722a757ec4b
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
