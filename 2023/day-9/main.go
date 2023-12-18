package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	input1 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part1:", part1(input1))

	file, _ = os.ReadFile("input2.txt")
	input2 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
	parts := strings.Split(input, "\n")

	sum := 0
	for _, p := range parts {
		splitted := strings.Split(p, " ")
		values := StringArrayToInt(splitted)
		sum += calculate(values)
	}

	return sum
}

func calculate(values []int) int {
	if slices.Max(values) == 0 && slices.Min(values) == 0 {
		return 0
	}

	var new []int
	for i := len(values) - 1; i > 0; i-- {
		v := values[i] - values[i-1]
		new = append([]int{v}, new...)
	}
	return calculate(new) + values[len(values)-1]
}

func StringArrayToInt(input []string) []int {
	var ret []int
	for _, s := range input {
		value, _ := strconv.Atoi(s)
		ret = append(ret, value)
	}
	return ret
}

func part2(input string) int {
	return 1
}
