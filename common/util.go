package util

import (
	"bufio"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

/**
 * Read the puzzle input file, filename, and call lineCallback for each line read from the file.
 */
func ReadPuzzleInput(filename string, logger log.Logger, lineCallback func(line string) error) error {
	level.Debug(logger).Log("msg", "Loading puzzle input", "filename", filename)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lineno int
	for lineno = 0; scanner.Scan(); lineno++ {
		line := scanner.Text()
		level.Debug(logger).Log("line", line)
		if err = lineCallback(line); err != nil {
			return err
		}
	}

	level.Debug(logger).Log("msg", "lines read", "count", lineno)

	return nil
}

/**
 * Sum the values of s and return the total.
 */
func SumSlice(s []int64) int64 {
	var total int64
	for _, val := range s {
		total += val
	}
	return total
}

/**
 * Reverse and return s
 */
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
