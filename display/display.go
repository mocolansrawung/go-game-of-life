package display

import (
	"fmt"
	"game-of-life-go/game"
	"os"
	"os/exec"
)

type Printer interface {
	Print(grid *game.Grid)
}

type ConsolePrinter struct{}

func NewConsolePrinter() Printer {
	return &ConsolePrinter{}
}

func (p *ConsolePrinter) Print(grid *game.Grid) {
	clearConsole()

	for _, row := range grid.Cells {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("⬛")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println()
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
