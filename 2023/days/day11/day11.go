package day11

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day11 struct{}

type Cell struct {
	X, Y   int
	Galaxy bool
}

type Pair struct {
	A, B *Cell
}

func (c *Cell) Print() {
	fmt.Printf("(%d,%d)\n", c.X, c.Y)
}

func (p *Pair) ToString() string {
	return fmt.Sprintf("(%d,%d) --> (%d,%d)", p.A.X, p.A.Y, p.B.X, p.B.Y)
}

func (p *Pair) Distance() int64 {
	return int64(math.Abs(float64(p.A.X-p.B.X)) + math.Abs(float64(p.A.Y-p.B.Y)))
}

func (d Day11) Part1(filename string, logger log.Logger) int64 {
	var value int64

	row := 0
	width := 0
	var emptyCols []int
	var cells []*Cell
	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		if lineno == 0 {
			width = len(line)
			// Construct emptyCols and then remove columns for each galaxy found per row.
			for i := 0; i < width; i++ {
				emptyCols = append(emptyCols, i)
			}
		}

		if !strings.Contains(line, "#") {
			// No galaxies on this line, double increment row
			row++
		} else {
			for i := 0; i < width; i++ {
				if line[i] == '#' {
					cells = append(cells, &Cell{X: i, Y: row, Galaxy: true})
					colIdx := slices.Index(emptyCols, i)
					if colIdx >= 0 {
						emptyCols = slices.Delete(emptyCols, colIdx, colIdx+1)
					}
				}
			}
		}

		row++
		return nil
	})

	// Process the empty columns in reverse to avoid increasing cells beyond
	// an empty column.
	slices.Reverse(emptyCols)
	for _, v := range emptyCols {
		for i := range cells {
			if cells[i].X > v {
				cells[i].X++
			}
		}
	}

	var pairs []*Pair
	for i := range cells {
		for j := i + 1; j < len(cells); j++ {
			pairs = append(pairs, &Pair{A: cells[i], B: cells[j]})
		}
	}

	level.Debug(logger).Log("len(pairs)", len(pairs))
	for i := range pairs {
		dist := pairs[i].Distance()
		value += dist
	}

	return value
}

func (d Day11) Part2(filename string, logger log.Logger) int64 {
	var value int64

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		return nil
	})

	return value
}
