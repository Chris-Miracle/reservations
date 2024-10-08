package handlers

import (
	"net/http"

	"github.com/chris-miracle/reservations/pgk/config"
	"github.com/chris-miracle/reservations/pgk/models"
	"github.com/chris-miracle/reservations/pgk/render"
)

// Repo is the repository for the handlers
var Repo *Repository

// Repository is the repository for the handlers
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{App: appConfig}
}

// NewHandlers creates the handlers for the application
func NewHandlers(repository *Repository) {
	Repo = repository
}

// Home is the handler for the home page
func (repository *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	repository.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (repository *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hi Chris."

	remoteIP := repository.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the handler for the reservation page
func (repository *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals is the handler for the generals page
func (repository *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the handler for the majors page
func (repository *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Availability is the handler for the availability page
func (repository *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact is the handler for the contact page
func (repository *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}


