package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/kennygrant/web-programming-in-go/examples/routes/router"
)

// handler says hello and echoes the request path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// main is the main entry point for your app
func main() {

	// Create our router
	r := router.New()
	r.Add("/foo", handler)

	// Set the router to handle requests
	http.Handle("/", r)

	// Ask the http package to listen on port 3000
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
