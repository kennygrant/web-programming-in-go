package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

// handler says hello and echoes the request path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// main is the main entry point for your app
func main() {

	// Attach a function to the default ServeMux/Router
	// for the path / and any path under it
	http.HandleFunc("/", handler)

	// Set up a new http server with some default timeouts and port
	server := &http.Server{
		// Set the port in the preferred string format
		Addr: ":3000",
		// The default server from net/http has no timeouts
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// Start the server with our self-signed cert and key
	// the keys provided are examples with no other use
	// don't use self-signed keys on a real server
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
