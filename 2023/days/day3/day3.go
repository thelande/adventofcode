package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day3 struct{}

var numRe = regexp.MustCompile(`[0-9]+`)

// Find and return the numbers in the line adjacent to the list of symbols.
func findNumbersNextToSymbols(line string, symbols [][]int, logger log.Logger) []int64 {
	var values []int64
	numbers := numRe.FindAllStringIndex(line, -1)

	level.Debug(logger).Log("line", line, "numbers", fmt.Sprintf("%d", numbers))
	for _, number := range numbers {
		for _, symbol := range symbols {
			//level.Debug(logger).Log("number[0]", number[0], "symbol[0]", symbol[0], "number[1]", number[1])
			if (number[0] <= symbol[0] && symbol[0] <= number[1]) || (number[0] <= symbol[1] && symbol[1] <= number[1]) {
				val, err := strconv.ParseInt(line[number[0]:number[1]], 10, 64)
				if err != nil {
					panic(err)
				}

				values = append(values, val)
			}
		}
	}
	level.Debug(logger).Log("values", fmt.Sprintf("%d", values))

	return values
}

func (d Day3) Part1(filename string, logger log.Logger) int64 {
	// Starting on the second line, look at the previous line for any symbols. For each
	// symbol found, look for numbers on the previous-previous line (top line), the
	// previous line (middle line), and the current line (bottom line).
	// Any numbers that touch a symbol (up/down/left/right/diagonal) are included.
	var topLine, midLine string
	var total int64
	symRe := regexp.MustCompile(`[^0-9.]`)

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		// Note the first line, but skip to the second to start processing.
		if lineno == 0 {
			midLine = line
			return nil
		}

		level.Debug(logger).Log("topLine", topLine, "midLine", midLine, "botLine", line)

		// Find the symbols on the middle line.
		symbols := symRe.FindAllStringIndex(midLine, -1)
		level.Debug(logger).Log("symbols", fmt.Sprintf("%d", symbols))

		// Skip processing if there are no symbols for this line.
		if len(symbols) > 0 {
			// Find the numbers on the top line, if set.
			if len(topLine) > 0 {
				total += util.SumSlice(findNumbersNextToSymbols(topLine, symbols, logger))
			}

			// Find the numbers on the middle line.
			total += util.SumSlice(findNumbersNextToSymbols(midLine, symbols, logger))

			// Find the numbers on the bottom line.
			total += util.SumSlice(findNumbersNextToSymbols(line, symbols, logger))
		}

		// Rotate the line trackers.
		topLine = midLine
		midLine = line

		return nil
	})

	return total
}

func valuesOnLineForGear(line string, numbers [][]int, gear []int) []int64 {
	var gearVals []int64

	// Look at the top line
	for _, num := range numbers {
		if (num[0] <= gear[0] && gear[0] <= num[1]) || (num[0] <= gear[1] && gear[1] <= num[1]) {
			val, err := strconv.ParseInt(line[num[0]:num[1]], 10, 64)
			if err != nil {
				panic(err)
			}
			gearVals = append(gearVals, val)
		}
	}

	return gearVals
}

func (d Day3) Part2(filename string, logger log.Logger) int64 {
	// Same approach as part 1, but only look for gears (*) and only
	// look at gears with two values returned.
	var topLine, midLine string
	var values []int64

	gearRe := regexp.MustCompile(`\*`)

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		// Note the first line, but skip to the second to start processing.
		if lineno == 0 {
			midLine = line
			return nil
		}

		level.Debug(logger).Log("topLine", topLine, "midLine", midLine, "botLine", line)

		// Find the gears on the middle line.
		gears := gearRe.FindAllStringIndex(midLine, -1)
		level.Debug(logger).Log("symbols", fmt.Sprintf("%d", gears))

		topNumbers := numRe.FindAllStringIndex(topLine, -1)
		midNumbers := numRe.FindAllStringIndex(midLine, -1)
		botNumbers := numRe.FindAllStringIndex(line, -1)

		for _, gear := range gears {
			var gearVals []int64
			gearVals = append(gearVals, valuesOnLineForGear(topLine, topNumbers, gear)...)
			gearVals = append(gearVals, valuesOnLineForGear(midLine, midNumbers, gear)...)
			gearVals = append(gearVals, valuesOnLineForGear(line, botNumbers, gear)...)

			if len(gearVals) > 2 {
				panic("Too many numbers attached to gear")
			} else if len(gearVals) == 2 {
				values = append(values, gearVals[0]*gearVals[1])
			}
		}

		// Rotate the line trackers.
		topLine = midLine
		midLine = line

		return nil
	})

	return util.SumSlice(values)
}
