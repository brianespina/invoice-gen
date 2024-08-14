package main

import (
	"fmt"
	"invoice-gen/client"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type model struct {
	clients       []client.Client
	form          client.ClientForm
	mode          string
	clientName    string
	clientContact string
}

func (m model) Init() tea.Cmd {
	if m.mode == "add" {
		return m.form.Form.Init()
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.form.Form.State == huh.StateCompleted {
		m.mode = "normal"
		m.clients = append(m.clients, client.Client{Name: m.form.Form.GetString("name")})
		m.form = client.NewCForm()
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			m.mode = "add"
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.mode {
	case "add":
		m.form.Form.Run()
		return ""
	case "normal":
		s := ""
		for _, client := range m.clients {
			s += client.Name + "\n"
		}
		return s
	}
	return "something"
}
func initialModel() model {
	m := model{
		clients: []client.Client{
			{
				Name: "Brian Espina",
			},
		},
		mode: "normal",
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
