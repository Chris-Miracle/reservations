package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":4545"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	parsedTenplate, _ := template.ParseFiles("./templates/" + templateName)
	err := parsedTenplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// main is the entry point for the application
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Listening on localhost%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
