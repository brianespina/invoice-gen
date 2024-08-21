package client

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	list   []Client
	cursor int
}

func NewList(db *sql.DB) *List {
	newList := &List{
		cursor: 0,
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
		newList.list = append(newList.list, client)
	}

	return newList
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
		}
	}
	return l, nil
}
