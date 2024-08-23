package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4545"

// main is the entry point for the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Listening on localhost%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
