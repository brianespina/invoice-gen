package main

import (
	"fmt"
	"os"

	"invoice-gen/client"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	clients []client.Client
	form    client.Form
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.form, cmd = m.form.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.clients = append(m.clients, client.Client{Name: m.form.FName})
			m.form = client.NewCForm()
			return m, cmd
		}
	}

	return m, cmd
}
func (m model) View() (s string) {
	for _, client := range m.clients {
		s += client.Name + "\n"
	}
	s += m.form.View()
	return
}
func initialModel() model {
	m := model{
		clients: []client.Client{
			{
				Name: "Brian Espina",
			},
		},
		form: client.NewCForm(),
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
