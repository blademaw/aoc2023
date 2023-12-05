package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isInRanges(num int, ranges [][]int) bool {
	for _, r := range ranges {
		if num >= r[0] && num < r[0]+r[1] {
			return true
		}
	}

	return false
}

func findSeed(locNum int, maps [][][]int) int {
	curNum := locNum

	for _, m := range maps {
		for _, line := range m {
			destStart, sourceStart, span := line[0], line[1], line[2]

			if curNum >= destStart && curNum < destStart+span {
				curNum = sourceStart + (curNum - destStart)
				break
			}
		}
	}

	return curNum
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as maps.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	maps := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	// Getting ranges
	r := regexp.MustCompile(`(\d+)`)
	seedNumsStr := r.FindAllString(maps[0], -1)
	seedNums := make([][]int, len(seedNumsStr)/2)

	for i := 0; i < len(seedNumsStr); i += 2 {
		start, _ := strconv.Atoi(seedNumsStr[i])
		span, _ := strconv.Atoi(seedNumsStr[i+1])

		seedNums[i/2] = []int{start, span}
	}

	// Getting maps and reversing
	mapInts := make([][][]int, len(maps[1:]))
	for i, mStr := range maps[1:] {
		mapLines := strings.Split(mStr, "\n")[1:]
		mapRanges := make([][]int, len(mapLines))

		for j, line := range mapLines {
			nums := make([]int, 3)
			for k, n := range r.FindAllString(line, -1) {
				nums[k], _ = strconv.Atoi(n)
			}
			mapRanges[j] = nums
		}

		mapInts[i] = mapRanges
	}

	for i, j := 0, len(mapInts)-1; i < j; i, j = i+1, j-1 {
		mapInts[i], mapInts[j] = mapInts[j], mapInts[i]
	}

	foundMin := false
	possibleMin := 0
	for !foundMin {
		if isInRanges(findSeed(possibleMin, mapInts), seedNums) {
			foundMin = true
			break
		}
		possibleMin++
	}

	fmt.Println(possibleMin)
}
