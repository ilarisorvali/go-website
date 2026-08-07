package content

import (
	"fmt"
	"html/template"
	"time"
)

type ContentCache struct {
	Items   map[string]*ContentItem
	Blog    []*ContentItem
	Kitchen []*ContentItem
}

// Metadata of a post
// Uses struct tags for marshaling purposes in markdown parsing
type FrontMatter struct {
	Title       string                `yaml:"Title"`
	Description string                `yaml:"Description"`
	Slug        string                `yaml:"Slug"`
	Draft       bool                  `yaml:"Draft"`
	Date        FrontMatterCustomDate `yaml:"Date"`
	Tags        []string              `yaml:"Tags"`
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
	Prev  *ContentItem
	Next  *ContentItem
	Items []*ContentItem
}

// Struct for custom date formatting when marshaling frontmatter
type FrontMatterCustomDate struct {
	time.Time
}

// This is a template for Go to use in parsing the date from markdown metadata
// See https://pkg.go.dev/time#Parse
// It has to be exactly 02.01.2006 here
// It's weird.
const dateTemplate = "02.01.2006"

// Custom date formatting for yaml marshaling with frontmatter
func (d *FrontMatterCustomDate) UnmarshalYAML(unmarshal func(any) error) error {
	var raw string
	if err := unmarshal(&raw); err != nil {
		return err
	}

	if raw == "" {
		return nil
	}

	t, err := time.Parse("02.01.2006", raw)
	if err != nil {
		return fmt.Errorf("invalid date %q (expected %s)", raw, dateTemplate)
	}

	d.Time = t
	return nil
}
