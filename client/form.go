package client

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ClientForm struct {
	form *huh.Form
}

func NewCForm() ClientForm {
	return ClientForm{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Key("val").Title("Client Name").Placeholder("John Doe"),
			),
		),
	}
}

func (c ClientForm) Init() tea.Cmd {
	return c.form.Init()
}

func (c ClientForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := c.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		c.form = f
	}
	return c, cmd
}

func (c ClientForm) View() string {
	if c.form.State == huh.StateCompleted {
		val := c.form.GetString("val")
		return fmt.Sprintf("you entered %s", val)
	}
	return c.form.View()
}
