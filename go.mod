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
	github.com/gincoat/core v0.0.0-20210329051702-7e362ac70582
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
