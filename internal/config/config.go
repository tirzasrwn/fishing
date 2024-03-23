package config

import (
	"log"
	"text/template"
)

// AppConfig holds the application config.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool

	Port       int
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
}
