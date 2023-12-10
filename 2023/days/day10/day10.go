package day10

import (
	"slices"

	"github.com/go-kit/log"
	util "github.com/thelande/adventofcode/common"
	"github.com/thelande/adventofcode/common/queue"
)

type Day10 struct{}

/**
 * Link the tree of cells from the string-based map.
 *
 * Returns the starting cell.
 */
func LinkTreeFromMap(cellMap [][]rune, cells [][]*Cell) *Cell {
	var start *Cell
	for i := range cells {
		for j := range cells[i] {
			switch cellMap[i][j] {
			case '|':
				cells[i][j].LinkNorth(cells, i, j)
				cells[i][j].LinkSouth(cells, i, j)
			case '-':
				cells[i][j].LinkEast(cells, i, j)
				cells[i][j].LinkWest(cells, i, j)
			case 'L':
				cells[i][j].LinkNorth(cells, i, j)
				cells[i][j].LinkEast(cells, i, j)
			case 'J':
				cells[i][j].LinkNorth(cells, i, j)
				cells[i][j].LinkWest(cells, i, j)
			case '7':
				cells[i][j].LinkSouth(cells, i, j)
				cells[i][j].LinkWest(cells, i, j)
			case 'F':
				cells[i][j].LinkSouth(cells, i, j)
				cells[i][j].LinkEast(cells, i, j)
			case 'S':
				start = cells[i][j]
			}
		}
	}

	// Figure out the neighbors of S

	// North
	if start.Y > 0 && cells[start.Y-1][start.X].South == start {
		start.LinkNorth(cells, start.Y, start.X)
	}
	// South
	if start.Y+1 < len(cells) && cells[start.Y+1][start.X].North == start {
		start.LinkSouth(cells, start.Y, start.X)
	}
	// East
	if start.X+1 < len(cells[0]) && cells[start.Y][start.X+1].West == start {
		start.LinkEast(cells, start.Y, start.X)
	}
	// West
	if start.X > 0 && cells[start.Y][start.X-1].East == start {
		start.LinkWest(cells, start.Y, start.X)
	}

	return start
}

func MakeTree(cellMap [][]rune) [][]*Cell {
	cells := make([][]*Cell, len(cellMap))
	for i, row := range cellMap {
		cells[i] = make([]*Cell, len(row))
		for j := range row {
			cells[i][j] = &Cell{X: j, Y: i}
		}
	}
	return cells
}

func BFS(start *Cell) int64 {
	var seen []*Cell
	var last *Cell

	seen = append(seen, start)
	queue := queue.NewQueue[*Cell]()
	queue.Push(start)

	for !queue.Empty() {
		val, err := queue.Pop()
		if err != nil {
			panic(err)
		}

		for _, c := range []*Cell{val.North, val.East, val.South, val.West} {
			if c != nil && !slices.Contains(seen, c) {
				c.Depth = val.Depth + 1
				queue.Push(c)
				seen = append(seen, c)
			}
		}

		last = val
	}

	return last.Depth
}

func (d Day10) Part1(filename string, logger log.Logger) int64 {
	var cellMap [][]rune

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		cellMap = append(cellMap, []rune(line))
		return nil
	})

	return BFS(LinkTreeFromMap(cellMap, MakeTree(cellMap)))
}

func (d Day10) Part2(filename string, logger log.Logger) int64 {
	var value int64
	var pipeMap [][]rune

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		pipeMap = append(pipeMap, []rune(line))
		return nil
	})

	return value
}
