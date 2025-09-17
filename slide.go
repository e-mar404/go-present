package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/glamour"
)

/* 
That will need to have a type like so-ish:
how would i want to make a new slide?
slideRenderer := newSlideRenderer(
	withGlamour(...termOptions)
	withDimensions(w, h)
)
slide := slidesRenderer.Render(curSlide)
p.viewport.SetContent(slidesRenderer.Render(curSlide))

this means that the struct would need to look like this:

slideRenderer {
	Height int
	Width int
	something that will hold the render function
}

then have a render function

the newSlideRenderer function however does need an argument with an interface of SlideRendrerOptions


type SlideRendererOptions func(*SlideRendrer) error
*/

type Renderer interface {
	Render(f fs.DirEntry) (string, error)
}

type SlideRenderer struct {
	Height int
	Width int
	renderer Renderer
}

func (sr SlideRenderer) Render(f fs.DirEntry) (string, error) {
	return sr.renderer.Render(f)
}

type SlideRendererOption func(*SlideRenderer) error 

type defaultSlideRenderer struct {}
func (dsr *defaultSlideRenderer) Render(f fs.DirEntry) (string, error) {
	if f.IsDir() {
		return "", fmt.Errorf("cannot make a slide form a directory")
	}

	// this is hard coded for now, should change later
	file, err := os.Open(filepath.Join("md", f.Name()))
	if err != nil {
		return "", fmt.Errorf("cannot open slide file: %v", err)
	}

	rawContent, _ := io.ReadAll(file)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)
	if err != nil {
		return "", err
	}

	renderedContent, err := renderer.Render(string(rawContent))
	if err != nil {
		return "", err
	}

	return renderedContent, nil
}

func NewSlideRenderer(options ...SlideRendererOption) (*SlideRenderer, error) {
	sr := &SlideRenderer{
		Width: 80,
		Height: 24,
		renderer: &defaultSlideRenderer{},
	}

	return sr, nil
}
