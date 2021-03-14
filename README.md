# Gincoat

## [Under Deveopment]

### What's is Gincoat?
Gincoat is a golang framework based on Gin framework with a style similar to Laravel's style, made for developing modern apis. it has an `MVC` like architecture.

### Features 
- Router
- Middlewares
- JWT tokens
- ORM
- Redis cache
- TLS 
- Package Integrator

### Architecture
the architecture is similar to MVC architecture, there is a `routes.go` file where you can define all your routes and their `handlers`, a handler crosponds to controller action in `MVC`
the request jouerny looks like this:

`request -> routing -> middleware -> handler -> json response`

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
