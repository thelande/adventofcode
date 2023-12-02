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
