package main

import (
	"fmt"
	"os"
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
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		str, num := fields[0], fields[1]
		nums := StringArrayToInt(strings.Split(num, ","))
		res := walk(str, nums)
		sum += res
	}
	return sum
}

func walk(str string, nums []int) int {
	if len(str) == 0 {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		if strings.Contains(str, "#") {
			return 0
		}

		return 1
	}

	if nums[0] > len(str) {
		return 0
	}

	total := 0
	if str[0] == '.' || str[0] == '?' {
		total += walk(str[1:], nums)
	}

	if str[0] == '?' || str[0] == '#' {
		s := str[:nums[0]]
		if !strings.Contains(s, ".") && len(str) >= nums[0] && (len(str) == nums[0] || str[nums[0]] != '#') {
			if len(str) == nums[0] {
				total += walk(str[nums[0]:], nums[1:])
			} else {
				total += walk(str[nums[0]+1:], nums[1:])
			}
		} else if strings.Count(str, "#") == nums[0] && len(str) == nums[0] {
			total += walk(str[nums[0]:], nums[1:])
		}
	}

	return total
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		str := strings.Repeat(fields[0]+"?", 5)
		str = str[:len(str)-1]
		num := strings.Repeat(fields[1]+",", 5)
		nums := StringArrayToInt(strings.Split(num[:len(num)-1], ","))

		res := walk(str, nums)
		sum += res
	}
	return sum
}

func StringArrayToInt(input []string) []int {
	var ret []int
	for _, s := range input {
		value, _ := strconv.Atoi(s)
		ret = append(ret, value)
	}
	return ret
}
