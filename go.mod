module github.com/gocondor/gocondor

replace (
	github.com/gocondor/core => C:/Users/harran/go/src/github.com/gocondor/core
	github.com/gocondor/gocondor/config => ./config
	github.com/gocondor/gocondor/httpd => ./httpd
	github.com/gocondor/gocondor/httpd/middlewares => ./httpd/middlewares
	github.com/gocondor/gocondor/integrations => ./integrations
	github.com/gocondor/gocondor/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/autotls v0.0.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210407100326-7c4af7b4c5a0
	github.com/joho/godotenv v1.3.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	gorm.io/gorm v1.21.6
)
