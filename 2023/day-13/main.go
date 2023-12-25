package main

import (
	"fmt"
	"os"
	"slices"
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
	grids := strings.Split(input, "\n\n")

	sum := 0
	for _, grid := range grids {
		rows := strings.Split(grid, "\n")
		r := findRowMirror(rows) * 100
		c := findColMirror(rows)
		sum += r + c
	}
	return sum
}

func findRowMirror(rows []string) int {
	for r := range rows {
		first := rows[:r]
		second := rows[r:]
		first = reverse(first)

		if len(first) > len(second) {
			first = first[:len(second)]
		} else if len(second) > len(first) {
			second = second[:len(first)]
		}

		if slices.Equal(first, second) && r != 0 {
			return r
		}
	}
	return 0
}

func findColMirror(rows []string) int {
	cols := transposeGrid(rows)
	return findRowMirror(cols)
}

func part2(input string) int {
	grids := strings.Split(input, "\n\n")

	sum := 0
	for _, grid := range grids {
		rows := strings.Split(grid, "\n")
		r := findRowMirrorPt2(rows) * 100
		c := findColMirrorPt2(rows)
		sum += r + c
	}
	return sum
}

func findRowMirrorPt2(rows []string) int {
	for r := range rows {
		first := rows[:r]
		second := rows[r:]
		first = reverse(first)

		if len(first) > len(second) {
			first = first[:len(second)]
		} else if len(second) > len(first) {
			second = second[:len(first)]
		}

		if countDiffs(first, second) == 1 {
			return r
		}
	}
	return 0
}

func findColMirrorPt2(rows []string) int {
	cols := transposeGrid(rows)
	return findRowMirrorPt2(cols)
}

func countDiffs(first, second []string) int {
	sum := 0

	for r, row := range first {
		for c, char := range row {
			if char != rune(second[r][c]) {
				sum++
			}
		}
	}
	return sum
}

func reverse(input []string) []string {
	out := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		out[i] = input[len(input)-i-1]
	}
	return out
}

func transposeGrid(rows []string) []string {
	arr := make([]string, len(rows[0]))

	for _, row := range rows {
		for c := range row {
			arr[c] = arr[c] + string(row[c])
		}
	}
	return arr

}
