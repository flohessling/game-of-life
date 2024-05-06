package model

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithKeys("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q", "quit"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Help}, {k.Quit}}
}

type model struct {
	Grid          [][]bool
	Width, Height int
	keys          keyMap
	help          help.Model
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Duration(80)*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func InitialModel() model {
	return model{
		Width:  100,
		Height: 40,
		Grid:   makeGrid(100, 40),
		keys:   keys,
		help:   help.New(),
	}
}

func makeGrid(width, height int) [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	// f-pentomino
	// --xx--
	// -xx---
	// --x---
	grid[19][50] = true
	grid[19][51] = true
	grid[20][49] = true
	grid[20][50] = true
	grid[21][50] = true

	return grid
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}

	case tickMsg:
		return m.updateGrid(), tick()
	}

	return m, nil
}

func (m model) View() string {
	var view string

	frame := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#25A065"))

	for col := range m.Grid {
		for row := range m.Grid[col] {
			if m.Grid[col][row] {
				view += "█"
			} else {
				view += " "
			}
		}
		view += "\n"
	}

	frameView := frame.Render(view)

	title := lipgloss.NewStyle().
		Width(100).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("#04B575")).
		Render("conway's game of life")

	helpView := m.help.ShortHelpView([]key.Binding{keys.Help, keys.Quit})

	return "\n" + title + "\n" + frameView + "\n" + "\n" + helpView + "\n"
}

func (m model) updateGrid() model {
	next := makeGrid(m.Width, m.Height)
	for col := range m.Grid {
		for row := range m.Grid[col] {
			alive := m.Grid[col][row]
			neighbors := m.countNeighbors(row, col)

			if alive && (neighbors < 2 || neighbors > 3) {
				next[col][row] = false
			} else if !alive && neighbors == 3 {
				next[col][row] = true
			} else {
				next[col][row] = alive
			}
		}
	}

	return model{
		Width:  m.Width,
		Height: m.Height,
		Grid:   next,
	}
}

func (m model) countNeighbors(row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighborRow := row + i
			neighborCol := col + j
			if neighborRow >= 0 && neighborRow < m.Width && neighborCol >= 0 &&
				neighborCol < m.Height &&
				m.Grid[neighborCol][neighborRow] {
				count++
			}
		}
	}

	return count
}
