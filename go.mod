module github.com/gocondor/gocondor

replace (
	github.com/gocondor/gocondor/config => ./config
	github.com/gocondor/gocondor/handlers => ./handlers
	github.com/gocondor/gocondor/middlewares => ./middlewares
	github.com/gocondor/gocondor/models => ./models
)

go 1.20

require (
	github.com/gocondor/core v1.5.0
	github.com/joho/godotenv v1.5.1
	github.com/julienschmidt/httprouter v1.3.0
	gorm.io/gorm v1.25.2
)

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/redis/go-redis/v9 v9.0.5 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/driver/postgres v1.5.2 // indirect
	gorm.io/driver/sqlite v1.5.2 // indirect
)
