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

type Day2 struct{}

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
	var red, green, blue, possibleCount int64
	red = 12
	green = 13
	blue = 14

	util.ReadPuzzleInput(filename, logger, func(line string) error {
		gameId, reveals := parseLine(line)
		for _, reveal := range reveals {
			if reveal.Blue > blue || reveal.Green > green || reveal.Red > red {
				return nil
			}
		}

		possibleCount += gameId

		return nil
	})

	return possibleCount
}

func (d Day2) Part2(filename string, logger log.Logger) int64 {
	var totalPower int64

	util.ReadPuzzleInput(filename, logger, func(line string) error {
		_, reveals := parseLine(line)
		var maxRed, maxBlue, maxGreen int64

		for _, reveal := range reveals {
			maxRed = max(maxRed, reveal.Red)
			maxBlue = max(maxBlue, reveal.Blue)
			maxGreen = max(maxGreen, reveal.Green)
		}

		totalPower += (maxRed * maxGreen * maxBlue)

		return nil
	})

	return totalPower
}
