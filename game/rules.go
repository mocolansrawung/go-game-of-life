package game

type cellState int

const (
	alive cellState = 1
	dead  cellState = 0
)

func calculateNextState(grid [][]int, row, col int) int {
	liveNeighbors := countLiveNeighbors(grid, row, col)
	currentState := grid[row][col]

	if currentState == int(alive) && (liveNeighbors == 2 || liveNeighbors == 3) {
		return int(alive)
	} else if currentState == 0 && liveNeighbors == 3 {
		return int(alive)
	} else {
		return int(dead)
	}
}

func countLiveNeighbors(grid [][]int, row, col int) int {
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

			// TODO: wrap the validation in another function
			if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < columns {
				liveNeighbors += grid[neighborRow][neighborCol]
			}
		}
	}

	return liveNeighbors
}
