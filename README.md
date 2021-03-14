# Gincoat

## [Under Deveopment]

### What's is Gincoat?
Gincoat is a golang framework based on Gin with a development experience similar to Laravel, made for developing modern apis and microservices. it has an `MVC` like architecture.

### Features 
- Router
- Middlewares
- JWT tokens
- Multiple database support
- ORM (GORM)
- MongoDB integration
- Redis cache
- TLS
- gRPC
- Package Integrator
- CLI with generators for files and code snippets
- Development Server with hot reloading
- Features Switch

### Architecture
the architecture is similar to `MVC` architecture, there is a `routes.go` file where you can define all your routes and their `handlers`, a handler crosponds to a controller action in `MVC`

The request jouerny looks like this:

`request -> routing -> middleware -> handler -> middleware -> json response`

### Folder structure 
```bash
├── gincoat
│   ├── config/ ---------------> app configuration
│   ├── core/ -----------------> contains the core packages of the framework
│   ├── httpd/-----------------> all http related functionalities go here 
│   │   ├── handlers/ ---------> contains the requests handlers
│   │   ├── middlewares/ ------> here you can puth your middlewares
│   ├── routes.go -------------> app routes are defined here
│   ├── integrations/ ---------> contains the integrations of third party packages into gin context
│   ├── logs/ -----------------> log files
│   ├── models/ ---------------> models
│   ├── ssl/ ------------------> ssl certificates
│   ├── .env ------------------> environment variables 
│   ├── .gitignore ----------------> gitignore file
│   ├── go.mod --------------------> go modules the project depends on
│   ├── LICENSE -------------------> license
│   ├── main.go -------------------> ssl certificates
│   ├── README.md -----------------> readme file
```
