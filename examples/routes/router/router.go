package router

import (
	"net/http"
	"strings"
	"sync"
)

// Route is a simplistic representation of a route
type Route struct {
	Path    string
	Handler http.HandlerFunc
}

// Router routes requests
type Router struct {
	mu     sync.RWMutex
	routes []*Route
}

// New returns a new router
func New() *Router {
	return &Router{}
}

// Add a route to our list of routes to evaluate
func (r *Router) Add(p string, fn http.HandlerFunc) {
	r.mu.Lock()
	r.routes = append(r.routes, &Route{Path: p, Handler: fn})
	r.mu.Unlock()
}

// ServeHTTP serves all HTTP requests
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Check the routes
	found := false
	r.mu.RLock()
	for _, route := range r.routes {
		// Return a match if our route prefix matches path
		if strings.HasPrefix(req.URL.Path, route.Path) {
			route.Handler(w, req)
			found = true
			break
		}
	}
	r.mu.RUnlock()

	// No route found, render 404
	if !found {
		http.Error(w, "404 page not found", http.StatusNotFound)
	}
}
