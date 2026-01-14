package server

import (
	"html/template"
)

type PageData struct {
	Title   string
	Posts   []Post
	Post    Post
	Recipes []Recipe
}

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
	Number      int      //ordering number of post
	Slug        string   //url slug of post
	Visible     bool     //is post visible on the website boolean
}

func NewRecipe() Recipe {
	var ret Recipe

	return ret
}
