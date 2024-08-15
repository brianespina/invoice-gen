package main

import (
	"fmt"
	"invoice-gen/client"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type mode int

const (
	clientView mode = iota
	timeSheetView
	clientAddView
)

type Client struct {
	name, rate string
}

func (c Client) FilterValue() string {
	return c.name
}
func (c Client) Title() string {
	return c.name
}
func (c Client) Description() string {
	return c.rate
}

type model struct {
	list client.ClientList
	mode mode
}

func (m *model) InitList() {
	m.list = client.New()
	m.list.Populate()
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.InitList()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+t":
			m.mode = timeSheetView
		case "ctrl+a":
			m.mode = clientAddView
		case "ctrl+v":
			m.mode = clientView
		}
	}
	if m.mode == clientView {
		cl, cmd := m.list.Update(msg)
		m.list = cl
		return m, cmd
	}
	return m, nil
}
func (m model) View() string {
	switch m.mode {
	case clientView:
		return m.list.View()
	case timeSheetView:
		return "Timesheet View"
	case clientAddView:
		return "Add Clients"
	}
	return m.list.View()
}
func New() *model {
	return &model{}
}
func main() {
	p := tea.NewProgram(New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error has occurd: %v", err)
		os.Exit(1)
	}
}
