package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTenplate, _ := template.ParseFiles("./templates/" + templateName)
	err := parsedTenplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
