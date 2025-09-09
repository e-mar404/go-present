package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&slides{},
		), // will need to do my own extension to be able to identify slide syntax
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // probably do not need this?? or maybe the ability to identify a slide syntax goes in here instead of the extension
		),
		// will have to make my own rendered to use lipgloss
		goldmark.WithRendererOptions( 
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	source := []byte(`# title 1
some text`)

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
}
