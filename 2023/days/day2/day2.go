package day2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	util "github.com/thelande/adventofcode/common"
)

type Reveal struct {
	Blue, Red, Green int64
}

type Day2 struct {
	Blue, Red, Green, Possible, TotalPower int64
}

func parseLine(line string) (int64, []Reveal) {
	var reveals []Reveal

	re := regexp.MustCompile(`^Game ([0-9]+): (.+)$`)
	matches := re.FindStringSubmatch(line)
	gameId, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}

	revealsStr := strings.Split(matches[2], ";")
	for _, reveal := range revealsStr {
		r := Reveal{}

		cubes := strings.Split(reveal, ",")
		for _, cube := range cubes {
			parts := strings.Split(strings.Trim(cube, " "), " ") // count and color
			count, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				panic(err)
			}

			switch parts[1] {
			case "blue":
				r.Blue = count
			case "red":
				r.Red = count
			case "green":
				r.Green = count
			}
		}

		reveals = append(reveals, r)
	}

	return gameId, reveals
}

func (d Day2) Part1(filename string, logger log.Logger) int64 {
	d.Red = 12
	d.Green = 13
	d.Blue = 14

	util.ReadPuzzleInput(filename, logger, func(line string) error {
		possible := true
		gameId, reveals := parseLine(line)
		for _, reveal := range reveals {
			if reveal.Blue > d.Blue || reveal.Green > d.Green || reveal.Red > d.Red {
				possible = false
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
	return d.TotalPower
}
