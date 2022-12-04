package utils

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	if err := templates.ExecuteTemplate(w, template, dados); err != nil {
		log.Println(err)
	}
}
