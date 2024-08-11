package game

type Grid struct {
	Cells [][]int
}

func NewGrid(rows, columns int) *Grid {
	grid := &Grid{
		Cells: make([][]int, rows),
	}
	for i := range grid.Cells {
		grid.Cells[i] = make([]int, columns)
	}
	return grid
}

func (g *Grid) AddGliderGun(topRow, leftColumn int) {
	coords := [36][2]int{
		{0, 24}, {1, 22}, {1, 24}, {2, 12}, {2, 13}, {2, 20}, {2, 21}, {2, 34}, {2, 35},
		{3, 11}, {3, 15}, {3, 20}, {3, 21}, {3, 34}, {3, 35}, {4, 0}, {4, 1}, {4, 10},
		{4, 16}, {4, 20}, {4, 21}, {5, 0}, {5, 1}, {5, 10}, {5, 14}, {5, 16}, {5, 17},
		{5, 22}, {5, 24}, {6, 10}, {6, 16}, {6, 24}, {7, 11}, {7, 15}, {8, 12}, {8, 13},
	}

	for _, coord := range coords {
		g.Cells[topRow+coord[0]][leftColumn+coord[1]] = 1
	}
}

func (g *Grid) NextGeneration() {
	rows := len(g.Cells)
	columns := len(g.Cells[0])
	newCells := make([][]int, rows)
	for i := range newCells {
		newCells[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			newCells[i][j] = calculateNextState(g.Cells, i, j)
		}
	}

	g.Cells = newCells
}
