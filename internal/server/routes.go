package server

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static/"))))
}
