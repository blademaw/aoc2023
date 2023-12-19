package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"math"
)

const PenaltyAmount int = 1e6

// Transform a universe into a row and column list that convey the cost to
// travel to row/col i as A[i]. For all-space rows/cols this is 10^6, and 1
// otherwise.
func findExpanded(universe [][]rune) ([]bool, []bool) {
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

	rows, cols := make([]bool, len(universe)), make([]bool, len(universe[0]))

	for i := range universe {
		if allSpaceRows(i) {
			rows[i] = true
		} else {
			rows[i] = false
		}
	}

	for j := range universe[0] {
		if allSpaceCols(j) {
			cols[j] = true
		} else {
			cols[j] = false
		}
	}

	return rows, cols
}

func BFS(universe [][]rune, galaxies [][]int, s int, rows []bool, cols []bool) []int {
	// Get orthogonal vertices easily
	orthogonalVertices := func(u []int) [][]int {
		i, j := u[0], u[1]
		transforms := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
		var vertices [][]int

		for _, t := range transforms {
			vi, vj := i+t[0], j+t[1]
			if (vi >= 0 && vi < len(universe)) && (vj >= 0 && vj < len(universe[0])) {
				vertices = append(vertices, []int{vi, vj})
			}
		}
		return vertices
	}

	// Start of BFS
	dist, pred := make([][]int, len(universe)), make([][][]int, len(universe))
	for i := range dist {
		dist[i] = make([]int, len(universe[0]))
		pred[i] = make([][]int, len(universe[0]))

		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}

	queue := [][]int{galaxies[s]}
	dist[galaxies[s][0]][galaxies[s][1]] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range orthogonalVertices(u) {
			if dist[v[0]][v[1]] == math.MaxInt {
				dist[v[0]][v[1]] = dist[u[0]][u[1]] + 1
				pred[v[0]][v[1]] = u
				queue = append(queue, v)
			}
		}
	}

	// Result is list of distances to all other galaxies
	res := make([]int, len(galaxies))
	for i, g := range galaxies {
		if i == s {
			res[i] = 0
			continue
		}

		// Reconstruct path to find if need to penalize shortest distance
		path := [][]int{g}

		for !(g[0] == galaxies[s][0] && g[1] == galaxies[s][1]) {
			path = append(path, pred[g[0]][g[1]])
			g = pred[g[0]][g[1]]
		}

		for j := range path[1:] {
			last, cur := path[j], path[j+1]
			if last[0] == cur[0] {
				// if rows are the same, check if col incurs extra cost
				if cols[cur[1]] {
					res[i] += PenaltyAmount
				} else {
					res[i] += 1
				}
			} else {
				// if cols are the same, check if row incurs extra cost
				if rows[cur[0]] {
					res[i] += PenaltyAmount
				} else {
					res[i] += 1
				}
			}
		}
	}

	return res
}

func shortestPathSum(universe [][]rune) int {
	var galaxies [][]int
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	rows, cols := findExpanded(universe)

	var total int = 0
	for i := range galaxies {
		for _, d := range BFS(universe, galaxies, i, rows, cols)[i:] {
			total += d
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

	fmt.Println(shortestPathSum(universe))
}
