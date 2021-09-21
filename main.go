package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type View int

const (
	SetupView View = iota
	InitialView
	HomeView
	NewEntryView
)

type model struct {
	view View
}

func initialModel() model {
	return model{
		view: InitialView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "yeet"
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
