package timelog

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Timelog struct {
	Name        string
	Description string
	Date        time.Time
	Log         float64
	Client      int
}

type TimeList struct {
	list  []Timelog
	table table.Model
	db    *sql.DB
}

func InitTimeList(db *sql.DB) TimeList {

	list := TimeList{
		list: []Timelog{},
		db:   db,
	}
	res, err := db.Query("SELECT * FROM timelog")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var name, description, date string
		var log float64
		var client, id int

		if err := res.Scan(&id, &name, &description, &log, &date, &client); err != nil {
			panic(err)
		}
		timeRow := Timelog{Name: name, Description: description, Log: log, Client: client}
		list.list = append(list.list, timeRow)
	}

	return list

}
func FilterLogs(list *TimeList, clientId int) {
	var filtered []Timelog
	for _, log := range list.list {
		if clientId == log.Client {
			filtered = append(filtered, log)
		}
	}
	list.list = filtered
	list.InitTable()
}
func (t TimeList) Init() tea.Cmd {
	return nil
}
func (t *TimeList) InitTable() {
	columns := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Time", Width: 10},
		{Title: "Client", Width: 10},
	}
	rows := []table.Row{}
	//get logs here

	for _, log := range t.list {
		logS := []string{
			log.Name,
			fmt.Sprintf("%.2f", log.Log),
			fmt.Sprintf("%d", log.Client),
		}
		rows = append(rows, logS)
	}
	table := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.table = table
}
func (t TimeList) View() string {
	return t.table.View()
}
func (t TimeList) Update(msg tea.Msg) (TimeList, tea.Cmd) {
	t.table, _ = t.table.Update(msg)
	return t, nil
}
