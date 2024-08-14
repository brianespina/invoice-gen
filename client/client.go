package client

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Client struct {
	Name string
}
type ClientForm struct {
	Form *huh.Form
}

func NewCForm() ClientForm {
	return ClientForm{
		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Key("val").Title("Client Name").Placeholder("John Doe"),
			),
		),
	}
}

func (f ClientForm) Init() tea.Cmd {
	return f.Form.Init()
}

func (f ClientForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, _ := f.Form.Update(msg)
	if ff, ok := form.(*huh.Form); ok {
		f.Form = ff
		return f, nil
	}
	return f, nil
}

func (f ClientForm) View() string {
	return f.Form.View()
}
