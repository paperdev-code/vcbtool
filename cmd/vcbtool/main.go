package main

import (
	"fmt"
	"log"
	"time"
	"vcb"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	vcb.Greet()

	p := tea.NewProgram(model(5))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type model int

func (m model) Init() tea.Cmd {
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tickMsg:
		m -= 1
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("This program exists in %d seconds. Press any key to exit.", m)
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
