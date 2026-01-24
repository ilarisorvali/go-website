---
Title: Longer code block
Description: Test Description2
Slug: test-post3
Kind: post
Date: 12.01.2025
---


# Huutis mint

This is another post loers laers :DDDDDDDDDD
asdasdasdasdasdasd
asdasdasdasdasdasdasdasdasdasdasdasdasdasdasd
asdasdasdasdasdassad

That's some text with a footnote.[^1]

This site is made with [Go](https://go.dev)

## Here is a code block lol

```go
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data, _ := models.LoadSingleFile(homePath)

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {

	posts := app.postCache

	data := models.TemplateData{
		Items: posts,
	}

	app.render(w, r, http.StatusOK, "posts.html", data)
}

func (app *application) viewPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post := app.postCache[slug]
	app.logger.Debug(slug)

	data := models.TemplateData{
		Item: post,
	}

	app.render(w, r, http.StatusOK, "post.html", data)
}
```

## asdasdasdasd

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris ut sagittis elit, sed consequat massa. Sed quis eleifend metus, ut gravida nulla. Aliquam erat volutpat. Fusce dictum laoreet orci ac vehicula. Nunc malesuada feugiat est, et pellentesque diam faucibus efficitur. Maecenas viverra, nisi vitae mattis rhoncus, massa risus pellentesque orci, in vestibulum est sem quis lorem. Nullam facilisis tincidunt nibh, ac cursus erat interdum ut. Vestibulum malesuada, nisi mattis fermentum volutpat, mauris erat fringilla risus, in laoreet magna augue non purus. Maecenas mattis erat eget ipsum fermentum maximus. Integer quis interdum ipsum. Integer porttitor sapien ullamcorper erat mattis porttitor. Praesent accumsan ex in metus tincidunt, id pharetra orci luctus.

[^1]: And that's the footnote.
