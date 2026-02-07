package server

import (
	"fmt"
	"net/http"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

const homePath = "./markdown/static/home.md"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadSingleFile(homePath)

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	tags := []string{"post"}

	data := models.TemplateData{
		Item:  app.LatestPostFromCache(tags),
		Items: app.FilterCacheByTags(tags),
	}
	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post := app.contentCache[slug]
	tags := []string{"post"}

	next, prev := app.NextPrevPostFromCache(slug, tags)

	data := models.TemplateData{
		Item: post,
		Next: next,
		Prev: prev,
	}

	fmt.Println(data)

	app.render(w, r, http.StatusOK, "post.html", data)
}

func (app *application) recipes(w http.ResponseWriter, r *http.Request) {
	kitchenItems := []string{"kitchen"}

	data := models.TemplateData{
		Items: app.FilterCacheByTags(kitchenItems),
	}
	app.render(w, r, http.StatusOK, "kitchen.html", data)
}
