package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	CtrlP key.Binding
	CtrlN key.Binding
	Quit key.Binding
}

func (km keyMap) ShortHelp() []key.Binding{
	return []key.Binding {km.CtrlP, km.CtrlN, km.Quit}
}

// TODO: where is this applied to, only saw the short help come up when running application
func (km keyMap) FullHelp() [][]key.Binding{
	return [][]key.Binding {
		{km.CtrlP, km.CtrlN, km.Quit},
	}
}

var keys = keyMap {
	CtrlP: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "previous slide"),
	),
	CtrlN: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "next slide"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}
