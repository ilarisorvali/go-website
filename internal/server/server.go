package server

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	models "github.com/ilarisorvali/sorvali-systems/internal/content"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	postCache     map[string]*models.ContentItem
	recipeCache   map[string]*models.ContentItem
}

func RunServer() error {
	addr := ":9000"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init template cache for application, check for errors		s
	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// init content cache for application, check for errors
	contentCache, err := newContentCache(models.Post)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// init content cache for application, check for errors
	recipeCache, err := newContentCache(models.Post)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		postCache:     contentCache,
		recipeCache:   recipeCache,
	}

	mux := app.addRoutes()

	logger.Info("starting server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
		os.Exit(1)
	}
	return nil
}
