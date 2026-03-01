---
Title: Building a website in go, a novice approach
Description: Test Description
Slug: test-post
Date: 28.03.2025
Tags: [post]
---

# Testing features

This is how the server is run:[^1]

```go
func RunServer() error {
	// Get the server port as a flag, default to 9000
	addr := flag.String("addr", ":9000", "HTTP port for the server to use")
	// Get the markdown content directory as a flag, default to .markdown
	contentDir := flag.String("contentdir", "./markdown", "directory to load markdown content from")
	// TODO add images directory, because now it's located in the app /static directory
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
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.addRoutes(),
	}

	logger.Info("starting server", "addr", *addr)
	err = srv.ListenAndServe()
	logger.Error(err.Error())

	return nil
}
```

Here's a picture of the dev setup:
![image](/images/screenshot.png "Dev setup")


### asdasdasdasd

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris ut sagittis elit, sed consequat massa. Sed quis eleifend metus, ut gravida nulla. Aliquam erat volutpat. Fusce dictum laoreet orci ac vehicula. Nunc malesuada feugiat est, et pellentesque diam faucibus efficitur. Maecenas viverra, nisi vitae mattis rhoncus, massa risus pellentesque orci, in vestibulum est sem quis lorem. Nullam facilisis tincidunt nibh, ac cursus erat interdum ut. Vestibulum malesuada, nisi mattis fermentum volutpat, mauris erat fringilla risus, in laoreet magna augue non purus. Maecenas mattis erat eget ipsum fermentum maximus. Integer quis interdum ipsum. Integer porttitor sapien ullamcorper erat mattis porttitor. Praesent accumsan ex in metus tincidunt, id pharetra orci luctus.


[^1]: See https://pkg.go.dev/net/http
