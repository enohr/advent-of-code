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
		cache := make(map[string]int)
		fields := strings.Fields(line)
		str, num := fields[0], fields[1]
		nums := StringArrayToInt(strings.Split(num, ","))
		sum += walk(str, nums, cache)
	}
	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		cache := make(map[string]int)
		fields := strings.Fields(line)
		str := strings.Repeat(fields[0]+"?", 5)
		str = str[:len(str)-1]
		num := strings.Repeat(fields[1]+",", 5)
		nums := StringArrayToInt(strings.Split(num[:len(num)-1], ","))

		sum += walk(str, nums, cache)
	}
	return sum
}

func walk(str string, nums []int, cache map[string]int) int {
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

	key := generateKey(str, nums)
	if value, ok := cache[key]; ok {
		return value
	}

	if nums[0] > len(str) {
		return 0
	}

	total := 0
	if str[0] == '.' || str[0] == '?' {
		total += walk(str[1:], nums, cache)
	}

	// This need a refactor asap
	if str[0] == '?' || str[0] == '#' {
		s := str[:nums[0]]
		if !strings.Contains(s, ".") && len(str) >= nums[0] && (len(str) == nums[0] || str[nums[0]] != '#') {
			if len(str) == nums[0] {
				total += walk(str[nums[0]:], nums[1:], cache)
			} else {
				total += walk(str[nums[0]+1:], nums[1:], cache)
			}
		} else if strings.Count(str, "#") == nums[0] && len(str) == nums[0] {
			total += walk(str[nums[0]:], nums[1:], cache)
		}
	}

	cache[key] = total
	return total
}

func generateKey(str string, nums []int) string {
	key := str + "_"

	for _, n := range nums {
		s := strconv.Itoa(n)
		key += s
	}

	return key
}

func StringArrayToInt(input []string) []int {
	var ret []int
	for _, s := range input {
		value, _ := strconv.Atoi(s)
		ret = append(ret, value)
	}
	return ret
}
