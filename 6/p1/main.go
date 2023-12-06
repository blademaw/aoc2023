package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Gets the margin of error for button presses b for a boat race of time t and
// record distance r.
func sizeOfMOE(t int, r int) int {
	// Computes bmax - bmin + 1, where bmin and bmax are the minimum and maximum
	// b such that d >= r + 1. Can compute this by defining time available t_a as
	// t - b and noting that the distance is d = t_a * b. Then to find the
	// function describing b, we have d = t_a * b = (t - b) * b, so b = (1/2)(t -
	// sqrt(t^2 - 4d)) or (1/2)(sqrt(t^2 - 4d) + t). Plugging in the desired
	// distance d = r + 1 we can find both bmin and bmax.
	bmin := math.Ceil((1.0 / 2.0) * (float64(t) - math.Sqrt(math.Pow(float64(t), 2.0)-4.0*float64(r+1))))
	bmax := math.Floor((1.0 / 2.0) * (math.Sqrt(math.Pow(float64(t), 2.0)-4.0*float64(r+1)) + float64(t)))

	return int(bmax) - int(bmin) + 1
}

func main() {
	file := flag.String("file", "data.txt", "input for races.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	// Converting times and record dists to integers
	r := regexp.MustCompile(`(\d+)`)

	timeStrs, recordStrs := r.FindAllString(lines[0], -1), r.FindAllString(lines[1], -1)

	times, records := make([]int, len(timeStrs)), make([]int, len(timeStrs))
	for i := range times {
		times[i], _ = strconv.Atoi(timeStrs[i])
		records[i], _ = strconv.Atoi(recordStrs[i])
	}

	prod := 1
	for i := range times {
		prod *= sizeOfMOE(times[i], records[i])
	}

	fmt.Println(prod)
}
