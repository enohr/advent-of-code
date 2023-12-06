package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	input1 := string(file)
  fmt.Println("Part1:", part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := string(file)
  fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
  return 1
}

func part2(input string) int {
  return 1
}
