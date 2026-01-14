package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/base.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/post.html",
	"templates/home.html",
))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "base")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loers Laers..."))
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loers Laers..."))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	msg := fmt.Sprintf("Slug extracted: %s", slug)
	w.Write([]byte(msg))
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
