package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// To generate key: openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem
func main() {
	//http vs https
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)

		fmt.Fprintf(w, "Handling incoming users")
	})
	port := 5050

	cert := "cert.pem"
	key := "key.pem"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS11,
	}
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	//Enable http2
	http2.ConfigureServer(server, &http2.Server{})
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}
	//HTTP 1
	// fmt.Println("Server is running on port:", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Version:", tlsVersion)
	} else {
		fmt.Println("Without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "v10"
	case tls.VersionTLS11:
		return "v11"
	case tls.VersionTLS12:
		return "v12"
	case tls.VersionTLS13:
		return "v13"
	default:
		return "Not found"
	}

}
