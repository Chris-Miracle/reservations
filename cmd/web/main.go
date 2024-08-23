package main

import (
	"fmt"
	"net/http"
	"reservations/pgk/handlers"
)

const portNumber = ":4545"

// main is the entry point for the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on localhost%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
