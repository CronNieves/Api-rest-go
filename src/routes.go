package main

import (
	"./handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"movies",
		"GET",
		"/movies",
		handlers.GetMovieList,
	},
	Route{
		"movies",
		"GET",
		"/movie/{id}",
		handlers.GetMovie,
	},
}
