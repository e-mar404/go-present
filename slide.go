package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

type slide struct {
	content string
}

func NewSlide(d fs.DirEntry) (slide, error) {
	if d.IsDir() {
		return slide{}, fmt.Errorf("cannot make a slide form a directory")
	}

	file, err := os.Open(filepath.Join("md", d.Name()))
	if err != nil {
		return slide{}, fmt.Errorf("cannot open slide file: %v", err)
	}

	rawContent, _ := io.ReadAll(file)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		return slide{}, err
	}
	
	renderedContent, err := renderer.Render(string(rawContent))
	if err != nil {
		return slide{}, err
	}

	return slide{
		content: renderedContent,
	}, nil
}

