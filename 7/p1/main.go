package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var CARDSTRENGTH = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func testEq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func handToType(hand string) HandType {
	m := make(map[rune]int)

	for _, c := range hand {
		m[c] += 1
	}

	counts := make([]int, len(m))
	i := 0
	for _, c := range m {
		counts[i] = c
		i++
	}
	slices.Sort(counts)

	if testEq(counts, []int{5}) {
		return FiveOfAKind
	} else if testEq(counts, []int{1, 4}) {
		return FourOfAKind
	} else if testEq(counts, []int{2, 3}) {
		return FullHouse
	} else if testEq(counts, []int{1, 1, 3}) {
		return ThreeOfAKind
	} else if testEq(counts, []int{1, 2, 2}) {
		return TwoPair
	} else if testEq(counts, []int{1, 1, 1, 2}) {
		return OnePair
	} else {
		return HighCard
	}
}

// Comparison function to compare two cards a and b. Outputs -1 if a < b, 1 if
// a > b, and 0 if a == b.
func compareCards(a string, b string) int {
	aType, bType := handToType(a), handToType(b)

	if aType < bType {
		return -1
	} else if aType > bType {
		return 1
	} else {

		for i := range a {
			if CARDSTRENGTH[rune(a[i])] > CARDSTRENGTH[rune(b[i])] {
				return 1
			} else if CARDSTRENGTH[rune(a[i])] < CARDSTRENGTH[rune(b[i])] {
				return -1
			}
		}
	}

	return 0
}

func scoreCamelCards(hands []string, handBidMap map[string]int) int {
	// Sort the cards, iterate and compute the score
	slices.SortFunc(hands, compareCards)

	res := 0
	for w, hand := range hands {
		res += (w + 1) * handBidMap[hand]
	}

	return res
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as hands of cards")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	hands, bidMap := make([]string, len(lines)), make(map[string]int)
	for i, line := range lines {
		contents := strings.Split(line, " ")

		hands[i] = contents[0]
		bidMap[contents[0]], _ = strconv.Atoi(contents[1])
	}

	fmt.Println(scoreCamelCards(hands, bidMap))
}
