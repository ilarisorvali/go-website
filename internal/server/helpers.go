package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data models.TemplateData) {

	// Retrieve template from templatecache with page name,
	// raises error if template does not exist
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	// Write template to placeholder buffer instead of
	// http.Responsewriter to catch ExecuteTemplate() errors
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	// Write contents of buffer to http.ResponseWriter
	buf.WriteTo(w)
}

// TODO add error handling
func newContentCache() (map[string]*models.ContentItem, error) {
	// kind indicates what kind of content to load ie. Post or Recipe
	kind := models.Post
	// init an empty map ot ac as the ContentItem cache
	cache, err := models.LoadMarkdownFiles("./markdown", kind)
	if err != nil {
		return nil, err
	}
	return cache, nil
}

func newTemplateCache() (map[string]*template.Template, error) {
	//Init an empty map to act as the template cache
	cache := map[string]*template.Template{}

	//Get slice of page filepath strings
	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.html')
		name := filepath.Base(page)

		// Parse the base template
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// ParseGlob() on current template set to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// ParseFiles() on current template set to add the current page template
		ts, err = ts.ParseFiles(page)
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
