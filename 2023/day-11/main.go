package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, _ := os.ReadFile("input1.txt")
	input1 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part1:", part1(input1))

	file, _ = os.ReadFile("input2.txt")
	input2 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
	matrix, points := generateMatrix(input)
	var emptyLines, emptyCols []int

	for i, line := range matrix {
		if !slices.Contains(line, '#') {
			emptyLines = append(emptyLines, i)
		}
	}
	for j := range matrix {
		var col []rune
		for i := range matrix {
			col = append(col, matrix[i][j])
		}
		if !slices.Contains(col, '#') {
			emptyCols = append(emptyCols, j)
		}
	}

	sum := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			linesExploded := CountExplosion(p1.X, p2.X, emptyLines)
			colsExploded := CountExplosion(p1.Y, p2.Y, emptyCols)

			line := int(math.Abs(float64(p1.X)-float64(p2.X))) - linesExploded
			col := int(math.Abs(float64(p1.Y)-float64(p2.Y))) - colsExploded
			distance := col + line + (2 * linesExploded) + (2 * colsExploded)

			sum += distance
		}
	}
	return sum
}

func CountExplosion(x, y int, empty []int) int {
	sum := 0
	start := int(math.Min(float64(x), float64(y)))
	end := int(math.Max(float64(x), float64(y)))

	for i := start; i < end; i++ {
		if slices.Contains(empty, i) {
			sum++
		}
	}
	return sum
}

func part2(input string) int {
	matrix, points := generateMatrix(input)
	var emptyLines, emptyCols []int

	for i, line := range matrix {
		if !slices.Contains(line, '#') {
			emptyLines = append(emptyLines, i)
		}
	}
	for j := range matrix {
		var col []rune
		for i := range matrix {
			col = append(col, matrix[i][j])
		}
		if !slices.Contains(col, '#') {
			emptyCols = append(emptyCols, j)
		}
	}

	sum := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			linesExploded := CountExplosion(p1.X, p2.X, emptyLines)
			colsExploded := CountExplosion(p1.Y, p2.Y, emptyCols)

			line := int(math.Abs(float64(p1.X)-float64(p2.X))) - linesExploded
			col := int(math.Abs(float64(p1.Y)-float64(p2.Y))) - colsExploded

			distance := col + line + (1000000 * linesExploded) + (1000000 * colsExploded)

			sum += distance
		}
	}
	return sum
}

func generateMatrix(input string) ([][]rune, []Point) {
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	var points []Point

	for x, line := range lines {
		matrix[x] = make([]rune, len(line))
		for y, column := range line {
			matrix[x][y] = column
			if column == '#' {
				points = append(points, Point{X: x, Y: y})
			}
		}
	}
	return matrix, points
}
