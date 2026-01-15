package models

import (
	"html/template"
	"strconv"
)

type contentItem struct {
	Title       string   //title of post
	Description string   //description of post
	Tags        []string //tags of post
	Slug        string   //url slug of post
	Order       int
	Content     template.HTML
}

type Pages struct {
	Title string
	Items []contentItem
}

func NewContentItem() contentItem {
	tags := []string{"testing", "go", "climbing"}

	return contentItem{
		Title:       "This is a test post.",
		Description: "Test post description lors lars",
		Tags:        tags,
		Slug:        "test-post",
		Order:       1,
		Content:     template.HTML("<p>asdasdasdasdsdasd</>"),
	}
}

func NewManyContentItems() Pages {
	var ret Pages
	ret.Title = "asdasdsadad"
	for i := 0; i < 8; i++ {
		ret.Items = append(ret.Items, NewContentItem())
		ret.Items[i].Slug = "test-slug" + strconv.Itoa(i)
		ret.Items[i].Title = "This is a test post " + strconv.Itoa(i)
	}
	return ret
}
