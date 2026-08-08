package server

import (
	"net/http"
)

// All application routes are defined here
// Returns an http handler containing the routes

func (app *application) addRoutes() *http.ServeMux {
	// new mux, http request multiplexer (router)
	mux := http.NewServeMux()

	// make images available with a FileServer route
	mux.Handle("GET /images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir(*app.imageDir))))

	// make static assets available via http.FileServer
	mux.Handle("GET /static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	//{$} is a catch-all preventer in go servemux,
	//ie. all unknown subtrees don't match to /
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /posts/{$}", app.posts)
	mux.HandleFunc("GET /posts/{slug}", app.viewPost)
	mux.HandleFunc("GET /kitchen/{$}", app.kitchen)
	mux.HandleFunc("GET /kitchen/{slug}", app.viewPost)
	mux.HandleFunc("GET /atom.xml", app.feed)

	return mux

}
