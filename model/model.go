package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Grid          [][]bool
	Width, Height int
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Duration(80)*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func InitialModel() model {
	return model{
		Width:  400,
		Height: 60,
		Grid:   makeGrid(400, 60),
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
	grid[12][40] = true
	grid[12][41] = true
	grid[13][39] = true
	grid[13][40] = true
	grid[14][40] = true

	// evolution to empty grid
	// -xxx-
	// -x-x-
	// -x-x-
	// -----
	// -x-x-
	// -x-x-
	// -xxx-
	// grid[12][39] = true
	// grid[12][40] = true
	// grid[12][41] = true
	// grid[13][39] = true
	// grid[13][41] = true
	// grid[14][39] = true
	// grid[14][41] = true
	// grid[16][39] = true
	// grid[16][41] = true
	// grid[17][39] = true
	// grid[17][41] = true
	// grid[18][39] = true
	// grid[18][40] = true
	// grid[18][41] = true

	return grid
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		return m.updateGrid(), tick()
	}
	return m, nil
}

func (m model) View() string {
	var view string

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

	return view
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
