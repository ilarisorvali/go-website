package server

import (
	"net/http"
)

//All application routes are defined here

func (app *application) addRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	//make static assets available via http.FileServer
	mux.Handle("/static/", http.StripPrefix("/static",
		http.FileServer(http.Dir("static/"))))

	//{$} is a catch-all preventer in go servemux,
	//ie. all unknown subtrees don't match to /
	mux.HandleFunc("/{$}", app.home)
	mux.HandleFunc("/posts/{$}", app.posts)
	mux.HandleFunc("/posts/{slug}", app.viewPost)

	return mux

}
