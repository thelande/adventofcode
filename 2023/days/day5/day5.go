package day5

import (
	"regexp"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gosuri/uiprogress"
	util "github.com/thelande/adventofcode/common"
)

const MAX_PROCS = 12

type Day5 struct{}

type SourceDestMap struct {
	DestRangeStart   int64
	SourceRangeStart int64
	Count            int64
}

type AlmanacMap struct {
	SourceDestMaps []SourceDestMap
	Source         string
	Dest           string
}

type SeedRange struct {
	Start, End int64
}

func (m *AlmanacMap) Translate(source int64) int64 {
	for _, srcDstMap := range m.SourceDestMaps {
		srcEnd := srcDstMap.SourceRangeStart + srcDstMap.Count - 1
		if srcDstMap.SourceRangeStart <= source && source <= srcEnd {
			// Found the source in a map, translate it.
			offset := source - srcDstMap.SourceRangeStart
			return srcDstMap.DestRangeStart + offset
		}
	}

	// Return the source if no maps contain it.
	return source
}

func GetSeeds(line string) []int64 {
	parts := strings.Split(line, ":")
	return util.NumListToSlice(parts[1])
}

// Returns the seed value in dest.
func GetValueForSeed(seed int64, dest string, maps map[string]*AlmanacMap, logger log.Logger) int64 {
	var currSrc, currDest string
	value := seed
	currSrc = "seed"
	for currDest != dest {
		level.Debug(logger).Log("value", value, "currSrc", currSrc, "currDest", currDest)
		value = maps[currSrc].Translate(value)
		currSrc = maps[currSrc].Dest
		currDest = maps[currSrc].Dest
	}
	value = maps[currSrc].Translate(value)
	return value
}

func (d Day5) Part1(filename string, logger log.Logger) int64 {
	var seeds []int64
	var maps map[string]*AlmanacMap = map[string]*AlmanacMap{}
	var curr *AlmanacMap

	mapRegexp := regexp.MustCompile(`^(.+)-to-(.+) map:$`)
	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		if lineno == 0 {
			// First line are the seeds
			seeds = GetSeeds(line)
		} else if len(line) > 0 {
			matches := mapRegexp.FindStringSubmatch(line)
			if len(matches) > 0 {
				curr = &AlmanacMap{
					Source:         matches[1],
					Dest:           matches[2],
					SourceDestMaps: []SourceDestMap{},
				}
				maps[curr.Source] = curr
			} else {
				values := util.NumListToSlice(line)
				curr.SourceDestMaps = append(curr.SourceDestMaps, SourceDestMap{
					DestRangeStart:   values[0],
					SourceRangeStart: values[1],
					Count:            values[2],
				})
			}
		}
		return nil
	})

	var minLoc int64 = -1
	for _, seed := range seeds {
		loc := GetValueForSeed(seed, "location", maps, logger)
		level.Debug(logger).Log("location", loc)
		if minLoc < 0 || loc < minLoc {
			minLoc = loc
		}
	}

	return minLoc
}

func (d Day5) Part2(filename string, logger log.Logger) int64 {
	var seedRanges []SeedRange
	var maps map[string]*AlmanacMap = map[string]*AlmanacMap{}
	var curr *AlmanacMap

	mapRegexp := regexp.MustCompile(`^(.+)-to-(.+) map:$`)
	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		if lineno == 0 {
			// First line are the seeds
			seeds := GetSeeds(line)
			for i := 0; i < len(seeds); i += 2 {
				r := SeedRange{
					Start: seeds[i],
					End:   seeds[i] + seeds[i+1] - 1,
				}
				seedRanges = append(seedRanges, r)
			}
		} else if len(line) > 0 {
			matches := mapRegexp.FindStringSubmatch(line)
			if len(matches) > 0 {
				curr = &AlmanacMap{
					Source:         matches[1],
					Dest:           matches[2],
					SourceDestMaps: []SourceDestMap{},
				}
				maps[curr.Source] = curr
			} else {
				values := util.NumListToSlice(line)
				curr.SourceDestMaps = append(curr.SourceDestMaps, SourceDestMap{
					DestRangeStart:   values[0],
					SourceRangeStart: values[1],
					Count:            values[2],
				})
			}
		}
		return nil
	})

	nRanges := len(seedRanges)
	minLoc := make(chan int64, nRanges)
	uiprogress.Start()
	for i, r := range seedRanges {
		level.Info(logger).Log("msg", "finding min for range", "i", i, "count", r.End-r.Start+1)
		FindMinLocationForRange(r, maps, logger, minLoc)
	}

	var currMin int64 = -1
	for i := 0; i < nRanges; i++ {
		loc := <-minLoc
		if currMin < 0 || loc < currMin {
			currMin = loc
		}
	}

	return currMin
}

func findSubrange(start, end int64, maps map[string]*AlmanacMap, logger log.Logger, minLoc chan int64) {
	bar := uiprogress.AddBar(int(end - start + 1))
	bar.AppendCompleted()
	var ourMin int64 = -1
	for seed := start; seed <= end; seed++ {
		// fmt.Printf("Range %2d / %02d - Seed %09d / %09d\r", i, nRanges, seed-r.Start+1, count)
		loc := GetValueForSeed(seed, "location", maps, logger)
		if ourMin < 0 || loc < ourMin {
			ourMin = loc
		}
		bar.Incr()
	}
	minLoc <- ourMin
}

func FindMinLocationForRange(r SeedRange, maps map[string]*AlmanacMap, logger log.Logger, minLoc chan int64) {
	// The total number of seeds to check
	count := r.End - r.Start + 1

	// How many seeds to check, per thread.
	countPerProc := count / MAX_PROCS

	// If the number of seeds to check per thread is less than one, then just run a single thread for
	// all seeds.
	if countPerProc < 1 {
		findSubrange(r.Start, r.End, maps, logger, minLoc)
		return
	}

	rangeChan := make(chan int64, MAX_PROCS)

	for i := 0; i < MAX_PROCS; i++ {
		offset := countPerProc * int64(i)
		endOffset := countPerProc * int64(i+1)
		go findSubrange(r.Start+offset, r.Start+endOffset, maps, logger, rangeChan)
	}

	var ourMin int64 = -1
	for i := 0; i < MAX_PROCS; i++ {
		loc := <-rangeChan
		if ourMin < 0 || loc < ourMin {
			ourMin = loc
		}
	}

	minLoc <- ourMin
}
