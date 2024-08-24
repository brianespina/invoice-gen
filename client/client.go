package client

import (
	"database/sql"

	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type view int

const (
	normal view = iota
	add
)

type Client struct {
	name  string
	email string
	rate  float32
	id    int
}

type ClientList struct {
	list *List
	view view
	db   *sql.DB
	form *huh.Form
}

func New(db *sql.DB) ClientList {
	//Initialize new Client list
	clientListInstance := ClientList{
		db:   db,
		list: NewList(db),
	}

	return clientListInstance
}
func (l *ClientList) resetForm() {
	form := huh.NewForm(
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
	)
	l.form = form
}

func (l *ClientList) addClient(name, email, rate string) {
	_, err := l.db.Exec("INSERT INTO client VALUES(NULL,?,?,?)", name, email, rate)
	if err != nil {
		panic(err)
	}
}
func (l ClientList) Init() tea.Cmd {
	return nil
}
func (l ClientList) View() string {
	switch l.view {
	case add:
		l.form.Init()
		if l.form.State == huh.StateCompleted {
			name := l.form.GetString("name")
			email := l.form.GetString("email")
			rate := l.form.GetString("rate")
			//add client here
			l.addClient(name, email, rate)
			return fmt.Sprintf("Client has been added: \n\n%s\nemail: %s\nrate: %s", name, email, rate)
		}
		return l.form.View()
	case normal:
		fallthrough
	default:
		return l.list.View()

	}
}

func (l ClientList) Update(msg tea.Msg) (ClientList, tea.Cmd) {
	switch l.view {
	case normal:
		listModel, _ := l.list.Update(msg)
		l.list = &listModel
	case add:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+n":
				l.view = normal
			}
		}
		if l.form.State == huh.StateCompleted {
			//refresh list
			l.list = NewList(l.db)
			l.view = normal
		}
		form, cmd := l.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			l.form = f
		}
		return l, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+n":
			//reset form here
			l.resetForm()
			l.view = add
		}
	}

	return l, nil
}
