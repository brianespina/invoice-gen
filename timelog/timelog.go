package timelog

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type view int

const (
	normal view = iota
	add
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
	form  *huh.Form
	view  view
}

func InitTimeList(db *sql.DB) TimeList {

	list := TimeList{
		list: []Timelog{},
		db:   db,
	}

	list.resetForm()
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
func (t *TimeList) resetForm() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Prompt("?").
				Key("name"),

			huh.NewInput().
				Title("Description").
				Prompt("?").
				Key("description"),
			huh.NewInput().
				Title("Log").
				Prompt("?").
				Key("log"),
			huh.NewInput().
				Title("Client").
				Prompt("?").
				Key("client"),
		),
	)
	t.form = form
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
func (t TimeList) addTime(name, description, log, client string) {
	_, err := t.db.Exec("INSERT INTO timelog (name, description, log, client)", name, description, log, client)
	if err != nil {
		panic(err)
	}
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
	switch t.view {
	case add:
		t.form.Init()

		if t.form.State == huh.StateCompleted {
			name := t.form.GetString("name")
			description := t.form.GetString("description")
			log := t.form.GetString("log")
			client := t.form.GetString("client")

			//add clienn here
			t.addTime(name, description, log, client)
			return fmt.Sprintf("%s, %s, %s, %s", name, description, log, client)
		}
		return t.form.View()
	case normal:
		fallthrough
	default:
		return t.table.View()
	}
}
func (t TimeList) Update(msg tea.Msg) (TimeList, tea.Cmd) {
	switch t.view {
	case add:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+n":
				t.view = normal
			}
		}
		if t.form.State == huh.StateCompleted {
			//refresh list
			t.view = normal
		}

		form, cmd := t.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			t.form = f
		}
		return t, cmd
	case normal:
		fallthrough
	default:
		t.table, _ = t.table.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+n":
			//reset form here
			t.resetForm()
			t.view = add
		}
	}

	return t, nil
}
