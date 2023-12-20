package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func isValid(config []rune, ns []int) int {
	var ns1 []int
	count := 0

	for _, r := range config {
		if r == '#' {
			count++
		} else if r == '.' {
			if count > 0 {
				ns1 = append(ns1, count)
				count = 0
			}
		}
	}
	if count > 0 {
		ns1 = append(ns1, count)
	}

	if len(ns1) != len(ns) {
		return 0
	}

	for i := range ns1 {
		if ns1[i] != ns[i] {
			return 0
		}
	}
	return 1
}

func numConfigs(config []rune, i int, ns []int) int {
	if i == len(config) {
		return isValid(config, ns)
	}

	if config[i] == '#' || config[i] == '.' {
		return numConfigs(config, i+1, ns)
	} else {
		c1, c2 := append([]rune{}, config...), append([]rune{}, config...)
		c1[i], c2[i] = '#', '.'
		return numConfigs(c1, i+1, ns) + numConfigs(c2, i+1, ns)
	}
}

func arrangements(line string) int {
	x := strings.Split(line, " ")
	config := []rune(x[0])

	var nums []int
	for _, n := range strings.Split(x[1], ",") {
		intN, _ := strconv.Atoi(n)
		nums = append(nums, intN)
	}

	return numConfigs(config, 0, nums)
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as record")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	res := 0
	for _, line := range lines {
		res += arrangements(line)
	}

	fmt.Println(res)
}
