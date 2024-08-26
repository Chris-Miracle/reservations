package main

import (
	"fmt"
	"log"
	"net/http"
	"reservations/pgk/config"
	"reservations/pgk/handlers"
	"reservations/pgk/render"
)

const portNumber = ":4545"

// main is the entry point for the application
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache")
	}

	app.TemplateCache = templateCache
	// set to false for development mode
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listening on localhost%s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
