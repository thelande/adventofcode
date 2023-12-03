package day3

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay3_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 4361},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day3{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day3.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 467835},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day3{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day3.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
