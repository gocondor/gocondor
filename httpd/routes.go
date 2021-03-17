// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package httpd

import (
	"github.com/gincoat/gincoat/core/routing"
	"github.com/gincoat/gincoat/httpd/handlers"
)

//RegisterRoutes to register your routes
func RegisterRoutes() {
	router := routing.ResolveRouter()

	//Define your routes here
	router.Get("/home", handlers.HomeShow)
	//router.Get("/update", middlewares.Logger, handlers.HomeUpdate)
}
