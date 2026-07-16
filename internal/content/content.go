package models

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	fm "github.com/adrg/frontmatter"
	figure "github.com/ilarisorvali/goldmark-picfig"
	"github.com/yuin/goldmark"
	hl "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

// Parses a post markdown file into metadata and HTML
// Returns metadata and HTML in a Post struct
func ParseMDToContent(file []byte) (*ContentItem, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			figure.PicFig,
			extension.GFM,
			extension.Footnote,
			hl.NewHighlighting(
				// neat built-in highlighting styles
				hl.WithStyle("dracula"),
			),
		),
	)

	//Get the yaml metadata part from MD file
	//Unmarshal metadata into Frontmatter struct
	var meta FrontMatter

	rest, err := fm.Parse(bytes.NewReader(file), &meta)
	if err != nil {
		return nil, err
	}

	// Parse the rest of the file into html content
	// First create memory buffer that implements io.Writer
	// Context is needed to hold state during conversion (footnotes, links, etc.)
	var buf bytes.Buffer
	context := parser.NewContext()

	// Convert rest to HTML into buf with parser context
	if err := markdown.Convert(rest, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}
	htmlTemplate := template.HTML(buf.String())

	// Final ContentItem
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

		//add to cache with slug as key
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
