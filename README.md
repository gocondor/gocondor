![gocondor logo](https://github.com/gocondor/gocondor.github.io/raw/master/img/logo.png)
# GoCondor

![Build Status](https://github.com/gocondor/gocondor/actions/workflows/build-master.yml/badge.svg)
![Test Status](https://github.com/gocondor/gocondor/actions/workflows/test-master.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/gocondor/gocondor?status.svg)](https://godoc.org/github.com/gocondor/gocondor)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocondor/gocondor)](https://goreportcard.com/report/github.com/gocondor/gocondor)

## What is GoCondor?
GoCondor is a golang web framework with an `MVC` like architecture, it's based on [Gin framework](https://github.com/gin-gonic/gin), it features a simple organized directory structure for your next project with a pleasant development experience, made for developing modern APIs and microservices.

## Features 
- Router
- Routing Groups
- Middlewares
- JWT tokens
- ORM (GORM)
- Sessions
- Authentication
- Cache (Redis)
- TLS
- Live-Reloading for development
- Features Control

## Architecture
The architecture is similar to `MVC`, where there is a routes file `http/routes.go` in which you can map all your app routes to their handlers.
Handlers are simply methods that get executed when the matching request is received, you can think of it like a controller's action in `MVC`.

#### The request journey:
`Request -> Routing -> Middleware -> Handler -> Middleware -> Json Response`

## Folder structure 
```bash
├── condor
│   ├── config/ ---------------> control what features to turn on
│   ├── httpd/-----------------> http related code
│   │   ├── handlers/ --------------> contains your requests handlers
│   │   ├── middlewares/ -----------> middlewares are defined here
│   ├── routes.go -------------> routes are mapped to their handlers here
│   ├── logs/ -----------------> logs file is here
│   ├── models/ ---------------> database models
│   ├── ssl/ ------------------> ssl certificates goes here
│   ├── .env ------------------> environment variables 
│   ├── .gitignore ------------> .gitignore file
│   ├── go.mod ----------------> Go modules that project depends on
│   ├── LICENSE ---------------> license
│   ├── main.go ---------------> main file
│   ├── README.md -------------> readme file
```

## Installation
To create a new GoCondor project you need to install the `gocondor cli` first

#### Install GoCondor cli
To install the `gocondor cli` globally open up your terminal and run the following command:
```bash
go get github.com/gocondor/installer/gocondor
```

#### Create a new project:
The command for creating a new project is the following:
```bash
gocondor new [project-name] [project-location]
```
where:
`project-name` is the name of your project
`project-location` is the remote repository that will host the project, usually people use `github.com`

Now let's create a project with the name `todo` and let's assume it's hosted on the repository `github.com/my-organization/todo`, here is the command to create that project
```bash
gocondor new todo github.com/my-organization/todo
```

## Getting started
Let's add the route `/hello`, and lets have `hello there!` as the response.
To do that Open the file `http/routes.go` in your editor, update the function `RegisterRoutes()`, make sure the it looks like below:
```go
func RegisterRoutes() {
    router := routing.Resolve()

    // Define your routes here
    router.Get("/hello", func(c *gin.Context) {
        message := "hello there!"

        c.JSON(http.StatusOK, gin.H{
            "message": message,
        })
    })
}
```
Next cd into the project folder and start the app by running the following command:
```bash
go run main.go
```
or you can start it using [Air](https://github.com/cosmtrek/air)
```bash
air main.go
```
Finally, open up your browser and navigate to `localhost:8000/hello`.

To learn how to create handlers files and how to add handlers to them check [handlers docs](https://gocondor.github.io/docs/handlers)


## Contribute
The framework consists of two main parts, each lives in a separate repository, the first part is the `core` which contains the framework core packages. the second part is `gocondor` which has the project folder structure and responsible of gluing everything together.

To contribute you simply need to clone these two repositories locally and create new branches from the `develop` branch, add your changes, then open up a `PR` to the `develop` branch.

Here is how you can clone and set up the development workflow in your local machine:

1. Create the organization `gocondor` directory in your workspace, make sure the full path to it looks like below:
```bash
$GOPATH/src/ginthub.com/gocondor
```
2. clone the repository `core` inside the organization `gocondor` directory:
```bash
git clone git@github.com:gocondor/core.git
```
3. clone the repository `gocondor`:
```bash
git clone git@github.com:gocondor/gocondor.git
```
4. cd into the project `gocondor` and open up `go.mod` in your editor and add the line `github.com/gocondor/gocondor/core => [full-local-path-to-core]` to the `replace`statement, make sure it looks something like this:
```go
module github.com/gocondor/gocondor

replace (
 github.com/gocondor/core => C:/Users/myname/go/src/github.com/gocondor/core

 github.com/gocondor/gocondor/config => ./config
 github.com/gocondor/gocondor/http => ./http
 github.com/gocondor/gocondor/http/middlewares => ./http/middlewares
 github.com/gocondor/gocondor/handlers => ./http/handlers
 github.com/gocondor/gocondor/models => ./models
)
```
Note:
this is needed to tell go that instead of using the remote core package use the local copy where we will be making the changes, once you are done, open a `PR` to `develop` branch.
