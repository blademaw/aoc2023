package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	file := flag.String("file", "data.txt", "the file to parse.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	total := 0
	for _, line := range lines {
		i, j := 0, len(line)-1

		for !unicode.IsDigit(rune(line[j])) && j >= 0 {
			j--
		}

		for !unicode.IsDigit(rune(line[i])) && i < len(line) {
			i++
		}

		n, _ := strconv.Atoi(string(line[i]) + string(line[j]))
		total += n
	}

	fmt.Println(total)
}
