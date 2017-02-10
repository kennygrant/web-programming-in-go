package main

import (
	"crypto/tls"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/acme/autocert"
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

	domains := "example.com www.example.com"
	email := "me@example.com"

	autocertDomains := strings.Split(domains, " ")
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Email:      email,                                      // Email for problems with certs
		HostPolicy: autocert.HostWhitelist(autocertDomains...), // Domains to request certs for
		Cache:      autocert.DirCache("secrets"),               // Cache certs in secrets folder
	}

	server := configuredTLSServer(":443", certManager)
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

// ConfiguredTLSServer returns a TLS server instance with a secure config
// See https://blog.cloudflare.com/exposing-go-on-the-internet/
func configuredTLSServer(port string, certManager *autocert.Manager) *http.Server {

	return &http.Server{
		// Set the port in the preferred string format
		Addr: port,

		// The default server from net/http has no timeouts
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		//	IdleTimeout:  120 * time.Second,

		// This TLS config follows recommendations in the above article
		TLSConfig: &tls.Config{
			// Pass in a cert manager if you want one set
			// this will only be used if the server Certificates are empty
			GetCertificate: certManager.GetCertificate,

			// VersionTLS11 or VersionTLS12 would exclude many browsers
			// inc. Android 4.x, IE 10, Opera 12.17, Safari 6
			// So unfortunately not acceptable as a default yet
			// Current default here for clarity
			MinVersion: tls.VersionTLS10,

			// Causes servers to use Go's default ciphersuite preferences,
			// which are tuned to avoid attacks. Does nothing on clients.
			PreferServerCipherSuites: true,
			// Only use curves which have assembly implementations
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
		},
	}

}
