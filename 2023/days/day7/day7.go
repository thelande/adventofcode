package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	util "github.com/thelande/adventofcode/common"
)

var CARD_VALUES = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var CARD_VALUES2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Day7 struct{}

type Hand struct {
	Cards       string
	UniqueCards map[rune]int
	Bid         int64
}

func (h *Hand) Print() {
	fmt.Printf("%s %q %d\n", h.Cards, h.UniqueCards, h.Bid)
}

func NewHandFromLine(line string) *Hand {
	parts := strings.Split(line, " ")
	bidVal, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}

	set := make(map[rune]int)
	for _, card := range parts[0] {
		val, ok := set[card]
		if !ok {
			val = 1
		} else {
			val++
		}
		set[card] = val
	}

	return &Hand{
		Cards:       parts[0],
		UniqueCards: set,
		Bid:         bidVal,
	}
}

func NewHandFromLine2(line string) *Hand {
	parts := strings.Split(line, " ")
	bidVal, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}

	set := make(map[rune]int)
	for _, card := range parts[0] {
		val, ok := set[card]
		if !ok {
			val = 1
		} else {
			val++
		}
		set[card] = val
	}

	// Check for jokers and increase the maximum card group by the number of jokers.
	jokers := set['J']
	var maxRune rune
	var maxCount int
	if jokers > 0 {
		for r, c := range set {
			if r != 'J' {
				if c > maxCount {
					maxRune = r
					maxCount = c
				}
			}
		}

		set[maxRune] += jokers
		delete(set, 'J')
	}

	return &Hand{
		Cards:       parts[0],
		UniqueCards: set,
		Bid:         bidVal,
	}
}

func SortHands(a, b *Hand) int {
	// First check if one hand has fewer unique cards than the other. If so,
	// the hand with fewer unique cards is stronger: five-of-a-kind has 1
	// unique card whereas a four-of-a-kind has 2 unique cards.
	// If the number of unique cards is equal (four-of-a-kind vs. full house, or
	// three-of-a-kind and two-pair), then compare the counts of each card. Finally,
	// if the two hands have the same type, compare the individual cards from
	// left-to-right.
	if len(a.UniqueCards) < len(b.UniqueCards) {
		// A has fewer unique cards, thus a stronger hand.
		return 1
	} else if len(a.UniqueCards) > len(b.UniqueCards) {
		// A has more unique cards, thus a weaker hand.
		return -1
	}

	// A and B have the same number of unique cards. Check the make up of the unique
	// cards. A full house will have two numbers with 3 and 2, and a four-of-a-kind
	// will have two numbers with 4 and 1. Two pair and three-of-a-kind will have
	// three numbers with 2, 2, 1 and 3, 1, 1 respectively.

	// Find the max count of the cards to use in the next check. The hand with the
	// higher card count is stronger.
	var maxCountA, maxCountB int
	for _, count := range a.UniqueCards {
		maxCountA = max(maxCountA, count)
	}
	for _, count := range b.UniqueCards {
		maxCountB = max(maxCountB, count)
	}

	if maxCountA > maxCountB {
		return 1
	} else if maxCountA < maxCountB {
		return -1
	}

	// Same type of hand. Compare cards one-by-one.
	for i := 0; i < 5; i++ {
		if CARD_VALUES[rune(a.Cards[i])] > CARD_VALUES[rune(b.Cards[i])] {
			return 1
		} else if CARD_VALUES[rune(a.Cards[i])] < CARD_VALUES[rune(b.Cards[i])] {
			return -1
		}
	}

	// Hands are the same
	return 0
}

func SortHands2(a, b *Hand) int {
	// First check if one hand has fewer unique cards than the other. If so,
	// the hand with fewer unique cards is stronger: five-of-a-kind has 1
	// unique card whereas a four-of-a-kind has 2 unique cards.
	// If the number of unique cards is equal (four-of-a-kind vs. full house, or
	// three-of-a-kind and two-pair), then compare the counts of each card. Finally,
	// if the two hands have the same type, compare the individual cards from
	// left-to-right.
	if len(a.UniqueCards) < len(b.UniqueCards) {
		// A has fewer unique cards, thus a stronger hand.
		return 1
	} else if len(a.UniqueCards) > len(b.UniqueCards) {
		// A has more unique cards, thus a weaker hand.
		return -1
	}

	// A and B have the same number of unique cards. Check the make up of the unique
	// cards. A full house will have two numbers with 3 and 2, and a four-of-a-kind
	// will have two numbers with 4 and 1. Two pair and three-of-a-kind will have
	// three numbers with 2, 2, 1 and 3, 1, 1 respectively.

	// Find the max count of the cards to use in the next check. The hand with the
	// higher card count is stronger.
	var maxCountA, maxCountB int
	for _, count := range a.UniqueCards {
		maxCountA = max(maxCountA, count)
	}
	for _, count := range b.UniqueCards {
		maxCountB = max(maxCountB, count)
	}

	if maxCountA > maxCountB {
		return 1
	} else if maxCountA < maxCountB {
		return -1
	}

	// Same type of hand. Compare cards one-by-one.
	for i := 0; i < 5; i++ {
		if CARD_VALUES2[rune(a.Cards[i])] > CARD_VALUES2[rune(b.Cards[i])] {
			return 1
		} else if CARD_VALUES2[rune(a.Cards[i])] < CARD_VALUES2[rune(b.Cards[i])] {
			return -1
		}
	}

	// Hands are the same
	return 0
}

func (d Day7) Part1(filename string, logger log.Logger) int64 {
	var value int64
	var hands []*Hand

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		hands = append(hands, NewHandFromLine(line))
		return nil
	})

	slices.SortFunc[[]*Hand, *Hand](hands, SortHands)
	for rank, hand := range hands {
		value += int64(rank+1) * hand.Bid
	}

	return value
}

func (d Day7) Part2(filename string, logger log.Logger) int64 {
	var value int64
	var hands []*Hand

	util.ReadPuzzleInput(filename, logger, func(line string, lineno int) error {
		hands = append(hands, NewHandFromLine2(line))
		return nil
	})

	slices.SortFunc[[]*Hand, *Hand](hands, SortHands2)
	for rank, hand := range hands {
		value += int64(rank+1) * hand.Bid
	}

	return value
}
