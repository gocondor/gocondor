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
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.7.1
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unrolled/secure v1.0.8
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)
