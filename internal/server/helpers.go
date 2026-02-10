package server

// This file contains a set of functions that don't yet have a better home

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"
	"slices"
	"time"

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
	// init an empty map ot ac as the ContentItem cache
	cache, err := models.LoadContentFiles("./markdown")
	if err != nil {
		return nil, err
	}
	return cache, nil
}

// A helper function for further post filtering in the templates
func hasTag(tags []string, tag string) bool {
	exists := slices.Contains(tags, tag)
	return exists
}

// Needed to register custom functions for the templates
var functions = template.FuncMap{
	"hasTag": hasTag,
	"Year": func() int {
		return time.Now().Year()
	},
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

		// Create blank template, add functions and parse the base template
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// ParseGlob() on base template set to add partials (navbar, footer, etc.)
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// ParseFiles() on current template set to add the current page template contents to the base
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

func tagSet(tags []string) map[string]struct{} {
	set := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		set[tag] = struct{}{}
	}
	return set
}

// Filter a map[string]*ContentItem by ContentItem tags.
// Makes use of the tagset function for ~speed~
func (app *application) FilterCacheByTags(tags []string) map[string]*models.ContentItem {
	required := tagSet(tags)
	result := make(map[string]*models.ContentItem)

	for slug, item := range app.contentCache {
		itemTags := tagSet(item.Meta.Tags)
		match := true
		for tag := range required {
			if _, ok := itemTags[tag]; !ok {
				match = false
				break
			}
		}

		if match && !item.Meta.Draft {
			result[slug] = item
		}
	}

	return result
}

func (app *application) LatestPostFromCache(tags []string) *models.ContentItem {
	var latest *models.ContentItem
	matchingItems := app.FilterCacheByTags(tags)

	for _, item := range matchingItems {
		if item.Meta.Date.IsZero() {
			continue
		}

		if latest == nil || item.Meta.Date.After(latest.Meta.Date.Time) {
			latest = item
		}
	}

	return latest
}

func (app *application) NextPrevPostFromCache(slug string, tags []string) (next, prev *models.ContentItem) {
	matchingItems := app.FilterCacheByTags(tags)
	currentPost, ok := matchingItems[slug]
	if !ok {
		return nil, nil
	}

	currentTime := currentPost.Meta.Date.Time

	for _, item := range matchingItems {
		if item.Meta.Date.After(currentTime) {
			if next == nil || item.Meta.Date.Before(next.Meta.Date.Time) {
				temp := item
				next = temp
			}

		}
		if item.Meta.Date.Before(currentTime) {
			if prev == nil || item.Meta.Date.After(prev.Meta.Date.Time) {
				temp := item
				prev = temp
			}

		}

	}
	return next, prev
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
