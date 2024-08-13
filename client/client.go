package client

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Client struct {
	Name string
}
type Form struct {
	FName string
	Form  *huh.Form
}

func NewCForm() Form {
	return Form{
		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Key("val").Title("Client Name").Placeholder("John Doe"),
			),
		),
	}
}

func (f Form) Init() tea.Cmd {
	return f.Form.Init()
}

func (f Form) Update(msg tea.Msg) (Form, tea.Cmd) {
	form, cmd := f.Form.Update(msg)
	if ff, ok := form.(*huh.Form); ok {
		f.Form = ff
	}
	return f, cmd
}

func (f Form) View() string {
	return f.Form.View()
}
