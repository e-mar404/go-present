package gopresent 

import (
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
	Render(string) (string, error)
}

type SlideRenderer struct {
	Height int
	Width int
	renderer Renderer
}

func (sr SlideRenderer) Render(in string) (string, error) {
	return sr.renderer.Render(in)
}

type SlideRendererOption func(*SlideRenderer) error 

func WithGlamourDefault () SlideRendererOption {
	return func(sr *SlideRenderer) error {
		glamourRenderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
		)
		sr.renderer = glamourRenderer
		return err
	}
}

func NewSlideRenderer(options ...SlideRendererOption) (*SlideRenderer, error) {
	sr := &SlideRenderer{
		Width: 80,
		Height: 24,
	}

	for _, opt := range options {
		if err := opt(sr); err != nil {
			return &SlideRenderer{}, nil
		}
	}

	return sr, nil
}
