package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

// NewRouter creates a new router which handles
// what to do when a given route is called and
// passes this information to be used by the
// handler.
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}
