package day8

import (
	"fmt"
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

// Returns the number of steps for the given starting location to reach
// the given ending location.
func stepsToEnd(start string, end *string, endRune *rune, directions string, nodes map[string]*Node) int64 {
	if end == nil && endRune == nil {
		panic("Neither end nor endRune provided.")
	}

	var value int64
	if end != nil {
		for curr := start; curr != *end; value++ {
			idx := int(value) % len(directions)
			dir := directions[idx]
			if dir == 'L' {
				curr = nodes[curr].LeftLabel
			} else {
				curr = nodes[curr].RightLabel
			}
		}
	} else {
		for curr := start; rune(curr[2]) != *endRune; value++ {
			idx := int(value) % len(directions)
			dir := directions[idx]
			if dir == 'L' {
				curr = nodes[curr].LeftLabel
			} else {
				curr = nodes[curr].RightLabel
			}
		}
	}

	return value
}

func (d Day8) Part1(filename string, logger log.Logger) int64 {
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
	end := "ZZZ"
	return stepsToEnd("AAA", &end, nil, directions, nodes)
}

func (d Day8) Part2(filename string, logger log.Logger) int64 {
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

	// Find and track the paths
	var paths []string
	for k := range nodes {
		if k[2] == 'A' {
			paths = append(paths, k)
		}
	}

	level.Debug(logger).Log("directions", directions)
	level.Debug(logger).Log("len(paths)", len(paths), "paths", fmt.Sprintf("%s", paths))

	return value
}
