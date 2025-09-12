package main

import (
	"io/fs"
	"path/filepath"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type presentation struct {
	curSlide	int
	slides   []slide
	viewport viewport.Model
}

var _ tea.Model = presentation{}

func (p presentation) Init() tea.Cmd {
	return nil
}

func (p presentation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
			case "q", "ctrl+c", "esc":
				return p, tea.Quit

			case "ctrl+n":
				if p.curSlide < len(p.slides) - 1 {
					p.curSlide++
				}
				p.viewport.SetContent(p.slides[p.curSlide].content)
				return p, nil

			case "ctrl+p":
				if p.curSlide > 0 {
					p.curSlide--
				}
				p.viewport.SetContent(p.slides[p.curSlide].content)
				return p, nil

			default:
				var cmd tea.Cmd
				p.viewport, cmd = p.viewport.Update(msg)
				return p, cmd 
		}
	default:
		return p, nil
	}
}

func (p presentation) View() string {
	return p.viewport.View()
}

func NewPresentation(path string) (*presentation, error) {
	var slides []slide
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		slide, err := NewSlide(d)
		if err != nil {
			return err
		}
		slides = append(slides, slide)
		return nil
	})
	if err != nil {
		return &presentation{}, err
	}
	
	vp := viewport.New(78, 20)
	vp.SetContent(slides[0].content)

	return &presentation{
		slides: slides, 
		curSlide: 0,
		viewport: vp,
	}, nil
}
