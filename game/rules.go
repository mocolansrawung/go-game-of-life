package game

func calculateNextState(grid Grid, row, col int) int {
	liveNeighbors := countLiveNeighbors(grid, row, col)
	currentState := grid[row][col]

	if currentState == 1 && (liveNeighbors == 2 || liveNeighbors == 3) {
		return 1
	} else if currentState == 0 && liveNeighbors == 3 {
		return 1
	} else {
		return 0
	}
}

func countLiveNeighbors(grid Grid, row, col int) int {
	rows := len(grid)
	columns := len(grid[0])
	liveNeighbors := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			neighborRow := row + i
			neighborCol := col + j

			if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < columns {
				liveNeighbors += grid[neighborRow][neighborCol]
			}
		}
	}

	return liveNeighbors
}
