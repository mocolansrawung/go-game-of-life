package main

import (
	"game-of-life-go/display"
	"game-of-life-go/game"

	"time"
)

func main() {
	grid := game.NewGrid(40, 40)
	game.AddGliderGun(grid, 1, 1)

	for {
		display.PrintGrid(grid)
		grid = game.NextGeneration(grid)
		time.Sleep(100 * time.Millisecond)
	}
}
