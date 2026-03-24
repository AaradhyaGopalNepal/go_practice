package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Root GET"))
		fmt.Println("Root GET")
	case http.MethodPost:
		w.Write([]byte("Root POST"))
		fmt.Println("Root POST")
	case http.MethodPut:
		w.Write([]byte("Root PUT"))
		fmt.Println("Root PUT")
	case http.MethodPatch:
		w.Write([]byte("Root PATCH"))
		fmt.Println("Root PATCH")
	case http.MethodDelete:
		w.Write([]byte("Root DELETE"))
		fmt.Println("Root DELETE")
	}
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}
func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Teacher GET"))
		fmt.Println("Teacher GET")
	case http.MethodPost:
		w.Write([]byte("Teacher POST"))
		fmt.Println("Teacher POST")
	case http.MethodPut:
		w.Write([]byte("Teacher PUT"))
		fmt.Println("Teacher PUT")
	case http.MethodPatch:
		w.Write([]byte("Teacher PATCH"))
		fmt.Println("Teacher PATCH")
	case http.MethodDelete:
		w.Write([]byte("Teacher DELETE"))
		fmt.Println("Teacher DELETE")
	}
	w.Write([]byte("Hello Teacher Route"))
	fmt.Println("Hello Teacher Route")
}
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Students GET"))
		fmt.Println("Students GET")
	case http.MethodPost:
		w.Write([]byte("Students POST"))
		fmt.Println("Students POST")
	case http.MethodPut:
		w.Write([]byte("Students PUT"))
		fmt.Println("Students PUT")
	case http.MethodPatch:
		w.Write([]byte("Students PATCH"))
		fmt.Println("Students PATCH")
	case http.MethodDelete:
		w.Write([]byte("Students DELETE"))
		fmt.Println("Students DELETE")
	}
	w.Write([]byte("Hello Students Route"))
	fmt.Println("Hello Students Route")
}
func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Execs GET"))
		fmt.Println("Execs GET")
	case http.MethodPost:
		w.Write([]byte("Execs POST"))
		fmt.Println("Execs POST")
	case http.MethodPut:
		w.Write([]byte("Execs PUT"))
		fmt.Println("Execs PUT")
	case http.MethodPatch:
		w.Write([]byte("Execs PATCH"))
		fmt.Println("Execs PATCH")
	case http.MethodDelete:
		w.Write([]byte("Execs DELETE"))
		fmt.Println("Execs DELETE")
	}
	w.Write([]byte("Hello Execs Route"))
	fmt.Println("Hello Execs Route")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	fmt.Println("Server is running")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(
			"Error starting the server", err,
		)
	}
}
