package server

import (
	"net/http"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

const homePath = "./markdown/static/home.md"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadSingleFile(homePath)

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {

	posts := app.postCache

	data := models.TemplateData{
		Items: posts,
	}

	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post := app.postCache[slug]
	app.logger.Debug(slug)

	data := models.TemplateData{
		Item: post,
	}

	app.render(w, r, http.StatusOK, "post.html", data)
}
