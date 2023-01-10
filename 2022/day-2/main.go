package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
	Lose     = 0
	Draw     = 3
	Win      = 6
)

var plays = map[string]int{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

var choices = map[int]int{
	Draw + Rock:     Rock,
	Draw + Scissors: Scissors,
	Draw + Paper:    Paper,
	Win + Rock:      Paper,
	Win + Scissors:  Rock,
	Win + Paper:     Scissors,
	Lose + Rock:     Scissors,
	Lose + Scissors: Paper,
	Lose + Paper:    Rock,
}

var results1 = map[string]int{
	"AX": Rock + Draw,
	"AY": Paper + Win,
	"AZ": Scissors + Lose,
	"BX": Rock + Lose,
	"BY": Paper + Draw,
	"BZ": Scissors + Win,
	"CX": Rock + Win,
	"CY": Paper + Lose,
	"CZ": Scissors + Draw,
}

var results2 = map[string]int{
	"AY": Draw + choices[Draw+plays["A"]],
	"AX": Lose + choices[Lose+plays["A"]],
	"AZ": Win + choices[Win+plays["A"]],
	"BY": Draw + choices[Draw+plays["B"]],
	"BX": Lose + choices[Lose+plays["B"]],
	"BZ": Win + choices[Win+plays["B"]],
	"CY": Draw + choices[Draw+plays["C"]],
	"CX": Lose + choices[Lose+plays["C"]],
	"CZ": Win + choices[Win+plays["C"]],
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		choices := strings.ReplaceAll(line, " ", "")
		sum = sum + results1[choices]
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		choices := strings.ReplaceAll(line, " ", "")
		sum = sum + results2[choices]
	}
	return sum
}
