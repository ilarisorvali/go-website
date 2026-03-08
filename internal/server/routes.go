package server

import (
	"net/http"
)

// All application routes are defined here
// Returns an http handler

func (app *application) addRoutes() *http.ServeMux {
	// new mux, http request multiplexer
	mux := http.NewServeMux()

	// make images available, longer route must come before the shorter
	mux.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir(*app.imageDir))))

	// make static assets available via http.FileServer
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	//{$} is a catch-all preventer in go servemux,
	//ie. all unknown subtrees don't match to /
	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("/posts/{$}", app.posts)
	mux.HandleFunc("/posts/{slug}", app.viewPost)
	mux.HandleFunc("/kitchen/{$}", app.kitchen)
	mux.HandleFunc("/kitchen/{slug}", app.viewPost)
	mux.HandleFunc("/atom.xml", app.feed)

	return mux

}
