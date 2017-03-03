package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
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
	var err error

	// Read the port from the environment variable, if set
	portEnv := os.Getenv("MY_SERVER_PORT")
	if len(portEnv) > 0 {
		// Convert string to an int
		config.Port, err = strconv.Atoi(portEnv)
		if err != nil {
			log.Print(err)
		}
	}

	fmt.Println("Starting server on port:", config.Port)

	// Attach a function to the default ServeMux/Router
	http.HandleFunc("/", handler)

	// Ask the http package to listen on port 3000
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
