package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// NoSurf prevents cross-site request forgery attacks
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
