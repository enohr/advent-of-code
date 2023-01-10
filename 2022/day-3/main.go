package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	FIRST_LOWER       = 97
	FIRST_UPPER       = 65
	FIRST_UPPER_PRIOR = 27
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
		line = string(line)
		first := line[:len(line)/2]
		second := line[len(line)/2:]
		for _, char := range first {
			if strings.Contains(second, string(char)) {
				sum = sum + calculatePriority(char)
				break
			}
		}
	}
	return sum
}

func part2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

		for _, char := range first {
			if strings.Contains(second, string(char)) && strings.Contains(third, string(char)) {
				sum = sum + calculatePriority(char)
				break
			}
		}
	}

	return sum

}

func calculatePriority(char rune) (value int) {
	if unicode.IsLower(char) {
		value = int(char) - FIRST_LOWER + 1
		return
	}
	value = int(char) - FIRST_UPPER + FIRST_UPPER_PRIOR
	return
}
