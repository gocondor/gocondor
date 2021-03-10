# Gincoat

## [Under Deveopment]

### What is Gincoat?
Gincoat is a golang framework based on Gin framework with a style similar to Laravel's style for developing modern apis. it has an `MVC` like routing style, middlewares, database, authentication with jwt and more.

### The architecture
the architecture is similar to MVC architecture, there is a `routes.go` file where you can defne all your routes and their `handlers`, a handler crosponds to controller action in `MVC`
the execusion jouerny looks like this:

`request -> routing -> middleware -> handler -> json response`

### the folder structure 
```bash
├── gincoat
│   ├── config/ 
│   ├── core/
│   ├── httpd/
│   │   ├── handlers/
│   │   ├── middlewares/
│   ├── routes.go
│   ├── integrations/
│   ├── logs/
│   ├── models/
│   ├── ssh/
├── .env
└── .gitignore
├── go.mod
├── LICENSE
├── README.md
```