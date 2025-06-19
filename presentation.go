package main

import (
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type presentation struct {
  pages []byte
  curPage int
  paginator paginator.Model
}

func newPresentation() presentation {
  file, _ := os.Open("test.md")
  pages, _ := io.ReadAll(file)

  _, height, _ := term.GetSize(os.Stdout.Fd())

  p := paginator.New()
  p.Type = paginator.Dots
  p.PerPage = height 
  p.ActiveDot = lipgloss.
    NewStyle().
    Foreground(lipgloss.AdaptiveColor{
      Light: "235",
      Dark: "252",
    }).Render("•")
  p.InactiveDot = lipgloss.
    NewStyle().
    Foreground(lipgloss.AdaptiveColor{
      Light: "250",
      Dark: "238",
    }).Render("•")
  p.SetTotalPages(len(pages))

  return presentation {
    pages: pages,
    curPage: 0,
    paginator: p,
  }
}

func (p presentation) Init() tea.Cmd {
  return nil
}

func (p presentation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "q", "esc", "ctrl+c":
      return p, tea.Quit
    case "l":
      p.curPage++
    case "h":
      p.curPage--
    }
  case tea.WindowSizeMsg:
  }
  p.paginator, cmd = p.paginator.Update(msg)
  return p, cmd 
}
 
func (p presentation) View() string {
  var buf strings.Builder
  start, end := p.paginator.GetSliceBounds(len(p.pages))
  for _, item := range p.pages[start:end] {
    buf.WriteByte(item)
  }

  buf.WriteString("\n\n " + p.paginator.View())
  buf.WriteString("\n\n h/l <-/-> page | q: quit\n")

  return buf.String()
}
