package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const NUMBERS = `1234567890`

// Returns the gear ratio for a given gear position and two positions of unique
// numbers.
func gearRatio(row int, col int, locs [][]int, schema []string) int {
	nums := []int{0, 0}

	for i, loc := range locs {
		// For each location, try to expand either direction until we hit a symbol
		// or dot to get each number.
		line := schema[row+loc[0]]
		start, end := col+loc[1], col+loc[1]

		for start > 0 && unicode.IsDigit(rune(line[start-1])) {
			start--
		}

		for end < len(line)-1 && unicode.IsDigit(rune(line[end+1])) {
			end++
		}

		nums[i], _ = strconv.Atoi(line[start : end+1])
	}

	return nums[0] * nums[1]
}

// Returns two positions relative to the gear of unique numbers surrounding a
// gear if they can be found.
func gearLocs(row int, col int, schema []string) ([][]int, bool) {
	num1, num2 := -1, -1

	var prevDigit bool

	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(schema) {
			continue
		}
		prevDigit = false

		for j := col - 1; j <= col+1; j++ {
			if j < 0 || j >= len(schema[0]) || (i == row && j == col) {
				prevDigit = false
				continue
			}

			if strings.ContainsAny(string(schema[i][j]), NUMBERS) {
				relativeLoc := (i-(row-1))*3 + (j - (col - 1))

				if num1 == -1 {
					num1 = relativeLoc
				} else if prevDigit && num2 == -1 {
					num1 = relativeLoc
				} else if prevDigit && num2 != -1 {
					num2 = relativeLoc
				} else if !prevDigit && num2 == -1 {
					num2 = relativeLoc
				} else {
					return nil, false
				}

				prevDigit = true
			} else {
				prevDigit = false
			}
		}
	}

	if num1 == -1 || num2 == -1 {
		return nil, false
	}

	return [][]int{{num1/3 - 1, num1%3 - 1}, {num2/3 - 1, num2%3 - 1}},
		true
}

func main() {
	file := flag.String("file", "data.txt", "the file of the schema.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	r, _ := regexp.Compile(`\*`)

	total := 0
	for i, line := range lines {
		idx := r.FindAllIndex([]byte(line), -1)

		for _, span := range idx {
			if locs, ok := gearLocs(i, span[0], lines); ok {
				total += gearRatio(i, span[0], locs, lines)
			}
		}
	}

	fmt.Println(total)
}
