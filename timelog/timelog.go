package timelog

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type Timelog struct {
	Name   string
	Date   time.Time
	Log    float32
	Client int
}

type TimeList struct {
	list  []Timelog
	table table.Model
}

func InitTimeList() TimeList {
	list := TimeList{
		list: []Timelog{
			{Name: "MNY Site Build", Log: 16.5, Client: 1},
			{Name: "Nova Website Updates", Log: .5, Client: 1},
			{Name: "Product Page Updates", Log: 4, Client: 2},
		},
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
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if t.table.Focused() {
				t.table.Blur()
			} else {
				t.table.Focus()
			}
		case "enter":
			return t, tea.Batch(
				tea.Printf("Let's go to %s!", t.table.SelectedRow()[1]),
			)
		}
	}
	t.table, cmd = t.table.Update(msg)
	return t, cmd
}
