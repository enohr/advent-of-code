package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int { return 1 }

func part2(input string) int { return 1 }
