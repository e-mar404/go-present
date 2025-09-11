package main

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type slides struct{}
type slideBlockRenderer struct{}

func (e slides) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(slideBlockParser{}, 500),
		),
	)

	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(slideBlockRenderer{}, 500),
	),
	)
}

var _ parser.BlockParser = slideBlockParser{}

type slideBlockParser struct{}

func (s slideBlockParser) Trigger() []byte {
	return []byte{':'}
}

// Open implements parser.BlockParser.
func (s slideBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, _ := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 || line[pos] != ':' {
		return nil, parser.NoChildren
	}
	slideChar := line[pos]

	fmt.Println(slideChar)

	return nil, parser.Continue
}

// Continue implements parser.BlockParser.
func (s slideBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	panic("unimplemented")
}

// Close implements parser.BlockParser.
func (s slideBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	panic("unimplemented")
}

func (s slideBlockParser) CanInterruptParagraph() bool {
	return true
}

func (s slideBlockParser) CanAcceptIndentedLine() bool {
	return false
}
