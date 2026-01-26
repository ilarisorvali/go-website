package models

import (
	"html/template"
	"time"
)

// metadata of a post
// uses struct tags for marshaling purposes in markdown parsing
type FrontMatter struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Slug        string    `yaml:"slug"`
	Draft       bool      `yaml:"tags"`
	Date        time.Time `yaml:"date"`
	Tags        []string  `yaml:"tags"`
}

// single item of content for website,
// md parsed to html with md frontmatter in FrontMatter struct
type ContentItem struct {
	Meta    FrontMatter   // metadata
	Content template.HTML // parsed main markdown content

}

// Final datatype sent to pages,
// contains a list of items and or one specific item to show
type TemplateData struct {
	Item  *ContentItem
	Items map[string]*ContentItem
}
