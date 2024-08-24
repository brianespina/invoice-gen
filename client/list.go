package client

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	list   []Client
	cursor int
	db     *sql.DB
}

func NewList(db *sql.DB) *List {
	newList := &List{
		cursor: 0,
		db:     db,
	}
	newList.list = getClients(db)
	return newList
}

func getClients(db *sql.DB) []Client {
	var list []Client

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
		list = append(list, client)
	}
	return list
}

func (l List) Init() tea.Cmd {
	return nil
}
func (l List) View() string {
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
func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
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
			current := l.list[l.cursor]
			//delete client here
			l.list = deleteClient(l.db)
		}
	}
	return l, nil
}
