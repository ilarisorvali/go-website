---
Title: Building a website in go, a novice approach
Description: Test Description
Slug: test-post
Date: 25.01.2025
Tags: [post, major]
---

asdasdasdasdasdasd
asdasdasdasdasdasdasdasdasdasdasdasdasdasdasd
asdasdasdasdasdassad

This site is made with [Go](https://go.dev)

## Ayy lmao

steps

- asdasdasdsdas
- russian bias

```go
func newTemplateCache() (map[string]*template.Template, error) {
	//Init an empty map to act as the template cache
	cache := map[string]*template.Template{}

	//Get slice of page filepath strings
	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.html')
		name := filepath.Base(page)

		// Parse the base template
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// ParseGlob() on current template set to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// ParseFiles() on current template set to add the current page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}
```


### asdasdasdasd

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris ut sagittis elit, sed consequat massa. Sed quis eleifend metus, ut gravida nulla. Aliquam erat volutpat. Fusce dictum laoreet orci ac vehicula. Nunc malesuada feugiat est, et pellentesque diam faucibus efficitur. Maecenas viverra, nisi vitae mattis rhoncus, massa risus pellentesque orci, in vestibulum est sem quis lorem. Nullam facilisis tincidunt nibh, ac cursus erat interdum ut. Vestibulum malesuada, nisi mattis fermentum volutpat, mauris erat fringilla risus, in laoreet magna augue non purus. Maecenas mattis erat eget ipsum fermentum maximus. Integer quis interdum ipsum. Integer porttitor sapien ullamcorper erat mattis porttitor. Praesent accumsan ex in metus tincidunt, id pharetra orci luctus.