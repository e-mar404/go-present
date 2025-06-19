package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
  items []string
  paginator paginator.Model
}

func newModel() model {
  var items []string
  for i := range(100) {
    text := fmt.Sprintf("Item %d", i)
    items = append(items, text)
  }
  
  p := paginator.New()
  p.Type = paginator.Dots
  p.PerPage = 10
  p.ActiveDot = lipgloss.
    NewStyle().
    Foreground(lipgloss.AdaptiveColor{
      Light: "235",
      Dark: "252",
    }).Render("•")
  p.InactiveDot= lipgloss.
    NewStyle().
    Foreground(lipgloss.AdaptiveColor{
      Light: "250",
      Dark: "238",
    }).Render("•")
  p.SetTotalPages(len(items))

  return model {
    items: items,
    paginator: p,
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "q", "esc", "ctrl+c":
      return m, tea.Quit
    }
  }
  m.paginator, cmd = m.paginator.Update(msg)
  return m, cmd 
}
 
func (m model) View() string {
  var buf strings.Builder
  buf.WriteString("\n Paginator\n\n")
  start, end := m.paginator.GetSliceBounds(len(m.items))
  for _, item := range m.items[start:end] {
    buf.WriteString(" - " + item + "\n\n")
  }

  buf.WriteString(" " + m.paginator.View())
  buf.WriteString("\n\n h/l <-/-> page | q: quit\n")

  return buf.String()
}

func main() {
  p := tea.NewProgram(newModel(), tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Could not run the tea program: %v\n", err)
    os.Exit(1)
  }
}
