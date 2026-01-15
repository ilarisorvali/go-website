package server

import (
	"html/template"
	"net/http"

	"github.com/ilarisorvali/sorvali-systems/internal/models"
)

var files = []string{
	"templates/base.html",
	"templates/header.html",
	"templates/footer.html",
	"templates/sidebar-left.html",

	"templates/home.html",

	"templates/posts_page.html",
	"templates/post.html",

	"templates/recipes.html",
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	post := models.NewManyContentItems()

	templates, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = templates.ExecuteTemplate(w, "base", post)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loers Laers..."))
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loers Laers..."))
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	app.logger.Info(slug)
	post := models.NewManyContentItems()

	templates, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = templates.ExecuteTemplate(w, "post", post)
	if err != nil {
		app.serverError(w, r, err)
	}
}
