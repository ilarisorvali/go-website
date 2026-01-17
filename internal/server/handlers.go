package server

import (
	"net/http"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

var postPath string = "./markdown"
var homePath string = "./markdown/home/home.md"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadSingleMarkdownFile(homePath)

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadMarkdownFiles(postPath)

	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	data, _ := models.LoadMarkdownFiles(postPath)
	app.logger.Debug(slug)

	var found models.ContentItem

	for _, item := range data.Items {
		if item.Meta.Title == slug && !item.Meta.Draft {
			found = item
			break
		}
	}

	data.Item = found

	app.render(w, r, http.StatusOK, "post.html", data)
}
