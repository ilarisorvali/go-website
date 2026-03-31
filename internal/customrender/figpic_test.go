package customrender

import (
	"bytes"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func newMarkdown() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithRenderer(
			renderer.NewRenderer(
				renderer.WithNodeRenderers(
					util.Prioritized(NewImageFigureRenderer(), 100),
				),
			),
		),
	)
}

func TestImageFigureRenderer(t *testing.T) {
	md := newMarkdown()

	input := "![alt text](image.jpg \"caption text\")"

	expected := `<figure>
<picture>
<source srcset="image.jpg">
<img src="image.jpg" alt="alt text" loading="lazy" decoding="async">
</picture>
<figcaption>caption text</figcaption>
</figure>
`

	var buf bytes.Buffer
	err := md.Convert([]byte(input), &buf)
	if err != nil {
		t.Fatalf("convert failed: %v", err)
	}

	got := buf.String()

	if got != expected {
		t.Errorf("unexpected output:\n got \n%s\n want \n%s", got, expected)
	}
}
