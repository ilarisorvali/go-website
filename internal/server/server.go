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

func RunServer() error {
	srv := NewServer()
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: srv,
	}

	template.Must(template.ParseGlob("templates/*.html"))

	log.Printf("Starting server on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
	return nil
}
