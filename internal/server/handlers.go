package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/post.html",
))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "base")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
