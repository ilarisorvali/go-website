package models

import (
	"fmt"
	"strconv"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type SrcSetRenderer struct {
	Sizes []int
}

func NewSrcSetRenderer() renderer.NodeRenderer {
	return &SrcSetRenderer{
		Sizes: []int{400, 800, 1200},
	}
}

func (r *SrcSetRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindImage, r.renderImage)
}

func (r *SrcSetRenderer) renderImage(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {

	if !entering {
		return ast.WalkContinue, nil
	}

	img := node.(*ast.Image)

	src := string(img.Destination)

	srcset := r.buildSrcset(src)

	_, err := fmt.Fprintf(w, `<img srcset="%s" sizes="(max-width: 800px) 100vw, 800px" loading="lazy">`, srcset)

	return ast.WalkSkipChildren, err
}

func (r *SrcSetRenderer) buildSrcset(src string) string {
	out := ""

	for i, s := range r.Sizes {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%s?w=%d %sw", src, s, strconv.Itoa(s))
	}

	return out
}
