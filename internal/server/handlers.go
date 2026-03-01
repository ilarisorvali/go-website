package server

import (
	"fmt"
	"net/http"

	models "github.com/ilarisorvali/go-website/internal/content"
)

const homePath = "./markdown/static/home.md"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadSingleFile(homePath)

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	//tags := []string{"post"}

	data := models.TemplateData{
		Items: app.contentCache.Blog,
	}
	fmt.Println(app.contentCache.Blog)
	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post := app.contentCache.Items[slug]
	tags := []string{"post"}

	next, prev := app.NextPrevPostFromCache(slug, tags)

	data := models.TemplateData{
		Item: post,
		Next: next,
		Prev: prev,
	}

	app.render(w, r, http.StatusOK, "post.html", data)
}

func (app *application) kitchen(w http.ResponseWriter, r *http.Request) {

	data := models.TemplateData{
		Items: app.contentCache.Kitchen,
	}
	app.render(w, r, http.StatusOK, "kitchen.html", data)
}
