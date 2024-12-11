package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chris-miracle/reservations/internal/config"
	"github.com/chris-miracle/reservations/internal/handlers"
	"github.com/chris-miracle/reservations/internal/models"
	"github.com/chris-miracle/reservations/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":4545"

var app config.AppConfig
var session *scs.SessionManager

// main is the entry point for the application
func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on localhost%s \n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gob.Register(models.Reservation{})

	app.InProduction = false

	//  Sessiion Props
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache")
		return err
	}

	app.TemplateCache = templateCache
	// set to false for development mode
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	return nil
}
