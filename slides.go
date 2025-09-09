package main

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type slideNode struct {
	ast.BaseBlock
}

type slideBlockParser struct {}

var defaultSlideBlockParser = &slideBlockParser{}

func NewSlideBlockParser() parser.BlockParser {
	return defaultSlideBlockParser
}

// the parse.BlockParser implementation is for now since i do not know yet how I should be modifying the parser i just want to see the md block "run" somehow
func (b *slideBlockParser) Trigger() []byte {
	return []byte{'x'}
}

func (b *slideBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	fmt.Println("slidesBlockParser called Open")
	withNewline := text.FindClosureOptions{
		Newline: true,
	}
	segments, found := text.Reader.FindClosure(reader, 'x', 'x', withNewline)
	fmt.Println(segments, found)
	return nil, parser.NoChildren
}

func (b *slideBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Continue
}

func (b *slideBlockParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
}

func (b *slideBlockParser) CanInterruptParagraph() bool {
	return true
}

func (b *slideBlockParser) CanAcceptIndentedLine() bool {
	return false 
}

type slides struct {}

func (s *slides) Extend(md goldmark.Markdown) {
}

