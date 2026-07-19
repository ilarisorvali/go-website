package server

import (
	"net/http"

	models "github.com/ilarisorvali/go-website/internal/content"
)

const homePath = "./markdown/static/home.md"

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	//empty content data struct passed as no data is needed on the page
	app.render(w, r, http.StatusOK, "home.html", models.TemplateData{})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {

	//empty content data struct passed as no data is needed on the page
	app.render(w, r, http.StatusOK, "about.html", models.TemplateData{})
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {

	data := models.TemplateData{
		Items: app.contentCache.Blog,
	}
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

func (app *application) feed(w http.ResponseWriter, r *http.Request) {
	// Set response headers for readers
	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.Write(*app.atomFeed)
}
