package day11

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay11_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 374},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day11{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day11.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay11_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 82000210},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day11{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day11.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
