package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}
	return r
}
