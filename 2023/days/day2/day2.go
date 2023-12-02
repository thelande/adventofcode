package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Reveal struct {
	Blue, Red, Green int64
}

type Day2 struct {
	Blue, Red, Green, Possible int64
}

func (d Day2) Part1(filename string, logger log.Logger) int64 {
	d.Blue = 14
	d.Red = 12
	d.Green = 13

	re := regexp.MustCompile(`^Game ([0-9]+): (.+)$`)
	util.ReadPuzzleInput(filename, logger, func(line string) error {
		possible := true
		matches := re.FindStringSubmatch(line)
		gameId, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			panic(err)
		}

		reveals := strings.Split(matches[2], ";")
		for _, reveal := range reveals {
			cubes := strings.Split(reveal, ",")
			for _, cube := range cubes {
				parts := strings.Split(strings.Trim(cube, " "), " ") // count and color
				count, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}

				level.Debug(logger).Log("color", parts[1], "count", count)

				switch parts[1] {
				case "blue":
					if count > d.Blue {
						possible = false
					}
				case "red":
					if count > d.Red {
						possible = false
					}
				case "green":
					if count > d.Green {
						possible = false
					}
				}
			}
		}

		if possible {
			d.Possible += gameId
		}

		return nil
	})

	return d.Possible
}

func (d Day2) Part2(filename string, logger log.Logger) int64 {
	return 0
}
