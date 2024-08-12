package main

import (
	"fmt"
	"os"

	"invoice-gen/client"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	//"time"
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
	if m.form.Form.State == huh.StateCompleted {
		m.clients = append(m.clients, client.Client{Name: m.form.FName})
		m.form = client.NewCForm()
	}
	return m, cmd
}
func (m model) View() string {
	if m.form.Form.State == huh.StateCompleted {
	}
	return m.form.View()
}
func initialModel() model {
	m := model{
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
