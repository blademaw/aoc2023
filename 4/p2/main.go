package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(\d+)`)

func matches(card string) int {
	cardNums := strings.Split(card, " | ")

	winningNums := make(map[string]bool)
	for _, n := range r.FindAllString(cardNums[0], -1) {
		winningNums[n] = true
	}

	res := 0
	for _, n := range r.FindAllString(cardNums[1], -1) {
		if _, ok := winningNums[n]; ok {
			res++
		}
	}

	return res
}

func main() {
	file := flag.String("file", "data.txt", "the file of scratchcards.")
	flag.Parse()

	lines, err := utils.ReadLines(*file)
	if err != nil {
		panic(err)
	}

	// Number of scratchcards for card i
	numCards := make([]int, len(lines))
	for i := range numCards {
		numCards[i] = 1
	}

	for i, cardLine := range lines {
		card := strings.Split(cardLine, ": ")[1]
		matches := matches(card)

		for j := 1; j < matches+1; j++ {
			numCards[i+j] += numCards[i]
		}
	}

	sum := 0
	for _, n := range numCards {
		sum += n
	}

	fmt.Println(sum)
}
