package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/chris-miracle/reservations/internal/config"
	"github.com/chris-miracle/reservations/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewTemplate(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, request *http.Request, templateName string, templateData *models.TemplateData) {
	var tmpCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from app config
		tmpCache = app.TemplateCache
	} else {
		tmpCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	template, ok := tmpCache[templateName]
	if !ok {
		log.Fatal("Template not found in cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, request)

	// log.Println(templateData)

	_ = template.Execute(buf, templateData)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmpCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl in the templates directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmpCache, err
	}

	// loop through the files and add them to the cache
	for _, page := range pages {
		// get the name of the template
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmpCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tmpCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmpCache, err
			}
		}

		tmpCache[name] = templateSet
	}

	return tmpCache, nil
}
