package main

import (
	"fmt"
	"log"
	"net/http"
)

// To generate key: openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem
func main() {
	//http vs https
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders")
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming users")
	})
	port := 3000
	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln("Could not start server", err)
	}
}
