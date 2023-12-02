package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

func Day1Part1(filename string, logger log.Logger) int64 {
	re := regexp.MustCompile(`[0-9]`)

	var values []int64
	util.ReadPuzzleInput(filename, logger, func(line string) error {
		level.Debug(logger).Log("line", line)
		split := re.FindAllString(line, -1)

		sVal := fmt.Sprintf("%s%s", split[0], split[len(split)-1])
		val, err := strconv.ParseInt(sVal, 10, 64)
		if err != nil {
			return err
		}
		level.Debug(logger).Log("val", val)
		values = append(values, val)

		return nil
	})

	return util.SumSlice(values)
}

func Day1Part2(filename string, logger log.Logger) int64 {
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)

	// valMap := map[string]string{
	// 	"one":   "1",
	// 	"two":   "2",
	// 	"three": "3",
	// 	"four":  "4",
	// 	"five":  "5",
	// 	"six":   "6",
	// 	"seven": "7",
	// 	"eight": "8",
	// 	"nine":  "9",
	// }

	var values []int64
	util.ReadPuzzleInput(filename, logger, func(line string) error {
		level.Debug(logger).Log("line", line)
		split := re.FindAllStringSubmatch(line, -1)

		level.Debug(logger).Log("split", fmt.Sprintf("%s", split))

		// first, ok := valMap[split[0]]
		// if !ok {
		// 	first = split[0]
		// }

		// last, ok := valMap[split[len(split)-1]]
		// if !ok {
		// 	last = split[len(split)-1]
		// }

		// sVal := fmt.Sprintf("%s%s", first, last)
		// val, err := strconv.ParseInt(sVal, 10, 64)
		// if err != nil {
		// 	return err
		// }
		// level.Debug(logger).Log("val", val)
		// values = append(values, val)

		return nil
	})

	return util.SumSlice(values)
}
