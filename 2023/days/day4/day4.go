package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day4 struct{}

func (d Day4) Part1(filename string, logger log.Logger) int64 {
	var total int64

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		parts := strings.Split(line, ":")    // card number and numbers
		parts = strings.Split(parts[1], "|") // winning numbers and our numbers
		winning := strings.Split(strings.Trim(parts[0], " "), " ")
		card := strings.Split(strings.Trim(parts[1], " "), " ")

		var winningNums, cardNums []int64
		for _, v := range winning {
			if v == "" {
				continue
			}
			val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			winningNums = append(winningNums, val)
		}

		for _, v := range card {
			if v == "" {
				continue
			}
			val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}
			cardNums = append(cardNums, val)
		}

		// Sort the lists, then loop over the card numbers and check if they
		// exist in the winning numbers list.
		var score int64
		slices.Sort(winningNums)
		slices.Sort(cardNums)

		level.Debug(logger).Log("cardNums", fmt.Sprintf("%d", cardNums))
		level.Debug(logger).Log("winningNums", fmt.Sprintf("%d", winningNums))

		winIdx := 0
		for _, num := range cardNums {
			level.Debug(logger).Log("num", num, "winIdx", winIdx, "winningNums[winIdx]", winningNums[winIdx], "score", score)
			for ; winIdx < len(winningNums) && winningNums[winIdx] < num; winIdx++ {
				// Skip over entries in winningNums until we get to a number
				// that is not less than the current card number.
			}

			if winIdx >= len(winningNums) {
				// We've exceeded the list of winning numbers.
				break
			}

			if winningNums[winIdx] == num {
				level.Debug(logger).Log("num", num, "winningNums[winIdx]", winningNums[winIdx])
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		total += score

		return nil
	})

	return total
}

func (d Day4) Part2(filename string, logger log.Logger) int64 {
	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		return nil
	})

	return 0
}
