package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// NewLoggedRouter returns a gorilla/mux router with logging and all routes
// defined for the webservice
func NewLoggedRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true) // "/path" and "/path/" will be viewed as the same path

	for _, r := range getRoutes() {
		var handler http.Handler

		handler = r.HandlerFunc
		handler = logger(handler, r.Name)

		router.Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Queries(r.Queries...).
			Handler(handler)
	}

	return router
}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Queries     []string
}

func getRoutes() []route {
	return []route{
		route{
			"MapSearch",
			"GET",
			"/maps",
			MapSearch,
			[]string{"x", "{x:[0-9]+}", "y", "{y:[0-9]+}"},
		},
	}
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
