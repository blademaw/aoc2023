package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
)

// Finds the next (row, col) location for a given pipe at (i, j) with a success
// indicator. Returns false if the implied transition is not valid or out of
// bounds.
func nextPipe(previ int, prevj int, i int, j int, char rune, maxI int,
	maxJ int) (int, int, bool) {
	var newi, newj int

	if char == '|' {
		if i < previ {
			newi, newj = i-1, j
		} else if i > previ {
			newi, newj = i+1, j
		} else {
			// If there was no vertical difference, cannot go into a vertical pipe
			return 0, 0, false
		}
	} else if char == '-' {
		if j < prevj {
			newi, newj = i, j-1
		} else if j > prevj {
			newi, newj = i, j+1
		} else {
			// If no horizontal difference, cannot go into horizontal pipe
			return 0, 0, false
		}
	} else if char == 'L' {
		if i > previ {
			newi, newj = i, j+1
		} else if j < prevj {
			newi, newj = i-1, j
		} else {
			// Can only enter L from above/east
			return 0, 0, false
		}
	} else if char == 'J' {
		if i > previ {
			newi, newj = i, j-1
		} else if j > prevj {
			newi, newj = i-1, j
		} else {
			return 0, 0, false
		}
	} else if char == '7' {
		if i < previ {
			newi, newj = i, j-1
		} else if j > prevj {
			newi, newj = i+1, j
		} else {
			return 0, 0, false
		}
	} else if char == 'F' {
		if i < previ {
			newi, newj = i, j+1
		} else if j < prevj {
			newi, newj = i+1, j
		} else {
			return 0, 0, false
		}
	} else {
		// . and S do not lead anywhere
		return 0, 0, false
	}

	if newi > maxI || newj > maxJ || newi < 0 || newj < 0 {
		return 0, 0, false
	}

	return newi, newj, true
}

func traverse(depth int, previ int, prevj int, i int, j int, pipes [][]rune, maxI int, maxJ int) (bool, int) {
	if pipes[i][j] == 'S' {
		return true, depth
	}

	if nextI, nextJ, ok := nextPipe(previ, prevj, i, j, pipes[i][j], maxI, maxJ); ok {
		return traverse(depth+1, i, j, nextI, nextJ, pipes, maxI, maxJ)
	}

	return false, 0
}

func findMaxDist(pipes [][]rune, i0 int, j0 int) int {
	maxI, maxJ := len(pipes)-1, len(pipes[0])-1

	// Find starting positions
	var start [][]int
	for _, transforms := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		posI, posJ := i0+transforms[0], j0+transforms[1]

		// Is only valid if in the grid and be navigated to by the prev square
		if posI <= maxI && posJ <= maxJ && posI >= 0 && posJ >= 0 {
			_, _, ok := nextPipe(i0, j0, posI, posJ, pipes[posI][posJ], maxI, maxJ)
			if ok {
				start = append(start, []int{posI, posJ})
			}
		}
	}

	k := 0
	found, depth := false, 0
	for !found && k < len(start) {
		i, j := start[k][0], start[k][1]

		found, depth = traverse(1, i0, j0, i, j, pipes, maxI, maxJ)
		k++
	}

	return depth / 2
}

func main() {
	// Main idea of algorithm: from root node (pipe) (S), traverse recursively
	// from each branch until you find the root node again. At each step, keep a
	// record of the depth (length) and whether the root node was found. Take
	// floor(len/2) to find the maximum furthest point.

	file := flag.String("file", "data.txt", "file to parse as pipes.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	pipes := make([][]rune, len(lines))
	for i, line := range lines {
		pipes[i] = []rune(line)
	}

	var i0, j0 int
	foundStart := false

	for i := range pipes {
		for j := range pipes[i] {
			if pipes[i][j] == 'S' {
				i0, j0 = i, j
				break
			}
		}
		if foundStart {
			break
		}
	}

	fmt.Println(findMaxDist(pipes, i0, j0))
}
