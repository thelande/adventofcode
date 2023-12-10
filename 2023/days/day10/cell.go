package day10

type Cell struct {
	X, Y                     int
	North, South, East, West *Cell
	Depth                    int64
}

func (c *Cell) LinkNorth(cells [][]*Cell, i, j int) {
	if i == 0 {
		return
	}
	c.North = cells[i-1][j]
}

func (c *Cell) LinkSouth(cells [][]*Cell, i, j int) {
	if i+1 == len(cells) {
		return
	}
	c.South = cells[i+1][j]
}

func (c *Cell) LinkEast(cells [][]*Cell, i, j int) {
	if j+1 == len(cells[i]) {
		return
	}
	c.East = cells[i][j+1]
}

func (c *Cell) LinkWest(cells [][]*Cell, i, j int) {
	if j == 0 {
		return
	}
	c.West = cells[i][j-1]
}
