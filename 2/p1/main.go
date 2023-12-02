package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var CUBEMAP = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isValidGame(game string) bool {
	obs := strings.Split(game, "; ")
	for _, ob := range obs {
		cubes := strings.Split(ob, ", ")
		for _, cube := range cubes {
			cubestr := strings.Split(cube, " ")
			cubeno, _ := strconv.Atoi(cubestr[0])

			if cubeno > CUBEMAP[cubestr[1]] {
				return false
			}
		}
	}

	return true
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	idSum := 0
	for _, gamestr := range lines {
		splitGame := strings.Split(gamestr, ": ")

		if isValidGame(splitGame[1]) {
			id, _ := strconv.Atoi(strings.Split(splitGame[0], " ")[1])
			idSum += id
		}
	}

	fmt.Println(idSum)
}
