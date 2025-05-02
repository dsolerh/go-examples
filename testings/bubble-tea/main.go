package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// "Month", "Rent", "Utilities", "Budget", "Total"
	columns := []table.Column{
		{Title: "Month", Width: 6},
		{Title: "Rent", Width: 12},
		{Title: "Utilities", Width: 12},
		{Title: "Budget", Width: 12},
		{Title: "Total", Width: 12},
	}
	rows := []table.Row{
		{"April", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
		{"May", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
		{"June", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

// func initModel() *model {
// 	m := new(model)
// 	m.selectedRow = 0

// 	re := lipgloss.NewRenderer(os.Stdout)
// 	baseStyle := re.NewStyle().Padding(0, 1)
// 	headerStyle := baseStyle.Foreground(lipgloss.Color("#01BE85")).Bold(true)
// 	selectedStyle := baseStyle.Bold(true)

// 	headers := []string{"Month", "Rent", "Utilities", "Budget", "Total"}
// 	rows := [][]string{
// 		{"April", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
// 		{"May", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
// 		{"June", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN"},
// 	}

// 	m.maxRows = len(rows) - 1

// 	m.table = table.New().
// 		Headers(headers...).
// 		Rows(rows...).
// 		Border(lipgloss.NormalBorder()).
// 		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
// 		StyleFunc(func(row, col int) lipgloss.Style {
// 			// fmt.Printf("row: %v\n", row)
// 			if row == -1 {
// 				return headerStyle
// 			}

// 			if row == m.selectedRow {
// 				return selectedStyle
// 			}

// 			return baseStyle.Foreground(lipgloss.Color("252"))
// 		}).
// 		Border(lipgloss.ThickBorder())

// 	return m
// }

// type model struct {
// 	table       *table.Table
// 	selectedRow int
// 	maxRows     int
// }

// func (m *model) Init() tea.Cmd { return nil }

// func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.table = m.table.Width(msg.Width)
// 		m.table = m.table.Height(msg.Height)
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "up":
// 			// do something to move the focus up
// 			if m.selectedRow > 0 {
// 				m.selectedRow--
// 			}
// 		case "down":
// 			// do something to move the focus down
// 			if m.selectedRow < m.maxRows {
// 				m.selectedRow++
// 			}
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		case "enter":
// 			// enter a new one
// 			m.maxRows++
// 			m.table.Row("June", "1470.00 BGN", "474.45 BGN", "1500.00 BGN", "3,444.45 BGN")
// 		}
// 	}
// 	return m, cmd
// }

// func (m *model) View() string {
// 	return "\n" + m.table.String() + "\n"
// }
