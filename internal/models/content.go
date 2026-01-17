package models

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Parses a post markdown file into metadata and HTML
// Returns metadata and HTML in a Post struct
func ParseMDToPost(file []byte) (ContentItem, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
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

	post := ContentItem{
		Title:       title,
		Description: description,
		Slug:        slug,
		Draft:       false,
		Date:        time.Now(),
		Kind:        ctype,
		Content:     htmlTemplate,
	}

	return post, nil

}

func LoadMarkdownPosts(dirPath string) (Pages, error) {
	//Init an empty Pages struct as go can't do it implicitly
	pages := Pages{
		Posts: []ContentItem{},
	}

	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		fullPath := filepath.Join(dirPath, file.Name())

		md, err := os.ReadFile(fullPath)
		if err != nil {
			return pages, err
		}

		post, err := ParseMDToPost(md)
		if err != nil {
			return pages, err
		}

		pages.Posts = append(pages.Posts, post)
	}

	return pages, nil

}
