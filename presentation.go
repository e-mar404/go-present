package gopresent

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
			// i just want to return p.NextSlide() instead of all the things below

			// if p.curSlide < len(p.slideFiles)-1 {
			// 	p.curSlide++
			// }
			// renderedContent, err := p.SlideRenderer.Render("hello")
			// if err != nil {
			// 	fmt.Println(err)
			// }

			// p.viewport.SetContent(renderedContent)
			return p, nil

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
func NewPresentation(basePath string) (presentation, error) {
	fmt.Println("reading from ", basePath)
	entries, _ := os.ReadDir(basePath)
	var files []os.DirEntry
	for _, entry := range entries {
		fmt.Println(entry)
		if !entry.IsDir() {
			files = append(files, entry)
		}
	}

	fmt.Printf("dir files: %v\n", files)

	renderer, err := NewSlideRenderer(
		WithGlamourDefault(),
	)
	if err != nil {
		return presentation{}, nil
	}

	vp := viewport.New(78, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	if len(files) == 0 {
		vp.SetContent("No content to show")
		return presentation{
			basePath:   basePath,
			slideFiles: files,
			curSlide:   -1,
			viewport:   vp,
		}, nil
	}

	fullPath := filepath.Join(basePath, files[0].Name())
	fileContent, _ := os.ReadFile(fullPath)
	initContent, _ := renderer.Render(string(fileContent))

	vp.SetContent(initContent)


	return presentation{
		basePath:   basePath,
		slideFiles: files,
		curSlide:   0,
		viewport:   vp,
	}, nil
}

func (p presentation) NextSlide() (presentation, tea.Cmd) {
	if p.curSlide == -1 {
		return p, nil
	}

	p.curSlide++
	path := filepath.Join(p.basePath, p.slideFiles[p.curSlide].Name())
	fmt.Printf("using file %v, with curslide num: %v\n", path, p.curSlide)
	raw, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("err reading file: %v\n", err)
	}

	renderedStr, err := p.SlideRenderer.renderer.Render(string(raw))
	if err != nil {
		fmt.Printf("err rendering file: %v\n", err)
	}
	p.viewport.SetContent(renderedStr)
	
	return p, nil
}
