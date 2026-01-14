package server

import (
	"html/template"
)

type Recipe struct {
	Title                   string
	Slug                    string
	Content                 template.HTML
	Description             string
	Order                   int
	Tags                    []string
	MetaDescription         string
	MetaPropertyTitle       string
	MetaPropertyDescription string
	MetaOgURL               string
}

type Post struct {
	Title       string //title of post
	Description string //description of post
	Content     template.HTML
	Tags        []string //tags of post
	Slug        string   //url slug of post
}

type Category struct {
	Name  string
	Pages []Post
	Order int
}

type PostData struct {
	Posts []Category
}

func NewPost() Post {
	tags := []string{"testing", "go", "climbing"}

	return Post{
		Title:       "This is a test post.",
		Description: "Test post description lors lars",
		Content:     template.HTML("<h2> Hello world </h2>"),
		Tags:        tags,
		Slug:        "test-post",
	}
}
