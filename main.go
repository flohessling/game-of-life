package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/flohessling/game-of-life/model"
)

func main() {
	p := tea.NewProgram(model.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("this did not go well: %v", err)
		os.Exit(1)
	}
}
