package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// This app is solely here so that this repo is go gettable,
// but it'd be nice if the book could host itself locally...
// just load and transform the markup files.

// handler says hello and echoes the request path
// eventually we should serve the markdown of the book converted to html.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// main is the main entry point for your app
func main() {

	// Attach a function to the default ServeMux/Router
	// for the path / and any path under it
	http.HandleFunc("/", handler)

	// Ask the http package to listen on port 3000
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
