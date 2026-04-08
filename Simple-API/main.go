package main

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		greeting := "Hello, " + name + "!"
		w.Write([]byte(greeting))
	})

	// Load the TLS certificate and key files
	cert := "cert.pem"
	key := "key.pem"

	// Configure the server to use TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}

	// Enable HTTP2
	http2.ConfigureServer(server, &http2.Server{})

	println("Server is running on port", server.Addr)
	err := server.ListenAndServeTLS(cert , key)
	if err != nil {
		panic(err)
	}


}