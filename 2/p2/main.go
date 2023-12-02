package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func getMinProd(games string) int {
  maxs := make(map[string]int)

  for _, obs := range strings.Split(games, "; ") {
    for _, cube := range strings.Split(obs, ", ") {
      res := strings.Split(cube, " ")
      n,_ := strconv.Atoi(res[0])

      if n > maxs[res[1]] {
        maxs[res[1]] = n
      }
    }
  }

  return maxs["red"] * maxs["green"] * maxs["blue"]
}

func main() {
  file := flag.String("file", "data.txt", "the file to parse.")
  flag.Parse()

  lines, _ := utils.ReadLines(*file)

  sumProd := 0
  for _, gamestr := range lines {
    splitGame := strings.Split(gamestr, ": ")
    sumProd += getMinProd(splitGame[1])
  }

  fmt.Println(sumProd)
}
