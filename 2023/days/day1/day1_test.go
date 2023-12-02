package day1_test

import (
	"testing"

	"github.com/prometheus/common/promlog"
	"github.com/thelande/adventofcode/2023/days/day1"
)

func TestDay1Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	tests := []struct {
		name     string
		filename string
		exp      int64
	}{
		{name: "sample", filename: "sample.txt", exp: 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := day1.Day1Part1(tt.filename, logger)
			if result != tt.exp {
				t.Fatalf(`Day1Part1("%s", logger) = %d, want %d`, tt.filename, result, tt.exp)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	tests := []struct {
		name     string
		filename string
		exp      int64
	}{
		{name: "sample", filename: "sample2.txt", exp: 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := day1.Day1Part2(tt.filename, logger)
			if result != tt.exp {
				t.Fatalf(`Day1Part2("%s", logger) = %d, want %d`, tt.filename, result, tt.exp)
			}
		})
	}
}
