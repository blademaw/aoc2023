package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func int64Sum(int64s []int64) int64 {
	var sum int64 = 0
	for _, n := range int64s {
		sum += n
	}
	return sum
}

func arrangements(line string, positions map[int64]int64) int64 {
	x := strings.Split(line, " ")
	initRecord := x[0]

	var fns []int64
	for _, nStr := range strings.Split(x[1], ",") {
		nint64, _ := strconv.Atoi(nStr)
		fns = append(fns, int64(nint64))
	}

	// Cycle to get proper input
	var initRecords []string
	for i := 0; i < 5; i++ {
		initRecords = append(initRecords, initRecord)
	}
	record := strings.Join(initRecords, "?")

	var ns []int64
	for i := 0; i < 5; i++ {
		ns = append(ns, fns...)
	}

	for i, cont := range ns {
		newPositions := make(map[int64]int64)

		for k, v := range positions {
			for n := k; n < int64(len(record))-int64Sum(ns[i+1:])+int64(len(ns[i+1:])); n++ {
				if n+cont-1 < int64(len(record)) && !strings.ContainsRune(record[n:n+cont], '.') {
					if (i == len(ns)-1 && !strings.ContainsRune(record[n+cont:], '#')) || (i < len(ns)-1 && n+cont < int64(len(record)) && record[n+cont] != '#') {
						if vN, ok := newPositions[n+cont+1]; ok {
							newPositions[n+cont+1] = vN + v
						} else {
							newPositions[n+cont+1] = v
						}
					}

				}
				if record[n] == '#' {
					break
				}
			}
			positions = newPositions
		}
	}

	var res int64 = 0
	for _, v := range positions {
		res += v
	}
	return res
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as records")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	positions := make(map[int64]int64)
	positions[0] = 1

	var res int64 = 0
	for _, line := range lines {
		res += arrangements(line, positions)
	}

	fmt.Println(res)
}
