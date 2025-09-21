package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type presentation struct {
	basePath      string
	curSlide      int
	slideFiles    []os.DirEntry
	SlideRenderer *SlideRenderer
	viewport      viewport.Model
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
			return p, p.NextSlide()

		case "ctrl+p":
			if p.curSlide > 0 {
				p.curSlide--
			}
			// content, _ := p.SlideRenderer.Render(p.slideFiles[p.curSlide])
			// p.viewport.SetContent(content)
			return p, nil

		default:
			var cmd tea.Cmd
			p.viewport, cmd = p.viewport.Update(msg)
			return p, cmd
		}
	case tea.WindowSizeMsg:
		p.viewport.Height = msg.Height
		p.viewport.Width = msg.Width

		return p, nil

	default:
		return p, nil
	}
}

func (p presentation) View() string {
	return p.viewport.View()
}

// TODO: maybe at some point i should think about making this function more extensible and calling it something like PresentationWithOptions that takes in any number of functions that take in a presentation and return it with its own config (should use an interface)
func NewPresentation(basePath string) (*presentation, error) {
	files, err := mdFilesFrom(basePath)
	if err != nil {
		return &presentation{}, err
	}

	renderer, err := NewSlideRenderer(
		WithGlamourDefault(),
	)
	if err != nil {
		return &presentation{}, err 
	}

	vp := viewport.New(78, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	if len(files) == 0 {
		vp.SetContent("No content to show")
		return &presentation{
			basePath:      basePath,
			slideFiles:    files,
			curSlide:      -1,
			viewport:      vp,
			SlideRenderer: renderer,
		}, nil
	}

	fullPath := filepath.Join(basePath, files[0].Name())
	fileContent, _ := os.ReadFile(fullPath)
	initContent, _ := renderer.Render(string(fileContent))

	vp.SetContent(initContent)

	return &presentation{
		basePath:      basePath,
		slideFiles:    files,
		curSlide:      0,
		viewport:      vp,
		SlideRenderer: renderer,
	}, nil
}

func (p *presentation) NextSlide() tea.Cmd {
	if p.curSlide == -1 || p.curSlide >= len(p.slideFiles)-1 {
		return nil
	}

	p.curSlide++

	path := filepath.Join(p.basePath, p.slideFiles[p.curSlide].Name())
	raw, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("err reading file: %v\n", err)
	}

	renderedStr, err := p.SlideRenderer.renderer.Render(string(raw))
	if err != nil {
		fmt.Printf("err rendering file: %v\n", err)
	}

	p.viewport.SetContent(renderedStr)

	return nil
}

func (p *presentation) PrevSlide() tea.Cmd {
	return nil
}
