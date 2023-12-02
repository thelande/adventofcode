package day1_test

import (
	"testing"

	"github.com/prometheus/common/promlog"
	"github.com/thelande/adventofcode/2023/days/day1"
)

func TestDay1Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	d := day1.Day1{}

	tests := []struct {
		name     string
		filename string
		exp      int64
	}{
		{name: "sample", filename: "sample.txt", exp: 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.Part1(tt.filename, logger)
			if result != tt.exp {
				t.Fatalf(`d.Part1("%s", logger) = %d, want %d`, tt.filename, result, tt.exp)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	d := day1.Day1{}

	tests := []struct {
		name     string
		filename string
		exp      int64
	}{
		{name: "sample", filename: "sample2.txt", exp: 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.Part2(tt.filename, logger)
			if result != tt.exp {
				t.Fatalf(`d.Part2("%s", logger) = %d, want %d`, tt.filename, result, tt.exp)
			}
		})
	}
}
