package display

import (
	"fmt"
	"game-of-life-go/game"
	"os"
	"os/exec"
)

func PrintGrid(grid game.Grid) {
	clearConsole()

	for _, row := range grid {
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
