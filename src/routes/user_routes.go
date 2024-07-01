package routes

import (
	"net/http"
	"rest/controllers"
)

func RegisterUserRoutes(router *http.ServeMux, parentRoute string) {
	if parentRoute == "" {
		parentRoute = "users"
	}
	router.HandleFunc(RouteMaker("GET", parentRoute, ""), controllers.GetUser)
	router.HandleFunc(RouteMaker("POST", parentRoute, ""), controllers.CreateUser)
}
