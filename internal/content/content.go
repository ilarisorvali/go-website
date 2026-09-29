package content

import (
	"bytes"
	temp "html/template"
	"os"
	"path/filepath"
	"sync"

	highlighting "github.com/yuin/goldmark-highlighting/v3"
	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
	"go.yaml.in/yaml/v4"
)

// Parses a post markdown file into metadata and HTML
// Returns metadata and HTML in a Post struct
func ParseMDToContent(source []byte) (*ContentItem, error) {
	var buf bytes.Buffer
	var frontMatter FrontMatter

	doc := parser.New(
		parser.WithExtensions(
			meta.Parser,
			highlighting.Parser,
		),
	).Parse(source)

	mData := doc.(*ast.Document).Metadata()

	data, err := yaml.Marshal(mData)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &frontMatter); err != nil {
		panic(err)
	}

	// HTML renderer
	r := html.New(
		html.WithExtensions(
			highlighting.NewHTMLRenderer(
				highlighting.WithStyle("nord"),
			),
		),
	)

	if err := r.Render(&buf, source, doc); err != nil {
			panic(err)
	}

	htmlTemplate := temp.HTML(buf.String())

	println(htmlTemplate)

	item := ContentItem{
		Meta:	frontMatter,
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
			// range over the channel to get jobs for the worker
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
