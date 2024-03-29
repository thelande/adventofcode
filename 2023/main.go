package main

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/thelande/adventofcode/2023/days/day1"
	"github.com/thelande/adventofcode/2023/days/day10"
	"github.com/thelande/adventofcode/2023/days/day11"
	"github.com/thelande/adventofcode/2023/days/day2"
	"github.com/thelande/adventofcode/2023/days/day3"
	"github.com/thelande/adventofcode/2023/days/day4"
	"github.com/thelande/adventofcode/2023/days/day5"
	"github.com/thelande/adventofcode/2023/days/day6"
	"github.com/thelande/adventofcode/2023/days/day7"
	"github.com/thelande/adventofcode/2023/days/day8"
	"github.com/thelande/adventofcode/2023/days/day9"
	util "github.com/thelande/adventofcode/common"

	"github.com/prometheus/common/promlog"

	kingpin "github.com/alecthomas/kingpin/v2"
)

var (
	app         = kingpin.New("aoc2023", "Advent of Code, 2023")
	day         = app.Arg("day", "Which day to run.").Required().String()
	puzzleInput = app.Arg("input", "Puzzle input file.").Required().String()

	logLevel = app.Flag("log.level", "The log level").Short('l').Default("info").String()
	part     = app.Flag("part", "Which part to run (default, 0, is both)").Default("0").Short('p').Int()

	logger log.Logger
)

func main() {
	kingpin.CommandLine.UsageWriter(os.Stdout)
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	promlogConfig := &promlog.Config{
		Level: &promlog.AllowedLevel{},
	}
	promlogConfig.Level.Set(*logLevel)
	logger = promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "adventofcode", "year", 2023, "day", *day)

	var dayObj util.Day
	switch *day {
	case "day1":
		dayObj = day1.Day1{}
	case "day2":
		dayObj = day2.Day2{}
	case "day3":
		dayObj = day3.Day3{}
	case "day4":
		dayObj = day4.Day4{}
	case "day5":
		dayObj = day5.Day5{}
	case "day6":
		dayObj = day6.Day6{}
	case "day7":
		dayObj = day7.Day7{}
	case "day8":
		dayObj = day8.Day8{}
	case "day9":
		dayObj = day9.Day9{}
	case "day10":
		dayObj = day10.Day10{}
	case "day11":
		dayObj = day11.Day11{}
	default:
		level.Error(logger).Log("msg", "Unknown day")
		os.Exit(1)
	}

	var result1, result2 int64
	if *part == 0 || *part == 1 {
		result1 = dayObj.Part1(*puzzleInput, logger)
	}
	if *part == 0 || *part == 2 {
		result2 = dayObj.Part2(*puzzleInput, logger)
	}
	level.Info(logger).Log("part 1", result1, "part 2", result2)
}
