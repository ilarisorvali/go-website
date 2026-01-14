package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func NewServer() http.Handler {
	//It is generally safer to declare and use your own servemux rather than
	//use the default one. All third party packages could in theory register
	//handlers to routes in the http.DefaultServeMux
	mux := http.NewServeMux()
	//Register mux as the http handler
	var handler http.Handler = mux

	//Add all application routes to mux in routes.go
	addRoutes(mux)

	return handler
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
