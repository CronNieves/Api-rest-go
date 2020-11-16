package main

import (
	"github.com/CronNieves/Api-rest-go/src/handlers"
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
		"moviesList",
		"GET",
		"/movies",
		handlers.GetMovieList,
	},
	Route{
		"moviesGet",
		"GET",
		"/movie/{id}",
		handlers.GetMovie,
	},
	Route{
		"moviesAdd",
		"POST",
		"/movie",
		handlers.MovieAdd,
	}, Route{
		"movieUpdate",
		"PUT",
		"/movie/{id}",
		handlers.MovieUpdate,
	},
}
