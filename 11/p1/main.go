package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"math"
)

// Expand a universe
func expandUniverse(universe [][]rune) [][]rune {
	// Identify rows and columns that are all space
	allSpaceRows := func(i int) bool {
		for _, r := range universe[i] {
			if r != '.' {
				return false
			}
		}
		return true
	}

	allSpaceCols := func(j int) bool {
		for _, c := range universe {
			if c[j] != '.' {
				return false
			}
		}
		return true
	}

	var rows, cols []int
	for i := range universe {
		if allSpaceRows(i) {
			rows = append(rows, i)
		}
	}
	for j := range universe[0] {
		if allSpaceCols(j) {
			cols = append(cols, j)
		}
	}

	rowTemplate := make([]rune, len(universe[0]))
	for i := range rowTemplate {
		rowTemplate[i] = '.'
	}

	// Add the necessary rows / columns
	addRow := func(i int) {
		universe = append(universe[:i+1], append([][]rune{rowTemplate}, universe[i+1:]...)...)
	}

	addCol := func(j int) {
		for i, line := range universe {
			universe[i] = append(line[:j+1], append([]rune{'.'}, line[j+1:]...)...)
		}
	}

	offset := 0
	for _, row := range rows {
		addRow(row + offset)
		offset++
	}

	offset = 0
	for _, col := range cols {
		addCol(col + offset)
		offset++
	}

	return universe
}

func shortestPathSum(universe [][]rune) float64 {
	var galaxies [][]int
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	// Uninhibited shortest path is just Manhattan distance
	total := 0.0
	for i, loc1 := range galaxies {
		for _, loc2 := range galaxies[i:] {
			total += math.Abs(float64(loc1[0]-loc2[0])) + math.Abs(float64(loc1[1]-loc2[1]))
		}
	}

	return total
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as a universe.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	universe := make([][]rune, len(lines))
	for i, line := range lines {
		universe[i] = []rune(line)
	}

	fmt.Println(int(shortestPathSum(expandUniverse(universe))))
}
