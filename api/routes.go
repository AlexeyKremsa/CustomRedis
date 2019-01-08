package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type route struct {
	description string
	method      string
	path        string
	queryPairs  string
	handlerFunc http.HandlerFunc
}

type routes []route

func newRouter(srv *server) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range allRoutes(srv) {
		var handler http.Handler
		handler = route.handlerFunc
		handler = loggerHandler(handler, route.description)
		handler = handlers.CompressHandler(handler)

		if route.queryPairs == "" {
			router.
				Methods(route.method).
				Path(route.path).
				Name(route.description).
				Handler(handler)
		} else {
			qp := strings.Split(route.queryPairs, ",")
			router.
				Methods(route.method).
				Path(route.path).
				Name(route.description).
				Handler(handler).
				Queries(qp...)
		}
	}

	return router
}

func allRoutes(srv *server) []route {
	return []route{
		{
			description: "Returns a simple response to check if server is alive",
			method:      "GET",
			path:        "/",
			handlerFunc: Index,
		},
		route{
			description: "Set string key and value",
			method:      "POST",
			path:        "/str",
			handlerFunc: srv.setStr,
		},
		route{
			description: "Set string key and value if key doesn`t exist",
			method:      "POST",
			path:        "/strnx",
			handlerFunc: srv.setStrNX,
		},
		route{
			description: "Get string by key",
			method:      "GET",
			path:        "/str",
			queryPairs:  "key,{key}",
			handlerFunc: srv.getStr,
		},
		route{
			description: "Set list key and value",
			method:      "POST",
			path:        "/list",
			handlerFunc: srv.setList,
		},
		route{
			description: "Get list by key",
			method:      "GET",
			path:        "/list",
			queryPairs:  "key,{key}",
			handlerFunc: srv.getList,
		},
		route{
			description: "Add elements to the end of the list",
			method:      "POST",
			path:        "/listinsert",
			handlerFunc: srv.listInsert,
		},
		route{
			description: "Removes and returns the last element of the list stored at key",
			method:      "GET",
			path:        "/listpop",
			queryPairs:  "key,{key}",
			handlerFunc: srv.listPop,
		},
		route{
			description: "Returns the element at index in the list stored at key",
			method:      "GET",
			path:        "/listindex",
			queryPairs:  "key,{key},index,{index}",
			handlerFunc: srv.listIndex,
		},
		route{
			description: "Set map key and value",
			method:      "POST",
			path:        "/map",
			handlerFunc: srv.setMap,
		},
		route{
			description: "Get map by key",
			method:      "GET",
			path:        "/map",
			queryPairs:  "key,{key}",
			handlerFunc: srv.getMap,
		},
		route{
			description: "Get map item by key",
			method:      "GET",
			path:        "/mapitem",
			queryPairs:  "key,{key},itemKey,{itemKey}",
			handlerFunc: srv.getMapItem,
		},
		route{
			description: "Delete value by key",
			method:      "DELETE",
			path:        "/del",
			queryPairs:  "key,{key}",
			handlerFunc: srv.Delete,
		},
		route{
			description: "Get all keys",
			method:      "GET",
			path:        "/keys",
			handlerFunc: srv.Keys,
		}}
}
