package server

import (
	"net/http"
)

//All application routes are defined here

func addRoutes(mux *http.ServeMux) {
	//make static assets available via http.FileServer
	mux.Handle("/static/", http.StripPrefix("/static",
		http.FileServer(http.Dir("static/"))))

	//{$} is a catch-all preventer in go servemux,
	//ie. all unknown subtrees don't match to /
	mux.HandleFunc("/{$}", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/posts", postsHandler)
	mux.HandleFunc("/posts/{slug}", postHandler)
}
