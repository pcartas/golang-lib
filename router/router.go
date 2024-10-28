package router

import (
    "net/http"
    "github.com/gorilla/mux"
)

// Route defines a single route, e.g. a human readable name, HTTP method and the pattern the function that will execute when the route is called.
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
    Middlewares []mux.MiddlewareFunc
}

// Routes is a slice of Route structs.
type Routes []Route

// NewRouter creates a new router with the provided routes and middlewares.
func NewRouter(routes Routes, generalMiddlewares []mux.MiddlewareFunc) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    // Add general middlewares to the router
    for _, middleware := range generalMiddlewares {
        router.Use(middleware)
    }

    // Register routes
    for _, route := range routes {
        r := router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
        
        // Add specific middlewares to the route
        for _, middleware := range route.Middlewares {
            r.Handler(middleware(r.GetHandler()))
        }
    }

    return router
}