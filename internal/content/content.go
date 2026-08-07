package content

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"sync"

	figure "github.com/mangoumbrella/goldmark-figure"

	fm "github.com/adrg/frontmatter"
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
			figure.Figure,
			extension.GFM,
			extension.Footnote,
			hl.NewHighlighting(
				hl.WithStyle("nord"),
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

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return data, err
	}

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

func LoadContentFilesParallel(dirPath string) (map[string]*ContentItem, error) {
	// amount of worker threads
	const WORKERS_AMOUNT = 8
	data := make(map[string]*ContentItem)

	// read all of the md file paths from md directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return data, err
	}

	// Jobs are files that need to be processed.
	mdjobs := make(chan os.DirEntry)

	// processed markdown files are sent to a channel, because
	// writing to a map in parallel is messy
	processedMd := make(chan *ContentItem)

	// waitgroup for workers
	var wg sync.WaitGroup

	// add files to job channel
	go func() {
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".md" {
				mdjobs <- file
			}
		}
		// close the channel
		close(mdjobs)
	}()

	// start the worker pool of WORKERS_AMOUNT size
	// each worker gets md files (paths) from the mdjobs channel
	for range WORKERS_AMOUNT {
		wg.Go(func() {
			for file := range mdjobs {
				fullPath := filepath.Join(dirPath, file.Name())
				//fmt.Println(fullPath)

				md, err := os.ReadFile(fullPath)
				if err != nil {
					continue
				}

				item, err := ParseMDToContent(md)
				if err != nil {
					continue
				}

				processedMd <- item
			}
		})
	}

	// run wg.Wait() inside a goroutine so the main thread doesn't get blocked
	go func() {
		wg.Wait()
		close(processedMd)
	}()

	// get results to a map
	for item := range processedMd {
		data[item.Meta.Slug] = item
	}

	return data, nil
}
