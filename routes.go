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

	for _, route := range AllRoutes(cr) {
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

func AllRoutes(cr *CustomRedis) []Route {
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
			Description: "Set string key and value if key doesn`t exist",
			Method:      "POST",
			Path:        "/strnx",
			HandlerFunc: cr.SetStrNX,
		},
		Route{
			Description: "Get string by key",
			Method:      "GET",
			Path:        "/str",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.GetStr,
		},
		Route{
			Description: "Set list key and value",
			Method:      "POST",
			Path:        "/list",
			HandlerFunc: cr.SetList,
		},
		Route{
			Description: "Get list by key",
			Method:      "GET",
			Path:        "/list",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.GetList,
		},
		Route{
			Description: "Get list by key",
			Method:      "POST",
			Path:        "/listpush",
			HandlerFunc: cr.PushList,
		},
		Route{
			Description: "Removes and returns the last element of the list stored at key",
			Method:      "GET",
			Path:        "/listpop",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.ListPop,
		},
		Route{
			Description: "Returns the element at index in the list stored at key",
			Method:      "GET",
			Path:        "/listindex",
			QueryPairs:  "key,{key},index,{index}",
			HandlerFunc: cr.ListIndex,
		},
		Route{
			Description: "Set map key and value",
			Method:      "POST",
			Path:        "/map",
			HandlerFunc: cr.SetMap,
		},
		Route{
			Description: "Get map by key",
			Method:      "GET",
			Path:        "/map",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.GetMap,
		},
		Route{
			Description: "Delete value by key",
			Method:      "DELETE",
			Path:        "/del",
			QueryPairs:  "key,{key}",
			HandlerFunc: cr.Delete,
		}}
}
