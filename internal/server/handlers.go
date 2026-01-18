package server

import (
	"net/http"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	post := app.contentCache["home"]

	data := models.TemplateData{
		Item: post,
	}

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {

	posts := app.contentCache

	data := models.TemplateData{
		Items: posts,
	}

	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post := app.contentCache[slug]
	app.logger.Debug(slug)

	data := models.TemplateData{
		Item: post,
	}

	app.render(w, r, http.StatusOK, "post.html", data)
}
