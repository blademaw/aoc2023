package utils

import (
	"os"
	"strings"
)

func ReadLines(file string) ([]string, error) {
  dat, err := os.ReadFile(file)
  if err != nil {
    return nil, err
  }

  return strings.Split(strings.TrimSpace(string(dat)), "\n"), nil
}
