package util

import "github.com/go-kit/log"

type Day interface {
	Part1(filename string, logger log.Logger) int64
	Part2(filename string, logger log.Logger) int64
}
