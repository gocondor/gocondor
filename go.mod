module github.com/gincoat/gincoat

replace (
	github.com/gincoat/gincoat/config => ./config
	github.com/gincoat/gincoat/httpd => ./httpd
	github.com/gincoat/gincoat/httpd/middlewares => ./httpd/middlewares
	github.com/gincoat/gincoat/integrations => ./integrations
	github.com/gincoat/gincoat/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gincoat/core v0.0.0-20210403060125-6a9147e1a71e
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
