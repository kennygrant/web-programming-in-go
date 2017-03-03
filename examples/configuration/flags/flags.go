package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

// Config holds our app config
type Config struct {
	Port int
}

var config = &Config{Port: 3000}

// handler says hello and echoes the request path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// main is the main entry point for your app
func main() {

	// Read the port with the flag package into a ptr to int
	flag.IntVar(&config.Port, "port", 3000, "The port the server listens on (default 3000)")
	flag.Parse()

	fmt.Println("Starting server on port:", config.Port)

	// Attach a function to the default ServeMux/Router
	http.HandleFunc("/", handler)

	// Ask the http package to listen on port 3000
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
