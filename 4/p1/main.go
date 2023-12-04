package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(\d+)`)

func cardPoints(card string) int {
	cardNums := strings.Split(card, " | ")

	winningNums := make(map[string]bool)
	for _, n := range r.FindAllString(cardNums[0], -1) {
		winningNums[n] = true
	}

	res := 0
	for _, n := range r.FindAllString(cardNums[1], -1) {
		if _, ok := winningNums[n]; ok {
			if res == 0 {
				res = 1
			} else {
				res *= 2
			}
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

	total := 0
	for _, cardLine := range lines {
		card := strings.Split(cardLine, ": ")[1]
		total += cardPoints(card)
	}

	fmt.Println(total)
}
