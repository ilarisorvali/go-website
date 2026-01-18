package models

import (
	"html/template"
	"time"
)

type ContentType string

const (
	Post   ContentType = "post"
	Recipe ContentType = "recipe"
)

type ContentItemMeta struct {
	Title       string      // meta: title of ContentItem
	Description string      // meta: description of ContentItem
	Slug        string      // meta: url slug of ContentItem
	Draft       bool        // is the ContentItem a draft, ie. is it visible
	Date        time.Time   // timestamp of ContentItem, not a real 'hard' timestamp
	Type        ContentType // "type" of ContentItem Post/Recipe
}

type ContentItem struct {
	Meta    ContentItemMeta // metadata
	Content template.HTML   // parsed main markdown content

}

type TemplateData struct {
	Item  *ContentItem
	Items map[string]*ContentItem
}
