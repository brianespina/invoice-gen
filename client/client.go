package client

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type view int

const (
	normal view = iota
	details
)

type Client struct {
	name  string
	email string
	rate  float32
}
type ClientList struct {
	list   []Client
	cursor int
	view   view
}

func (l *ClientList) Populate() {
	l.list = []Client{
		{name: "White Sheep Digital", rate: 20, email: "sample@email.com"},
		{name: "Marc Frojentein", rate: 8.5, email: "sample@email.com"},
		{name: "Hicaliber", rate: 14, email: "sample@email.com"},
	}
}
func New() ClientList {
	return ClientList{
		cursor: 0,
	}
}
func (l ClientList) Init() tea.Cmd {
	return nil
}

func (l ClientList) View() string {
	switch l.view {
	case details:
		var s string
		client := l.list[l.cursor]
		s += fmt.Sprintf("%s\n", client.name)
		s += fmt.Sprintf("%s\n", client.email)
		s += fmt.Sprintf("%.2f\n", client.rate)
		return s
	case normal:
		fallthrough
	default:
		var s string
		for i, client := range l.list {
			cursor := ""
			if l.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %d %s\n", cursor, i, client.name)
		}
		return s
	}
}

func (l ClientList) Update(msg tea.Msg) (ClientList, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return l, tea.Quit
		case "j":
			if l.cursor == len(l.list)-1 {
				l.cursor = 0
			} else {
				l.cursor++
			}
		case "k":
			if l.cursor == 0 {
				l.cursor = len(l.list) - 1
			} else {
				l.cursor--
			}
		case "d":
			l.view = details
		case "l":
			l.view = normal
		}
	}
	return l, nil
}
