package main

import (
	"fmt"
	"strings"
)

var content = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

func main() {
	fmt.Println(part1(content))
	fmt.Println(part2(content))
}

func part1(input string) int {
	var text string
	for i := 0; i < len(input)-4; i++ {
		text = ""
		for j := 0; j < 4; j++ {
			char := input[i+j]
			if strings.Contains(text, string(char)) {
				break
			}
			text = text + string(char)
		}
		if len(text) == 4 {
			return i + 4
		}
	}
	return 0
}

func part2(input string) int {
	var text string
	for i := 0; i < len(input)-14; i++ {
		text = ""
		for j := 0; j < 14; j++ {
			char := input[i+j]
			if strings.Contains(text, string(char)) {
				break
			}
			text = text + string(char)
		}
		if len(text) == 14 {
			return i + 14
		}
	}
	return 0
}
