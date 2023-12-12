package day10

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay10_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day10{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day10.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay10_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		// {name: "sample2", filename: "sample2.txt", want: 4},
		{name: "sample3", filename: "sample3.txt", want: 4},
		// {name: "sample4", filename: "sample4.txt", want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day10{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day10.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
