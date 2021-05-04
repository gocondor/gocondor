package authentication

import "github.com/gocondor/core/routing"

func RegisterAuthRoutes() {
	router := routing.Resolve()

	router.Post("/login", Login)
	router.Get("/logout", Logout)
	router.Post("/register", Register)
}
