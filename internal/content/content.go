package models

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"gopkg.in/yaml.v2"
)

// Parses a post markdown file into metadata and HTML
// Returns metadata and HTML in a Post struct
func ParseMDToContent(file []byte) (*ContentItem, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			meta.Meta,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()

	err := markdown.Convert(file, &buf, parser.WithContext(context))
	if err != nil {
		panic(err)
	}

	//Get the yaml metadata part from MD file
	metaData := meta.Get(context)
	htmlTemplate := template.HTML(buf.String())

	//Extract metadata from MD file
	metaBuf, err := yaml.Marshal(metaData)
	if err != nil {
		return nil, err
	}

	//Unmarshal metadata into Frontmatter struct
	var meta FrontMatter
	if err := yaml.Unmarshal(metaBuf, &meta); err != nil {
		return nil, err
	}

	// This is a template for Go to use in parsing the date from markdown metadata
	// See https://pkg.go.dev/time#Parse
	// It has to be exactly 02.01.2006 here
	// It's weird.
	const dateTemplate = "02.1.2006"

	item := ContentItem{
		Meta:    meta,
		Content: htmlTemplate,
	}

	return &item, nil

}

// TODO add error handling
func LoadContentFiles(dirPath string) (map[string]*ContentItem, error) {
	//Init an empty map
	data := map[string]*ContentItem{}

	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		fullPath := filepath.Join(dirPath, file.Name())

		// Read	file
		md, err := os.ReadFile(fullPath)
		if err != nil {
			return data, err
		}

		// Parse file to ContentItem
		item, err := ParseMDToContent(md)
		if err != nil {
			return data, err
		}

		// TODO Only add to cache the requested kind

		data[item.Meta.Slug] = item

	}

	return data, nil

}

// TODO add error handling
func LoadSingleFile(filepath string) (TemplateData, error) {
	data := TemplateData{}

	file, _ := os.ReadFile(filepath)

	item, err := ParseMDToContent(file)
	if err != nil {
		return data, err
	}

	data.Item = item

	return data, nil
}

// tagSet makes a fake "set" out of slices because go doesn't have one implemented built in
// keys are the tags and values are empty structs (that take up zero space)
//
//	map[string]struct{}{
//		"major": {},
//		"minor": {},
//	}
func tagSet(tags []string) map[string]struct{} {
	set := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		set[tag] = struct{}{}
	}
	return set
}

// Filter a map[string]*ContentItem by ContentItem tags.
// Makes use of the tagset function for ~speed~
func FilterByTags(items map[string]*ContentItem, tags []string) map[string]*ContentItem {
	required := tagSet(tags)
	result := make(map[string]*ContentItem)

	for slug, item := range items {
		itemTags := tagSet(item.Meta.Tags)
		fmt.Println(itemTags)
		match := true
		for tag := range required {
			if _, ok := itemTags[tag]; !ok {
				match = false
				break
			}
		}

		if match {
			result[slug] = item
		}
	}

	return result
}
