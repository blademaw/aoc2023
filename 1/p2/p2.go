package main

import (
	"aoc2023/utils"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var NMAP = map[string]int{
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9,
}

func parseSliceToNum(line string) (string, bool) {
  if unicode.IsDigit(rune(line[0])) {
    return string(line[0]), true
  }

  for num, digit := range NMAP {
    if strings.HasPrefix(line, num) {
      return strconv.Itoa(digit), true
    }
  }

  return "", false
}

func lineNumber(line string) int {
  var first, last string
  var ok bool

  for i := 0; i < len(line); i++ {
    first, ok = parseSliceToNum(line[i:])
    if !ok {
      continue
    } else {
      break
    }
  }

  for j := len(line)-1; j >= 0; j-- {
    last, ok = parseSliceToNum(line[j:])
    if !ok {
      continue
    } else {
      break
    }
  }

  res, _ := strconv.Atoi(first + last)
  return res
}

func main() {
  file := flag.String("file", "data.txt", "the file to parse.")
  flag.Parse()

  lines, _ := utils.ReadLines(*file)

  total := 0
  for _, line := range lines {
    total += lineNumber(line)
  }

  fmt.Println(total)
}

