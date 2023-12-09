package day8

import (
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day8 struct{}

type Node struct {
	Label, LeftLabel, RightLabel string
}

func NewNodeFromLine(line string) *Node {
	node := &Node{}
	parts := strings.Split(line, "=")
	node.Label = strings.Trim(parts[0], " ")

	parts = strings.Split(parts[1], ",")
	node.LeftLabel = strings.Trim(parts[0], "( ")
	node.RightLabel = strings.Trim(parts[1], ") ")

	return node
}

func (d Day8) Part1(filename string, logger log.Logger) int64 {
	var value int64
	var directions string
	nodes := make(map[string]*Node)

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		if lineno == 0 {
			directions = line
		} else if lineno == 1 {
			// skip
		} else {
			node := NewNodeFromLine(line)
			nodes[node.Label] = node
		}
		return nil
	})

	level.Debug(logger).Log("directions", directions)
	for curr := "AAA"; curr != "ZZZ"; value++ {
		idx := int(value) % len(directions)
		dir := directions[idx]
		if dir == 'L' {
			curr = nodes[curr].LeftLabel
		} else {
			curr = nodes[curr].RightLabel
		}
	}

	return value
}

func (d Day8) Part2(filename string, logger log.Logger) int64 {
	var value int64
	// var directions string
	// var complete bool = false
	// nodes := make(map[string]*Node)

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		// if lineno == 0 {
		// 	directions = line
		// } else if lineno == 1 {
		// 	// skip
		// } else {
		// 	node := NewNodeFromLine(line)
		// 	nodes[node.Label] = node
		// }
		return nil
	})

	// Loop until we get all paths to finish at the same time.
	// for value = 0; !complete; value++ {
	// 	idx := int(value) % len(directions)
	// }

	return value
}
