package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate renders a template
func RenderTemplateTest(w http.ResponseWriter, templateName string) {
	parsedTenplate, _ := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	err := parsedTenplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error

	// check to see if the template is in cache
	_, imMap := tc[templateName]
	if !imMap {
		// create the template
		log.Println("creating template cache")
		err = createTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in cache
		log.Println("using template from cache")
	}

	tmpl = tc[templateName]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(templateName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache map
	tc[templateName] = tmpl
	return nil
}
