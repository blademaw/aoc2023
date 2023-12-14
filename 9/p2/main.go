package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func allZeros(ns []int) bool {
	for _, n := range ns {
		if n != 0 {
			return false
		}
	}
	return true
}

func computeDiffs(ns []int) []int {
	diffArr := make([]int, len(ns)-1)
	for i := 1; i < len(ns); i++ {
		diffArr[i-1] = ns[i] - ns[i-1]
	}

	return diffArr
}

func nextRes(ns []int) int {
	diffs := make([][]int, len(ns))
	diffs[0] = ns

	// Compute diffs until 0
	for i := 1; i < len(ns); i++ {
		if allZeros(diffs[i-1]) {
			break
		}

		diffs[i] = computeDiffs(diffs[i-1])
	}

	// Subtract each difference from the first term to get the 0-th
	i, sum := 0, 0
	for len(diffs[i]) > 0 && i < len(diffs) {
		if i%2 == 0 {
			sum += diffs[i][0]
		} else {
			sum -= diffs[i][0]
		}
		i++
	}

	return sum
}

func main() {
	file := flag.String("file", "data.txt", "the file of OASIS records.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	sum := 0
	for _, line := range lines {
		sNums := strings.Split(line, " ")

		ns := make([]int, len(sNums))
		for i, s := range sNums {
			ns[i], _ = strconv.Atoi(s)
		}

		sum += nextRes(ns)
	}

	fmt.Println(sum)
}
