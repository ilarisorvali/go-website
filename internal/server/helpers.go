package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"

	"github.com/ilarisorvali/sorvali-systems/internal/models"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data models.Pages) {

	//Retrieve template from templatecache with page name,
	//raises error if template does not exist
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}
	// Write out the provided HTTP status code ('200 OK', '400 Bad Request' etc).
	w.WriteHeader(status)
	// Execute the 'page' template set
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	//Init map to act as the template cache
	cache := map[string]*template.Template{}

	//Get slice of page filepath strings
	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.html')
		name := filepath.Base(page)

		// This essentially merges together individual pages with base.html and partials
		// into many named templates that the router then calls based on the name of the template
		files := []string{
			"./ui/html/templates/base.html",
			"./ui/html/partials/header.html",
			page,
		}

		// Parse the files into a template set.
		ts, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
