package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const SYMBOLS = `~!@#$%^&*_+-=[]{}|;:,/<>?`

// Determines if a row and span in the schema is near a symbol.
func nearSymbol(row int, span []int, schema []string) bool {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(schema) {
			continue
		}

		for j := span[0] - 1; j <= span[1]; j++ {
			if j < 0 || j >= len(schema[0]) || (i == row && j >= span[0] && j < span[1]) {
				continue
			}

			if strings.ContainsAny(string(schema[i][j]), SYMBOLS) {
				fmt.Println("found num", schema[row][span[0]:span[1]])
				return true
			}
		}
	}

	return false
}

func main() {
	file := flag.String("file", "data.txt", "the file of the schema.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	schema := strings.TrimSpace(string(dat))
	lines, _ := utils.ReadLines(*file)
	_ = schema

	r, _ := regexp.Compile(`(\d+)`)

	total := 0
	for i, line := range lines {
		idx := r.FindAllIndex([]byte(line), -1)

		for _, numSpan := range idx {
			if nearSymbol(i, numSpan, lines) {
				n, _ := strconv.Atoi(line[numSpan[0]:numSpan[1]])
				total += n
			}
		}
	}

	fmt.Println(total)
}
