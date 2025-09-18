package gopresent 

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
 	dir := os.Args[1]
	presentation, err := NewPresentation(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if _, err := tea.NewProgram(presentation, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Bummer, there was an error during the presentation:", err)
		os.Exit(1)
	}
}
