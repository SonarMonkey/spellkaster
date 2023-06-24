package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

// define a model to store the application state
type model struct {
	// items to choose from
	choices []string
	// position of the cursor
	cursor int
}

// define the initial state of model
func initialModel() model {
	return model{
		choices: []string{"Update system", "Search packages", "Install packages"},
	}
}

// define an init function to do nothing for now
func (m model) Init() tea.Cmd {
	return nil
}

// define an update function to activate on keypress
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
		switch msg.String() {
			case "ctrl+c", "q":
			return m, tea.Quit

			case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

// define a view function to display the model
func (m model) View() string {
	s := "What would you like to do?\n\n"

	// iterate over the choices
	for i, choice := range m.choices {
		// no cursor
		cursor := " "

		// cursor!!!
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

// put all of it together
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err !=nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}