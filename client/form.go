package client

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ClientForm struct {
	form *huh.Form
}

func NewClientForm() ClientForm {
	return ClientForm{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Client Name").
					Prompt("?").
					Key("name"),

				huh.NewInput().
					Title("Client Email").
					Prompt("?").
					Key("email"),
				huh.NewInput().
					Title("Client Rate").
					Prompt("?").
					Key("rate"),
			),
		),
	}
}

func (m ClientForm) Init() tea.Cmd {
	return m.form.Init()
}

func (m ClientForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m ClientForm) View() string {
	if m.form.State == huh.StateCompleted {
		name := m.form.GetString("name")
		email := m.form.GetString("email")
		rate := m.form.GetString("rate")
		return fmt.Sprintf("You selected: %s, %s, %s", name, email, rate)
	}
	return m.form.View()
}
