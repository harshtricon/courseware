package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	content string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Courseware CLI\n\n%s\n\nPress 'q' to quit.", m.content)
}

func RunUI(content string) {
	p := tea.NewProgram(model{content: content})
	if err := p.Start(); err != nil {
		fmt.Println("Error running UI:", err)
	}
}
