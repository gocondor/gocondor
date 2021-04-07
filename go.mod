module github.com/gocondor/condor

replace (
	github.com/gocondor/condor/config => ./config
	github.com/gocondor/condor/httpd => ./httpd
	github.com/gocondor/condor/httpd/middlewares => ./httpd/middlewares
	github.com/gocondor/condor/integrations => ./integrations
	github.com/gocondor/condor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210407100326-7c4af7b4c5a0
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
