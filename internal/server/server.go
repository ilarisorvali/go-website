package server

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	contentCache  map[string]*models.ContentItem
}

func RunServer() error {
	addr := flag.String("addr", ":9000", "HTTP port for the server to use")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init template cache for application, check for errors		s
	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// init content cache for application, check for errors
	cache, err := newContentCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		contentCache:  cache,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.addRoutes(),
	}

	logger.Info("starting server", "addr", addr)
	err = srv.ListenAndServe()
	logger.Error(err.Error())

	return nil
}
