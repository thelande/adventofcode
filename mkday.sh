#!/bin/bash
#
# Create base files for a new day.
#
set -e

function usage {
    echo "usage: ${0##*/} <year> <day number>"
}

test -z "$2" && { usage; exit 2; }
YEAR="$1"
DAY="$2"

PAT='^[0-9]{1,2}$'
if ! [[ $DAY =~ $PAT ]]; then
	echo "error: <day number> must be a number."
	exit 1
fi

DAY_DIR="$YEAR/days/day$DAY"
DAY_FILE="$DAY_DIR/day$DAY.go"
DAY_TEST_FILE="$DAY_DIR/day${DAY}_test.go"
DAY_INPUT_FILE="$DAY_DIR/input.txt"

# Make sure the day does not already exist.
if [[ -f "$DAY_FILE" ]]; then
    echo "error: $DAY_FILE already exists."
    exit 1
fi

mkdir -p "$DAY_DIR"

cat<<! >$DAY_FILE
package day$DAY

import (
    "github.com/go-kit/log"
    util "github.com/thelande/adventofcode/common"
)

type Day$DAY struct{}

func (d Day$DAY) Part1(filename string, logger log.Logger) int64 {
    var value int64

    util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
        return nil
    })

    return value
}

func (d Day$DAY) Part2(filename string, logger log.Logger) int64 {
    var value int64

    util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
        return nil
    })

    return value
}
!

cat<<! >$DAY_TEST_FILE
package day$DAY

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay${DAY}_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day$DAY{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day$DAY.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay${DAY}_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day$DAY{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day$DAY.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
!
