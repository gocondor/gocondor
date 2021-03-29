# Gincoat

![Build Status](https://github.com/gincoat/gincoat/actions/workflows/build-master.yml/badge.svg)
![Test Status](https://github.com/gincoat/gincoat/actions/workflows/test-master.yml/badge.svg)

## [Under Development]

## What is Gincoat?

Gincoat is a golang web framework with an `MVC` like architecture, it's based on [Gin framework](https://github.com/gin-gonic/gin), it has a development experience similar to Laravel, made for developing modern APIs and microservices.

## Features 
- Live-Reloading for development
- Router
- Middlewares
- JWT tokens
- ORM (GORM)
- Cache (Redis)
- TLS
- Context Package Integrator
- Features Control

## Create a new project 
To create a new Gincoat project you need to install the `Gincoat cli` first

#### Install Gincoat cli 
To install the `Gincoat cli` run the following command
```bash
go get github.com/gincoat/installer/gincoat
```

#### Create a new project:
To create a new project run the following command:
```bash
gincoat new my-project github.com/my-organization/my-project
```

## Getting started
First create a project by following the steps above.
Now Let's create a route to handle our first request
let's start by defining the handler function for the request, to do that first create a file with the name `example.go` in `httpd/handlers` folder and then add to it the following code:
```go
package handlers

import (
"github.com/gin-gonic/gin"
)

func ExampleShow(c *gin.Context) {
    message := "Hello from example handler!"
c.JSON(200, gin.H{
"message": message,
})
}
```
Next lets define the route, to do that open up the file `httpd/routes.go`, then inside the function `RegisterRoutes()` add to this line `router.Get("/", handlers.ExampleShow)` and make sure it looks like below:
```go
func RegisterRoutes() {
router := routing.Resolve()
// Define your routes here
router.Get("/", handlers.ExampleShow)
}

```
next start the app by running the following code from inside the project:
```bash
gincoat run:dev
```
Finally, open up your browser and navigate to `localhost:8000`


## Architecture
the architecture is similar to `MVC` architecture, there is a `routes.go` file where you can define all your routes and their `handlers`, the handler is simply a method to handle the received request, you can think of it like a controller action in `MVC`

### The request journey looks like below:
`request -> routing -> middleware -> handler -> middleware -> json response`

## Folder structure 
```bash
├── gincoat
│   ├── config/ ---------------> control what features to turn on
│   ├── httpd/-----------------> http related code
│   │   ├── handlers/ --------------> contains your http requests handlers
│   │   ├── middlewares/ -----------> middlewares are defined here
│   ├── routes.go -------------> your routes are defined here
│   ├── integrations/ ---------> contains the integrations of third party packages into gin context
│   ├── logs/ -----------------> log files
│   ├── models/ ---------------> database models
│   ├── ssl/ ------------------> ssl certificates
│   ├── .env ------------------> environment variables 
│   ├── .gitignore ------------> gitignore file
│   ├── go.mod ----------------> go modules the project depends on
│   ├── LICENSE ---------------> license
│   ├── main.go ---------------> ssl certificates
│   ├── README.md -------------> readme file
```
