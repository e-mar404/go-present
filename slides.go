package main

import (
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

func (b *slideBlockParser) Trigger() []byte {
	return []byte(":::slide")
}

func (b *slideBlockParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	// this is for now to see what happens
	return nil, parser.NoChildren
}

// finish implementing parser.BlockParser

type slides struct {}

func (s *slides) Extend(md goldmark.Markdown) {
}

