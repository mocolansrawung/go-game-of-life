package game

type Grid [][]int

func NewGrid(rows, columns int) Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]int, columns)
	}
	return grid
}

func AddGliderGun(grid Grid, topRow, leftColumn int) {
	coords := [36][2]int{
		{0, 24}, {1, 22}, {1, 24}, {2, 12}, {2, 13}, {2, 20}, {2, 21}, {2, 34}, {2, 35},
		{3, 11}, {3, 15}, {3, 20}, {3, 21}, {3, 34}, {3, 35}, {4, 0}, {4, 1}, {4, 10},
		{4, 16}, {4, 20}, {4, 21}, {5, 0}, {5, 1}, {5, 10}, {5, 14}, {5, 16}, {5, 17},
		{5, 22}, {5, 24}, {6, 10}, {6, 16}, {6, 24}, {7, 11}, {7, 15}, {8, 12}, {8, 13},
	}

	for _, coord := range coords {
		grid[topRow+coord[0]][leftColumn+coord[1]] = 1
	}
}

func NextGeneration(grid Grid) Grid {
	rows := len(grid)
	columns := len(grid[0])
	newGrid := NewGrid(rows, columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			newGrid[i][j] = calculateNextState(grid, i, j)
		}
	}

	return newGrid
}
