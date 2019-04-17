package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route describes a specific route to handle unique requests
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of Route types
type Routes []Route

// NewRouter creates a mux.Router from the constant `routes` struct
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"EventHandler",
		"POST",
		"/",
		EventHandler,
	},
	Route{
		"Users",
		"GET",
		"/users",
		UsersHandler,
	},
}
