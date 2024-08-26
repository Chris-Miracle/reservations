package handlers

import (
	"net/http"
	"reservations/pgk/config"
	"reservations/pgk/models"
	"reservations/pgk/render"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (repository *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hi Chris."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StrinpMap: stringMap,
	})
}
