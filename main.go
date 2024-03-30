package main

import (
	"fmt"
  "log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
  fmt.Println(logfilePath)
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

  log.Println("Starting program")

	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct{}

func (m model) Init() tea.Cmd {
	// Return a command to perform when the program starts.
	// For now, return `nil`, meaning "do nothing."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
    case tea.KeyMsg:
      log.Println("Key pressed:", msg.String())
      if msg.String() == "q" {
        log.Println("Quitting program")
        return m, tea.Quit
      }
    }

  return m, nil
}

func (m model) View() string {
	// Return a string that will be rendered in the terminal.
	return "Hello, Bubble Tea!\n\nPress 'q' to quit.";
}
