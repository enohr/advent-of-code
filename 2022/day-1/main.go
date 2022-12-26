package main

import (
	"fmt"
	"os"
	"sort"
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
	elves := parse(input)
	max := elves[0]

	for _, val := range elves {
		if val > max {
			max = val
		}
	}
	return max
}

func part2(input string) int {
	elves := parse(input)
	sum := 0
	sort.Ints(elves)

	for i := 0; i < 3; i++ {
		sum += elves[len(elves)-1-i]
	}
	return sum

}

func parse(input string) []int {
	splitted := strings.Split(input, "\n\n")
	ret := []int{}
	for _, group := range splitted {
		sum := 0
		for _, line := range strings.Split(group, "\n") {
			num, _ := strconv.Atoi(line)
			sum += num
		}
		ret = append(ret, sum)
	}

	return ret
}
