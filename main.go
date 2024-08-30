package main

import (
	"database/sql"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
	"invoice-gen/client"
	"invoice-gen/timelog"
	"os"
)

type mode int

const (
	clientView mode = iota
	timeSheetView
	clientAddView
)

type model struct {
	list     client.ClientList
	timeList timelog.TimeList
	mode     mode
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	if m.mode == timeSheetView {
		cl, cmd := m.timeList.Update(msg)
		m.timeList = cl
		return m, cmd
	}
	return m, nil
}
func (m model) View() string {
	switch m.mode {
	case clientView:
		return m.list.View()
	case timeSheetView:
		return m.timeList.View()
	case clientAddView:
		return "Add Clients"
	}
	return m.list.View()
}
func New(db *sql.DB) *model {
	m := &model{
		list:     client.New(db),
		timeList: timelog.InitTimeList(db),
	}
	return m
}
func main() {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := tea.NewProgram(New(db))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error has occurd: %v", err)
		os.Exit(1)
	}
}
