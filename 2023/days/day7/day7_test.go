package day7

import (
	"testing"

	"github.com/prometheus/common/promlog"
)

func TestDay7_Part1(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 6440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day7{}
			if got := d.Part1(tt.filename, logger); got != tt.want {
				t.Errorf("Day7.Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay7_Part2(t *testing.T) {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{name: "sample", filename: "sample.txt", want: 5905},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day7{}
			if got := d.Part2(tt.filename, logger); got != tt.want {
				t.Errorf("Day7.Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDay7Part1(b *testing.B) {
	promlogConfig := &promlog.Config{
		Level: &promlog.AllowedLevel{},
	}
	promlogConfig.Level.Set("warn")
	logger := promlog.New(promlogConfig)
	d := Day7{}
	for i := 0; i < b.N; i++ {
		d.Part1("sample.txt", logger)
	}
}
