package main

import (
	"log"
	"os"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input string) int { return 1 }

func part2(input string) int { return 1 }
