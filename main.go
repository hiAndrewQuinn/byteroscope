package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	// Return a command to perform when the program starts.
	// For now, return `nil`, meaning "do nothing."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This function is called when messages are received.
	// Return the updated model and an optional command.
	return m, tea.Quit
}

func (m model) View() string {
	// Return a string that will be rendered in the terminal.
	return "Hello, Bubble Tea!"
}
