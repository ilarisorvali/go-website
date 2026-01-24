package models

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Parses a post markdown file into metadata and HTML
// Returns metadata and HTML in a Post struct
func ParseMDToContent(file []byte) (*ContentItem, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
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

	//Extract metadata from MD file
	metaData := meta.Get(context)
	htmlTemplate := template.HTML(buf.String())

	//Extract concrete values from metadata
	//If value doesn't exist in metadata, init as zero value ""
	title, _ := metaData["Title"].(string)
	description, _ := metaData["Description"].(string)
	slug, _ := metaData["Slug"].(string)

	var ctype ContentType
	if t, ok := metaData["Type"].(string); ok && t == "post" {
		ctype = Post
	} else {
		ctype = Recipe
	}

	meta := ContentItemMeta{
		Title:       title,
		Description: description,
		Slug:        slug,
		Draft:       false,
		Date:        time.Now(),
		Type:        ctype,
	}

	post := ContentItem{
		Meta:    meta,
		Content: htmlTemplate,
	}

	return &post, nil

}

// TODO add error handling
func LoadContentFiles(dirPath string, kind ContentType) (map[string]*ContentItem, error) {
	//Init an empty Pages struct as go can't do it implicitly
	data := map[string]*ContentItem{}

	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		fullPath := filepath.Join(dirPath, file.Name())

		md, err := os.ReadFile(fullPath)
		if err != nil {
			return data, err
		}

		item, err := ParseMDToContent(md)
		if err != nil {
			return data, err
		}

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
