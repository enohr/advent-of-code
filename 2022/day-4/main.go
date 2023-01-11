package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {

		sections := strings.Split(line, ",")
		first := sections[0]
		second := sections[1]

		firstSection := strings.Split(first, "-")
		secondSection := strings.Split(second, "-")

		if checkSections(firstSection, secondSection) {
			sum++
		}
	}
	return sum
}

func checkSections(first, second []string) bool {
	firstInitial, _ := strconv.Atoi(first[0])
	firstEnd, _ := strconv.Atoi(first[1])
	secondInitial, _ := strconv.Atoi(second[0])
	secondEnd, _ := strconv.Atoi(second[1])

	return firstInitial <= secondInitial && firstEnd >= secondEnd || firstInitial >= secondInitial && firstEnd <= secondEnd
}

func part2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sections := strings.Split(line, ",")
		first := sections[0]
		second := sections[1]
		firstSection := strings.Split(first, "-")
		secondSection := strings.Split(second, "-")
		if checkOverlaps(firstSection, secondSection) {
			sum++
		}

	}
	return sum
}

func checkOverlaps(first, second []string) bool {
	firstInitial, _ := strconv.Atoi(first[0])
	firstEnd, _ := strconv.Atoi(first[1])
	secondInitial, _ := strconv.Atoi(second[0])
	secondEnd, _ := strconv.Atoi(second[1])
	return firstInitial <= secondEnd && firstEnd >= secondInitial || secondInitial <= firstEnd && secondEnd >= firstInitial
}
