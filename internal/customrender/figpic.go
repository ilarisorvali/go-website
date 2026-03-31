package customrender

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// Renderer struct where render methods are implemented
// Acts as a renderer for nodes of type Image to following:
//
// <figure>
//
//	<picture>
//	  <source srcset=xyz... type=hjk...>
//	  <img src=xyz... alt=hjk... loading="lazy" decoding="async"
//	</picture>
//	<figcaption>asdf</figcaption>
//
// <figure>

type ImageFigureRenderer struct{}

// Return a new ImageFigureRenderer
func NewImageFigureRenderer() renderer.NodeRenderer {
	return &ImageFigureRenderer{}
}

// Method to register custom rendering functions for ast nodes of kind Image
// With priority overrides the default Image node HTML renderer
func (r *ImageFigureRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindImage, r.RenderImage)
}

func (r *ImageFigureRenderer) RenderImage(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}
