package day10

import (
	"slices"
)

var (
	// Pipes that connect from the north
	northPipes = []rune{'|', 'F', '7', 'S'}
	// Pipes that connect from the south
	southPipes = []rune{'|', 'L', 'J', 'S'}
	// Pipes that connect from the east
	eastPipes = []rune{'-', '7', 'J', 'S'}
	// Pipes that connect from the west
	westPipes = []rune{'-', 'L', 'F', 'S'}
)

type Cell struct {
	X, Y                     int
	North, South, East, West *Cell
	Depth                    int64
	Path                     []*Cell
	Symbol                   rune
}

func (c *Cell) IsGround() bool {
	return c.Symbol == '.'
}

// Returns true if c and x are linked pipes.
func (c *Cell) IsLinked(x *Cell) bool {
	// False if either are ground.
	if c.IsGround() || x.IsGround() {
		return false
	}

	if (c.North == x && slices.Contains(southPipes, c.Symbol) && slices.Contains(northPipes, x.Symbol)) ||
		(c.South == x && slices.Contains(northPipes, c.Symbol) && slices.Contains(southPipes, x.Symbol)) ||
		(c.East == x && slices.Contains(westPipes, c.Symbol) && slices.Contains(eastPipes, x.Symbol)) ||
		(c.West == x && slices.Contains(eastPipes, c.Symbol) && slices.Contains(westPipes, x.Symbol)) {
		return true
	}

	return false
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
