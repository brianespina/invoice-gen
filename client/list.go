package client

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"invoice-gen/timelog"
)

const (
	clientListView view = iota
	timelistView
)

type List struct {
	list     []Client
	cursor   int
	timeList timelog.TimeList
	db       *sql.DB
	view     view
}

func NewList(db *sql.DB) *List {
	newList := &List{
		cursor: 0,
		db:     db,
	}
	newList.list = newList.getClients()
	return newList
}

func (l List) getClients() []Client {
	var list []Client

	//Query Clients form Db
	rows, err := l.db.Query("SELECT * FROM client")
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
		list = append(list, client)
	}
	return list
}
func (l List) deleteClient(id int) []Client {
	stmt, err := l.db.Prepare("DELETE FROM client WHERE client_id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return l.getClients()
}

func (l List) Init() tea.Cmd {
	return nil
}
func (l List) View() string {
	switch l.view {
	case timelistView:
		return l.timeList.View()
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
func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	current := l.list[l.cursor]
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
		case "ctrl+d":
			//delete client here
			l.list = l.deleteClient(current.id)
			l.cursor = 0
		case "enter":
			l.timeList = timelog.InitTimeList()
			timelog.FilterLogs(&l.timeList, current.id)
			l.view = timelistView
		case "esc":
			l.view = clientListView
		}

	}
	return l, nil
}
