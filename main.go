package main

import (
	"fmt"
	"os"

	"invoice-gen/client"

	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type model struct {
	clients []client.Client
	form    client.Form
	mode    string
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.mode == "add" {
		m.form, cmd = m.form.Update(msg)
		log.Println("press")
		if m.form.IsComplete() {
			m.clients = append(m.clients, client.Client{Name: m.form.FName})
			m.mode = "normal"
		}
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.mode = "add"
			//m.clients = append(m.clients, client.Client{Name: m.form.FName})
			return m, cmd
		case "esc":
			m.mode = "normal"
		}
	}
	return m, cmd
}
func (m model) View() (s string) {
	switch m.mode {
	case "add":
		s = m.form.View()
	case "normal":
		for _, client := range m.clients {
			s += client.Name + "\n"
		}
	}
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
		mode: "normal",
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
