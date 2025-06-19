package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(newPresentation(), tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Could not start presentation: %v\n", err)
    os.Exit(1)
  }
}
