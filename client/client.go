package client

import (
	"database/sql"
	"fmt"
	"invoice-gen/timelog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type view int

const (
	normal view = iota
	details
	add
)

type Client struct {
	name  string
	email string
	rate  float32
	id    int
}

type ClientList struct {
	list   []Client
	cursor int
	view   view
	db     *sql.DB
	form   *huh.Form
}

func (l *ClientList) Db(db *sql.DB) {
	l.db = db
}
func New(db *sql.DB) ClientList {
	//Initialize new Client list
	list := ClientList{
		cursor: 0,
		db:     db,
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

	//Query Clients form Db
	rows, err := db.Query("SELECT * FROM client")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//Populate the list
	for rows.Next() {
		client := Client{}
		if err := rows.Scan(&client.id, &client.name, &client.email, &client.rate); err != nil {
			panic(err)
		}
		list.list = append(list.list, client)
	}

	return list
}
func (l ClientList) Init() tea.Cmd {
	return l.form.Init()
}

func (l ClientList) View() string {
	switch l.view {
	case details:
		var s string
		client := l.list[l.cursor]
		s += fmt.Sprintf("%s\n", client.name)
		s += fmt.Sprintf("%s\n", client.email)
		s += fmt.Sprintf("%.1f\n\n", client.rate)
		t := timelog.InitTimeList()
		timelog.FilterLogs(&t, l.cursor)
		s += t.View()
		return s
	case add:
		if l.form.State == huh.StateCompleted {
			name := l.form.GetString("name")
			email := l.form.GetString("email")
			rate := l.form.GetString("rate")
			return fmt.Sprintf("You selected: %s, %s, %s", name, email, rate)
		}
		return l.form.View()
	case normal:
		fallthrough
	default:
		var s string
		for i, client := range l.list {
			cursor := " "
			if l.cursor == i {
				cursor = "|"
			}
			s += fmt.Sprintf("%s %s\n", cursor, client.name)
		}
		return s
	}
}

func (l ClientList) Update(msg tea.Msg) (ClientList, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
		case "enter":
			l.view = details
		case "ctrl+n":
			l.view = add
		case "esc":
			l.view = normal
		}
	}

	form, cmd := l.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		l.form = f
	}

	return l, cmd
}
