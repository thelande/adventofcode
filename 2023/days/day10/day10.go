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
 */
func LinkTreeFromMap(cellMap [][]rune, cells [][]*Cell, partTwo bool) {
	for i := range cells {
		for j := range cells[i] {
			cell := cells[i][j]
			if i > 0 {
				cell.LinkNorth(cells, i, j)
			}
			if i+1 < len(cells) {
				cell.LinkSouth(cells, i, j)
			}
			if j > 0 {
				cell.LinkWest(cells, i, j)
			}
			if j+1 < len(cells[i]) {
				cell.LinkEast(cells, i, j)
			}
		}
	}
}

func FindStart(cellMap [][]rune, cells [][]*Cell) *Cell {
	var start *Cell
	for i := range cells {
		for j := range cells[i] {
			if cellMap[i][j] == 'S' {
				start = cells[i][j]
				break
			}
		}
	}

	if start == nil {
		panic("Never found start")
	}

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

func MakeTree(cellMap [][]rune, partTwo bool) [][]*Cell {
	cells := make([][]*Cell, len(cellMap))
	for i, row := range cellMap {
		cells[i] = make([]*Cell, len(row))
		for j := range row {
			cells[i][j] = &Cell{X: j, Y: i, Symbol: cellMap[i][j]}
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
			if c != nil && c.IsLinked(val) && !slices.Contains(seen, c) {
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

	cells := MakeTree(cellMap, false)
	LinkTreeFromMap(cellMap, cells, false)
	return BFS(FindStart(cellMap, cells))
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
