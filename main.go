package main

import (
	"fmt"
	"os"

	"invoice-gen/client"

	tea "github.com/charmbracelet/bubbletea"
	//"time"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	state    string
	form     client.ClientForm
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}
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
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "ctrl+n":
			m.state = "client_add"
		case "esc":
			m.state = "normal"
			//default:
			//	log.Println(msg.String())
		}

	}
	if m.state == "client_add" {
		return m.form.Update(msg)
	}
	return m, nil
}
func (m model) View() string {
	s := ""
	switch m.state {
	case "normal":
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}
	case "client_add":
		s += m.form.View()
	}

	s += "\nPress q to quit.\n"
	return s
}
func initialModel() model {
	m := model{
		choices:  []string{"Buy Gata", "Buy Pork", "Buy Sili"},
		selected: make(map[int]struct{}),
		state:    "normal",
		form:     client.NewCForm(),
	}
	return m
}
func main() {
	//p := tea.NewProgram(client.NewCForm())
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error has occurd: %v", err)
		os.Exit(1)
	}
}
