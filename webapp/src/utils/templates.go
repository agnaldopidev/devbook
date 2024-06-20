package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Carrega views
func CarregarTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Executa os templates
func ExecuteTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
