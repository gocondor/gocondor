# Gincoat

![Build Status](https://github.com/gincoat/gincoat/actions/workflows/build-master.yml/badge.svg)
![Test Status](https://github.com/gincoat/gincoat/actions/workflows/test-master.yml/badge.svg)

## [Under Development]

### What is Gincoat?
Gincoat is a golang framework based on Gin with a development experience similar to Laravel, made for developing modern apis and microservices. it has an `MVC` like architecture.

### Features 
- Development Server with live-reloading
- Router
- Middlewares
- JWT tokens
- ORM (GORM)
- Redis cache
- TLS
- gRPC
- Context Package Integrator
- Features Control

### Create a new project 
To create a new Gincoat project you need to install the Gincoat installer cli tool first

#### Installing Gincoat cli 
To install the cli run the following command
```bash
go get github.com/gincoat/installer/gincoat
```

#### Create a new project:
To create a new project run the following command:
```bash
gincoat new my-project github.com/my-organization/my-project
```

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
│   ├── .gitignore ------------> gitignore file
│   ├── go.mod ----------------> go modules the project depends on
│   ├── LICENSE ---------------> license
│   ├── main.go ---------------> ssl certificates
│   ├── README.md -------------> readme file
```
