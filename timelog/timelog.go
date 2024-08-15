package timelog

import (
	"fmt"
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
	list []Timelog
}

func InitTimeList() TimeList {
	return TimeList{
		list: []Timelog{
			{Name: "MNY Site Build", Log: 16.5, Client: 0},
			{Name: "Nova Website Updates", Log: .5, Client: 1},
			{Name: "Product Page Updates", Log: 4, Client: 2},
		},
	}
}
func FilterLogs(list *TimeList, clientId int) {
	var filtered []Timelog
	for _, log := range list.list {
		if clientId == log.Client {
			filtered = append(filtered, log)
		}
	}
	list.list = filtered
}
func (t TimeList) Init() tea.Cmd {
	return nil
}
func (t TimeList) View() string {
	var s string
	for _, time := range t.list {
		s += fmt.Sprintf("%s\n", time.Name)
		s += fmt.Sprintf("%.2f\n\n", time.Log)
	}
	return s
}
func (t TimeList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}
