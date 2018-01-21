package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route Define a HTTP Route with given logical name, http method, route pattern and handler function
type Route struct {
	Description string
	Method      string
	Path        string
	QueryPairs  string
	HandlerFunc http.HandlerFunc
}

// Routes Describe all service API
type Routes []Route

// NewRouter creates a new router
func NewRouter(cr *CustomRedis) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range allRoutes(cr) {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = loggerHandler(handler, route.Description)
		handler = handlers.CompressHandler(handler)

		if route.QueryPairs == "" {
			router.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Description).
				Handler(handler)
		} else {
			qp := strings.Split(route.QueryPairs, ",")
			router.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Description).
				Handler(handler).
				Queries(qp...)
		}
	}

	return router
}

func allRoutes(cr *CustomRedis) []Route {
	return []Route{
		{
			Description: "Returns a simple response to check if server is alive",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: Index,
		},
		Route{
			Description: "Set string key and value",
			Method:      "POST",
			Path:        "/str",
			HandlerFunc: cr.SetStr,
		},
		Route{
			Description: "Get value by key",
			Method:      "GET",
			Path:        "/str",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.GetStr,
		}}
}
