package client

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"invoice-gen/timelog"
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
	id    int
}

type ClientList struct {
	list   []Client
	cursor int
	view   view
	db     *sql.DB
}

func (l *ClientList) Db(db *sql.DB) {
	l.db = db
}
func New(db *sql.DB) ClientList {
	//Initialize new Client list
	list := ClientList{
		cursor: 0,
		db:     db,
	}

	//Query Clients form Db
	//Populate the list
	rows, err := db.Query("SELECT * FROM client")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

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
	return nil
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
		case "enter":
			l.view = details
		case "esc":
			l.view = normal
		}
	}
	return l, nil
}
