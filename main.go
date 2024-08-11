package main

import (
	"game-of-life-go/display"
	"game-of-life-go/game"
	"time"
)

func main() {
	grid := game.NewGrid(40, 40)
	grid.AddGliderGun(1, 1)

	printer := display.NewConsolePrinter()

	for {
		printer.Print(grid)
		grid.NextGeneration()
		time.Sleep(100 * time.Millisecond)
	}
}
