package server

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	models "github.com/ilarisorvali/go-website/internal/content"
)

type application struct {
	logger        *slog.Logger //universal logger that can be called wherever application is present
	templateCache map[string]*template.Template
	contentCache  *models.ContentCache
	imageDir      *string
	atomFeed      *[]byte
}

func RunServer() error {
	// Get the server port as a flag, default to 9000
	addr := flag.String("addr", ":9000", "HTTP port for the server to use")
	// Get the markdown content directory as a flag, default to .markdown
	contentDir := flag.String("contentdir", "./markdown", "directory to load markdown content from")
	// Get the image content directory as a flag, default to .markdown/images
	imageDir := flag.String("imagedir", "./images", "directory to load images content from")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init template cache for application, check for errors
	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// init content cache for application, check for errors
	cache, err := newContentCache(contentDir)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		contentCache:  cache,
		imageDir:      imageDir,
	}

	app.generateAtomFeed()

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.addRoutes(),
	}

	logger.Info("starting server", "addr", *addr)
	err = srv.ListenAndServe()
	logger.Error(err.Error())

	return nil
}
