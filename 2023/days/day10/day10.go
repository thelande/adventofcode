package day10

import (
	"fmt"
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

func PrintMap(cells [][]*Cell) {
	fmt.Printf("  ")
	for x := range cells[0] {
		fmt.Printf("%d", x)
	}
	fmt.Println()
	for i := range cells {
		fmt.Printf("%d ", i)
		for j := range cells[i] {
			fmt.Printf("%c", cells[i][j].Symbol)
		}
		fmt.Println()
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

func BFS2(start *Cell) bool {
	var seen []*Cell

	seen = append(seen, start)
	queue := queue.NewQueue[*Cell]()
	queue.Push(start)

	for !queue.Empty() {
		val, err := queue.Pop()
		if err != nil {
			panic(err)
		}

		if val.IsExit() {
			fmt.Println("Path")
			for _, n := range val.Path {
				fmt.Printf("(%d,%d) [%c]\n", n.X, n.Y, n.Symbol)
			}
			return true
		}

		for _, c := range []*Cell{val.North, val.East, val.South, val.West} {
			if c != nil && !slices.Contains(seen, c) {
				push := false
				// Ground connects to ground
				if val.IsGround() && c.IsGround() {
					push = true
				} else if val.IsGround() {
					// Ground can enter pipes that open to the ground.
					if val.South == c && slices.Contains([]rune{'|', '7', 'F'}, c.Symbol) {
						push = true
					} else if val.North == c && slices.Contains([]rune{'|', 'L', 'J'}, c.Symbol) {
						// Look south
						push = true
					} else if val.East == c && slices.Contains([]rune{'-', 'F', 'L'}, c.Symbol) {
						// Look west
						push = true
					} else if val.West == c && slices.Contains([]rune{'-', '7', 'J'}, c.Symbol) {
						// Look east
						push = true
					}
				} else if !val.IsGround() && (val.IsLinked(c) || val.ExitsTo(c)) {
					push = true
				}

				if push {
					c.Path = append(c.Path, val.Path...)
					c.Path = append(c.Path, val)
					queue.Push(c)
					seen = append(seen, c)
				}
			}
		}
	}

	return false
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
	var cellMap [][]rune

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		cellMap = append(cellMap, []rune(line))
		return nil
	})

	cells := MakeTree(cellMap, true)
	var groundCells []*Cell
	for i := range cells {
		for j := range cells[i] {
			if cells[i][j].IsGround() {
				groundCells = append(groundCells, cells[i][j])
			}
		}
	}

	LinkTreeFromMap(cellMap, cells, true)

	PrintMap(cells)

	BFS2(cells[6][2])

	for _, cell := range groundCells {
		if !BFS2(cell) {
			value++
			fmt.Println(cell.Path)
		}
	}

	return value
}
