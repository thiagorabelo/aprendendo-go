package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "joao.pedro@gmail.com"}
		templates.ExecuteTemplate(w, "index.html", u)
	})

	fmt.Println("Executando o servidor na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
