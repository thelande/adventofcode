package day9

import (
	"fmt"
	"slices"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day9 struct{}

func DiffSeq(seq []int64) []int64 {
	diffs := make([]int64, 0, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		diffs = append(diffs, seq[i+1]-seq[i])
	}
	return diffs
}

func ExtrapolateNext(seq []int64) int64 {
	if util.SumSlice(seq) != 0 {
		// Returns once we hit 0s
		return seq[len(seq)-1] + ExtrapolateNext(DiffSeq(seq))
	}
	return 0
}

func ComputeVal(seq []int64, logger log.Logger) int64 {
	level.Debug(logger).Log("seq", fmt.Sprintf("%d", seq))
	if len(seq) > 1 {
		return ComputeVal(DiffSeq(seq), logger)
	}
	return -1 * seq[0]
}

func (d Day9) Part1(filename string, logger log.Logger) int64 {
	var value int64

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		seq := util.NumListToSlice(line)
		seq = append(seq, 0)
		val := ComputeVal(seq, logger)
		level.Debug(logger).Log("seq", fmt.Sprintf("%d", seq), "val", val)
		value += val

		return nil
	})

	return value
}

func (d Day9) Part2(filename string, logger log.Logger) int64 {
	var value int64

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		seq := util.NumListToSlice(line)
		slices.Reverse(seq)
		seq = append(seq, 0)
		val := ComputeVal(seq, logger)
		level.Debug(logger).Log("seq", fmt.Sprintf("%d", seq), "val", val)
		value += val

		return nil
	})

	return value
}
