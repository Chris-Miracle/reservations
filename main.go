package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4545"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page, the sum is %d", sum)
}

// addValues adds two values and returns the result
func addValues(x, y int) int {
	return  x + y
}


// main is the entry point for the application
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)


	fmt.Printf("Listening on localhost%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
