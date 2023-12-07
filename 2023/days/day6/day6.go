package day6

import (
	"fmt"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day6 struct{}

func (d Day6) Part1(filename string, logger log.Logger) int64 {
	var times, distances []int64
	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		parts := strings.Split(line, ":")
		if lineno == 0 {
			times = util.NumListToSlice(parts[1])
		} else {
			distances = util.NumListToSlice(parts[1])
		}
		return nil
	})

	level.Debug(logger).Log("times", fmt.Sprintf("%d", times), "distances", fmt.Sprintf("%d", distances))

	var spreads []int64
	for i, x := range times {
		var minZ, maxZ int64
		y := distances[i]
		for minZ = 1; minZ+(y/minZ) >= x; minZ++ {
		}
		for maxZ = x; maxZ+(y/maxZ) >= x; maxZ-- {
		}
		spreads = append(spreads, maxZ-minZ+1)
	}

	level.Debug(logger).Log("spreads", fmt.Sprintf("%d", spreads))

	return util.MultSlice(spreads)
}

func (d Day6) Part2(filename string, logger log.Logger) int64 {
	return d.Part1(filename, logger)
}
