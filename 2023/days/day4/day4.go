package day4

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	util "github.com/thelande/adventofcode/common"
)

type Day4 struct{}

type Card struct {
	WinningNumbers []int64
	CardNumbers    []int64
	Copies         int
}

func NewCardFromLine(line string) *Card {
	var card Card
	card.Copies = 1

	parts := strings.Split(line, ":")    // card number and numbers
	parts = strings.Split(parts[1], "|") // sWinning numbers and our numbers
	sWinning := strings.Split(strings.Trim(parts[0], " "), " ")
	sCard := strings.Split(strings.Trim(parts[1], " "), " ")

	for _, v := range sWinning {
		if v == "" {
			continue
		}
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		card.WinningNumbers = append(card.WinningNumbers, val)
	}

	for _, v := range sCard {
		if v == "" {
			continue
		}
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		card.CardNumbers = append(card.CardNumbers, val)
	}
	slices.Sort(card.WinningNumbers)
	slices.Sort(card.CardNumbers)

	return &card
}

// Returns the number of winning numbers matched on this card.
func (card Card) Matches(logger log.Logger) int {
	var count int

	winIdx := 0
	for _, num := range card.CardNumbers {
		// level.Debug(logger).Log("num", num, "winIdx", winIdx, "winningNums[winIdx]", card.WinningNumbers[winIdx], "count", count)
		for ; winIdx < len(card.WinningNumbers) && card.WinningNumbers[winIdx] < num; winIdx++ {
			// Skip over entries in winningNums until we get to a number
			// that is not less than the current card number.
		}

		if winIdx >= len(card.WinningNumbers) {
			// We've exceeded the list of winning numbers.
			break
		}

		if card.WinningNumbers[winIdx] == num {
			// level.Debug(logger).Log("num", num, "winningNums[winIdx]", card.WinningNumbers[winIdx])
			count++
		}
	}

	return count
}

func (d Day4) Part1(filename string, logger log.Logger) int64 {
	var total int64

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		card := NewCardFromLine(line)

		level.Debug(logger).Log("cardNums", fmt.Sprintf("%d", card.CardNumbers))
		level.Debug(logger).Log("winningNums", fmt.Sprintf("%d", card.WinningNumbers))

		matches := card.Matches(logger)
		if matches > 0 {
			total += int64(math.Pow(2, float64(matches-1)))
		}

		return nil
	})

	return total
}

func (d Day4) Part2(filename string, logger log.Logger) int64 {
	var cards []Card

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		cards = append(cards, *NewCardFromLine(line))
		return nil
	})

	numCards := len(cards)
	count := int64(numCards)
	for idx, card := range cards {
		matches := card.Matches(logger)
		level.Debug(logger).Log("idx", idx, "card.Copies", card.Copies, "matches", matches)

		if matches == 0 {
			continue
		}

		// Increase the number of copies for the next matches number of cards.
		for i := 1; i <= matches && idx+i < numCards; i++ {
			level.Debug(logger).Log("idx+i", idx+i)
			cards[idx+i].Copies += card.Copies
			count += int64(card.Copies)
		}
	}

	return count
}
