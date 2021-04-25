![gocondor logo](https://github.com/gocondor/gocondor.github.io/raw/master/img/logo.png)

# GoCondor

![Build Status](https://github.com/gocondor/gocondor/actions/workflows/build-master.yml/badge.svg)
![Test Status](https://github.com/gocondor/gocondor/actions/workflows/test-master.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/gocondor/gocondor?status.svg)](https://godoc.org/github.com/gocondor/gocondor)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocondor/gocondor)](https://goreportcard.com/report/github.com/gocondor/gocondor)

## [Under Development]

## What is GoCondor?

GoCondor is a golang web framework with an `MVC` like architecture, it's based on [Gin framework](https://github.com/gin-gonic/gin), it features a development experience similar to Laravel, made for developing APIs and microservices.
Full documentation is available at [GoCondor Docs](https://gocondor.github.io/docs/).

## Features
- Router
- Middlewares
- JWT tokens
- ORM (GORM)
- Cache (Redis)
- TLS
- Live-Reloading for development
- Features Control

## Create a new project 
To create a new GoCondor project you need to install the `GoCondor cli` first

#### Install GoCondor cli 
To install the `GoCondor cli` run the following command
```bash
go get github.com/gocondor/installer/gocondor
```

#### Create a project using GoCondor cli:
To create a new project run the following command:
```bash
gocondor new my-project github.com/my-organization/my-project
```

## Getting started
Let's create a route to handle our first request
First define the route in `http/routes.go`
```go
    router.Get("/", handlers.ExampleShow)
```
Next define the handler function `ExampleShow`, for that let's create the file `http/handlers/example.go` with the following content:
```go
package handlers

import (
"github.com/gin-gonic/gin"
"github.com/my-organization/my-project/http/handlers"
)

func ExampleShow(c *gin.Context) {
    message := "Hello from example handler!"

    c.JSON(200, gin.H{
        "message": message,
    })
}
```
Note: make sure to update the import path of the package `handlers` `github.com/my-organization/my-project/http/handlers` accordingly

Next start the app by running the following code from inside the project:
```bash
gocondor run:dev
```
Finally, open up your browser and navigate to `localhost:8000`

## Folder structure 
```bash
├── gocondor
│   ├── config/ ---------------> control what features to turn on
│   ├── http/------------------> http related code
│   │   ├── handlers/ --------------> contains your http requests handlers
│   │   ├── middlewares/ -----------> middlewares are defined here
│   ├── routes.go -------------> routes are defined here
│   ├── logs/ -----------------> log files
│   ├── models/ ---------------> database models
│   ├── ssl/ ------------------> ssl certificates
│   ├── .env ------------------> environment variables 
│   ├── .gitignore ------------> gitignore file
│   ├── go.mod ----------------> go modules the project depends on
│   ├── LICENSE ---------------> license
│   ├── main.go ---------------> main file
│   ├── README.md -------------> readme file
```
## Architecture
the architecture is similar to `MVC` architecture, there is a `routes.go` file where you can define all your routes and their `handlers`, the handler is simply a method that gets executed when the request is received, you can think of it like a controllers action in `MVC`

### The request journey looks like below:
`request -> routing -> middleware -> handler -> middleware -> json response`


## Contribute
The framework consists of two main parts, each lives in a separate repository, the first part is the `core` which contains the framework core packages. the second part is `gocondor` which has the project folder structure and responsible of gluing everything together.

To contribute you simply need to clone these two repositories locally and create new branches from the `develop` branch, add your changes, then open up a `PR` to the `develop` branch.

Here is how you can clone and set up the development workflow in your local machine:

1. Create the organization `gocondor` directory in your workspace, make sure the full path to it looks like below:
```bash
$GOPATH/src/ginthub.com/gocondor
```
2. clone the repository `core` inside the organization `gocondor` directory:
```bash
git clone git@github.com:gocondor/core.git
```
3. clone the repository `gocondor`:
```bash
git clone git@github.com:gocondor/gocondor.git
```
4. cd into the project `gocondor` and open up `go.mod` in your editor and add the line `github.com/gocondor/gocondor/core => [full-local-path-to-core]` to the `replace`statement, make sure it looks something like this:
```go
module github.com/gocondor/gocondor

replace (
 github.com/gocondor/core => C:/Users/myname/go/src/github.com/gocondor/core

 github.com/gocondor/gocondor/config => ./config
 github.com/gocondor/gocondor/http => ./http
 github.com/gocondor/gocondor/http/middlewares => ./http/middlewares
 github.com/gocondor/gocondor/handlers => ./http/handlers
 github.com/gocondor/gocondor/models => ./models
)
```
Note:
this is needed to tell go that instead of using the remote core package use the local copy where we will be making the changes, once you are done, open a `PR` to `develop` branch.
