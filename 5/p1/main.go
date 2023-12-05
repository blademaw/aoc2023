package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processMaps(sourceNum int, maps [][][]int) int {
	curNum := sourceNum

	for _, m := range maps {
		for _, line := range m {
			destStart, sourceStart, span := line[0], line[1], line[2]

			if curNum >= sourceStart && curNum < sourceStart+span {
				curNum = destStart + (curNum - sourceStart)
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

	// Getting seeds
	r := regexp.MustCompile(`(\d+)`)
	seedNumsStr := r.FindAllString(maps[0], -1)
	seedNums := make([]int, len(seedNumsStr))

	for i, n := range seedNumsStr {
		seedNums[i], _ = strconv.Atoi(n)
	}

	// Getting maps
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

	res := make([]int, len(seedNums))
	for i, n := range seedNums {
		res[i] = processMaps(n, mapInts)
	}

	minimum := res[0]
	for _, n := range res[1:] {
		if n < minimum {
			minimum = n
		}
	}

	fmt.Println(minimum)
}
