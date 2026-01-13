package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux
	addRoutes(mux)

	return handler
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templ, err := template.ParseFiles("views/pages/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	templ.Execute(w, nil)
}

func RunServer() error {
	srv := NewServer()
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: srv,
	}

	log.Printf("Starting server on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
	return nil
}
