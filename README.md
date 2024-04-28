# Game of Life

This is a simple implementation of Conway's Game of Life running in the terminal.

It is written in Go and using bubbletea for the terminal rendering.

## TODOs

- Check if `tea.WithAltScreen` screen size can be queried for the grid size.
- Add helper menu to display `q` and `ctrl-c` to quit.
- Make the starting pattern configurable (currently it's hardcoded in the initial model).
