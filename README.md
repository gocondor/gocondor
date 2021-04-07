# GoCondor

![Build Status](https://github.com/gocondor/gocondor/actions/workflows/build-master.yml/badge.svg)
![Test Status](https://github.com/gocondor/gocondor/actions/workflows/test-master.yml/badge.svg)

## [Under Development]

## What is GoCondor?

GoCondor is a golang web framework with an `MVC` like architecture, it's based on [Gin framework](https://github.com/gin-gonic/gin), it provides you with an easy-to-use directory structure with a development experience similar to Laravel, made for developing modern APIs and microservices.
Full docmentation is available at [GoCondor Docs](https://gocondor.github.io/docs/).

## Features 
- Router
- Middlewares
- JWT tokens
- ORM (GORM)
- Cache (Redis)
- TLS
- Context Package Integrator
- Live-Reloading for development
- Features Control

## Create a new project 
To create a new GoCondor project you need to install the `GoCondor cli` first

#### Install Condor cli 
To install the `GoCondor cli` run the following command
```bash
go get github.com/gocondor/installer/gocondor
```

#### Create a project using GoCondor cli:
To create a new project run the following command:
```bash
Gocondor new my-project github.com/my-organization/my-project
```

## Getting started
First create a project by following the steps above.
Now Let's create a route to handle our first request
let's start by defining the handler function for the request, create a file with the name `example.go` in `httpd/handlers` folder with the following content:
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
Next start the app by running the following code from inside the project:
```bash
gocondor run:dev
```
Finally, open up your browser and navigate to `localhost:8000`


## Architecture
the architecture is similar to `MVC` architecture, there is a `routes.go` file where you can define all your routes and their `handlers`, the handler is simply a method that gets executed when the request is received, you can think of it like a controllers action in `MVC`

### The request journey looks like below:
`request -> routing -> middleware -> handler -> middleware -> json response`

## Folder structure 
```bash
├── gocondor
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
