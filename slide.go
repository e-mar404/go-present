package gopresent

import (
	"github.com/charmbracelet/glamour"
)

type Renderer interface {
	Render(string) (string, error)
}

type SlideRenderer struct {
	Height   int
	Width    int
	renderer Renderer
}

func (sr *SlideRenderer) Render(in string) (string, error) {
	return sr.renderer.Render(in)
}

type SlideRendererOption func(*SlideRenderer) error

func WithGlamourDefault() SlideRendererOption {
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
		Width:  80,
		Height: 24,
	}

	for _, opt := range options {
		if err := opt(sr); err != nil {
			return &SlideRenderer{}, nil
		}
	}

	return sr, nil
}
