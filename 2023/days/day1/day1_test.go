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
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.Part1(tt.filename, logger)
			if result != tt.want {
				t.Fatalf(`d.Part1("%s", logger) = %d, want %d`, tt.filename, result, tt.want)
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
		want     int64
	}{
		{name: "sample", filename: "sample2.txt", want: 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := d.Part2(tt.filename, logger)
			if result != tt.want {
				t.Fatalf(`d.Part2("%s", logger) = %d, want %d`, tt.filename, result, tt.want)
			}
		})
	}
}
