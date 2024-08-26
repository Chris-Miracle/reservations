package config

import "text/template"

// AppConfig is the application configuration
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}
