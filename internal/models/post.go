package models

import (
	"html/template"
)

type ContentData struct {
	Title       string   //title of post
	Description string   //description of post
	Tags        []string //tags of post
	Slug        string   //url slug of post
	Order       int
	Content     template.HTML
}

func NewPost() (ContentData, error) {
	tags := []string{"testing", "go", "climbing"}

	return ContentData{
		Title:       "This is a test post.",
		Description: "Test post description lors lars",
		Tags:        tags,
		Slug:        "test-post",
		Order:       1,
		Content:     template.HTML("<p>asdasdasdasdsdasd</>"),
	}, nil
}
