package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/chris-miracle/reservations/internal/config"
	"github.com/chris-miracle/reservations/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

func NewTemplate(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(request.Context(), "flash")
	templateData.Error = app.Session.PopString(request.Context(), "error")
	templateData.Warning = app.Session.PopString(request.Context(), "warning")
	
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, request *http.Request, templateName string, templateData *models.TemplateData) error {
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
		// log.Fatal("Template not found in cache")
		return errors.New("Can't get template from cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, request)

	// log.Println(templateData)

	_ = template.Execute(buf, templateData)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmpCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl in the templates directory
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return tmpCache, err
	}

	// loop through the files and add them to the cache
	for _, page := range pages {
		// get the name of the template
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tmpCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return tmpCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return tmpCache, err
			}
		}

		tmpCache[name] = templateSet
	}

	return tmpCache, nil
}
