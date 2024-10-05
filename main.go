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
		}
	}
	listModel, cmd := m.list.Update(msg)
	m.list = listModel
	return m, cmd
}
func (m model) View() string {
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
