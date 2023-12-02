package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	valMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var fwdRePatSlice, bwdRePatSlice, keys []string
	var fwdRePat, bwdRePat string
	for k := range valMap {
		keys = append(keys, k)
	}
	for _, k := range keys {
		revK := util.Reverse(k)
		fwdRePatSlice = append(fwdRePatSlice, k)
		bwdRePatSlice = append(bwdRePatSlice, revK)
		valMap[revK] = valMap[k]
	}

	fwdRePatSlice = append(fwdRePatSlice, "[0-9]")
	fwdRePat = fmt.Sprintf("(%s)", strings.Join(fwdRePatSlice, "|"))

	bwdRePatSlice = append(bwdRePatSlice, "[0-9]")
	bwdRePat = fmt.Sprintf("(%s)", strings.Join(bwdRePatSlice, "|"))

	fwdRe := regexp.MustCompile(fwdRePat)
	bwdRe := regexp.MustCompile(bwdRePat)

	level.Debug(logger).Log("fwdRePat", fwdRePat)
	level.Debug(logger).Log("bwdRePat", bwdRePat)
	level.Debug(logger).Log("valMap", fmt.Sprintf("%s", valMap))

	var values []int64
	util.ReadPuzzleInput(filename, logger, func(line string) error {
		level.Debug(logger).Log("line", line)
		revLine := util.Reverse(line)

		fwd := fwdRe.FindAllString(line, -1)
		bwd := bwdRe.FindAllString(revLine, -1)

		level.Debug(logger).Log("fwd", fmt.Sprintf("%s", fwd))
		level.Debug(logger).Log("bwd", fmt.Sprintf("%s", bwd))

		first, ok := valMap[fwd[0]]
		if !ok {
			first = fwd[0]
		}

		last, ok := valMap[bwd[0]]
		if !ok {
			last = bwd[0]
		}

		sVal := fmt.Sprintf("%s%s", first, last)
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
