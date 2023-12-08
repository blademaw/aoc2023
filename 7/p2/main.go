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
	'J': -1,
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'Q': 9,
	'K': 10,
	'A': 11,
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

// Go doesn't have this built-in except for reflect.DeepEqual, which
// is apparently slow? Just test to see if two int lists are the same.
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
	jokers := 0 // Jokers just pretend to be the highest-count card

	for _, c := range hand {
		if c == 'J' {
			jokers++
		} else {
			m[c] += 1
		}
	}

	// Avoid initializing a list of length 0
	var counts []int
	if len(m) == 0 {
		counts = []int{0}
	} else {
		counts = make([]int, len(m))
	}

	// Put the card count map into a list and sort to identify types
	i := 0
	for _, c := range m {
		counts[i] = c
		i++
	}
	slices.Sort(counts)
	counts[len(counts)-1] += jokers

	// Find a type
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
