package main

import (
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route Define a HTTP Route with given logical name, http method, route pattern and handler function
type Route struct {
	Description string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes Describe all service API
type Routes []Route

// NewRouter creates a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = loggerHandler(handler, route.Description)
		handler = handlers.CompressHandler(handler)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Description).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		Description: "Health check",
		Method: "GET",
		Path: "/",
		HandlerFunc: Heartbit,
	},
}