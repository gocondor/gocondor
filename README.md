![gocondor logo](https://gocondor.github.io/img/logo.png)
# GoCondor

![Build Status](https://github.com/gocondor/gocondor/actions/workflows/build-main.yml/badge.svg)
![Test Status](https://github.com/gocondor/gocondor/actions/workflows/test-main.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/gocondor/gocondor?status.svg)](https://godoc.org/github.com/gocondor/gocondor)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocondor/gocondor)](https://goreportcard.com/report/github.com/gocondor/gocondor)

## What is GoCondor?
GoCondor is a Golang web framework for building APIs, it features a simple directory structure for your next project with a pleasant & fast development experience, suitable for small & medium size projects and microservices

## Main Features 
- Routing
- Middlewares
- JWT tokens
- Multiple Databases ORM (GORM)
- Cache (Redis)
-  HTTPS (TLS)
- Multiple Environment variables sources (.env file & OS)



## Architecture
The architecture is similar to `MVC`, where there is a routes file `./routes.go` in which you can map all your app routes to their handlers.
Handlers are simply methods that handles the requests to the given routes

#### The request journey:
`Request -> Router -> Optional Middleware -> Handler -> Optional Middleware ->  Response`

## Folder structure 
```bash
├── gocondor
│   ├── config/ --------------------------> main configs
│   ├── handlers/ ------------------------> route's handlers
│   ├── logs/ ----------------------------> app log files
│   ├── middlewares/ ---------------------> app middlewares
│   ├── models/ --------------------------> database models
│   ├── tls/ -----------------------------> tls certificates
│   ├── storage/ -------------------------> a place to store files
│   ├── .env -----------------------------> environment variables 
│   ├── .gitignore -----------------------> .gitignore
│   ├── go.mod ---------------------------> Go modules
│   ├── LICENSE --------------------------> license
│   ├── main.go --------------------------> go main file
│   ├── README.md ------------------------> readme file
│   ├── register-global-middlewares.go ---> register global middlewares
│   ├── routes.go ------------------------> app routes
│   ├── run-auto-migrations.go -----------> database migrations
```

## Installation
To create a new GoCondor project you need to install the `gocondor cli` first

#### Install GoCondor cli
To install the `gocondor cli` globally open up your terminal and run the following command:
```bash
go install github.com/gocondor/installer/gocondor@latest
```

#### Create a new project:
The command for creating a new project is the following:
```bash
gocondor new [project-name] [remote-location]
# example:
# gocondor new my-project github.com/gocondor/my-project
```
where:
`project-name` is the name of your project
`remote-location` is the remote repository that will host the project, usually people use `github.com`


## Getting started
Let's create an app that returns `hello world` as a response.
First create a project with the name `helloapp` by running the following command:
```bash
gocondor new helloapp github.com/gocondor/helloapp
```
Next `cd` into the project and create the file `handlers/greeting-handlers.go` with the following content:

 `#file: handlers/greeting-handlers.go`
```go
package handlers

import "github.com/gocondor/core"

// Returns hello world
func SayHelloWorld(c *core.Context) *core.Response {
	// create a map variable with the desired response
	res := map[string]string{
		"message": "Hello World!",
	}
	// convert the map into JSON
	resJson := c.MapToJson(res)

	// return the response
	return c.Response.WriteJson([]byte(resJson))
}
```
Next, add the route `router.Get("/", handlers.SayHelloWorld)` to the file `./routes.go` like below:

```go
package main

import (
	"github.com/gocondor/core"
	"github.com/gocondor/helloapp/handlers"
)

// Register the app routes
func registerRoutes() {
	router := core.ResolveRouter()
	//#############################
	//# App Routes            #####
	//#############################

	// Define your routes here...
	router.Get("/", handlers.SayHelloWorld)
}
```
Next, build the project by running the following command,in the terminal:
```go
go build
```
this will produce a `binary` file with the name `helloapp` in the root of your project

Next, run the app by executing the binary by running the following command:
```go
./helloapp
```

Finally, open up your browser and navigate to `localhost:8000`.

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
