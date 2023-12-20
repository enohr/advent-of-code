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

type Pos struct {
	X int
	Y int
}

func part1(input string) int {
	matrix, starting := generateMatrix(input)
	moves := 0

	goal := starting
	previous := starting

	current := FindFirstMove(starting, matrix)
	for current != goal {
		next := GetNextMove(current, previous, matrix)
		previous = current
		current = next
		moves++
	}
	return (moves / 2) + 1
}

func part2(input string) int {
	matrix, starting := generateMatrix(input)
	var path []Pos

	goal := starting
	previous := starting

	path = append(path, starting)
	current := FindFirstMove(starting, matrix)
	for current != goal {
		path = append(path, current)
		next := GetNextMove(current, previous, matrix)
		previous = current
		current = next
	}

	matrix = clearMatrix(path, matrix)

	tiles := 0
	for x, line := range matrix {
		for y := range line {
			if matrix[x][y] != '.' {
				continue
			}
			if sum := rayCasting(x, y, matrix); sum%2 != 0 {
				tiles++
			}
		}
	}
	return tiles
}

// If we cross the loop even number of times, it's a tile. It's like a Ray Casting Algorithm
// I got this idea from:
// https://github.com/hyper-neutrino/advent-of-code/blob/main/2023/day10p2.py
func rayCasting(x, y int, matrix [][]rune) int {
	height := len(matrix)
	width := len(matrix[0])
	sum := 0
	// Going diagonally it's easiar because only has two possible corners
	for i := 1; (x+i) < height && (y+i) < width; i++ {
		dX := x + i
		dY := y + i
		value := matrix[dX][dY]
		// Ignore corners
		if value == 'L' || value == '7' || value == '.' {
			continue
		}
		sum++
	}
	return sum
}

func ReplaceStarting(path []Pos) rune {
	possibles := [6]rune{'-', '|', 'J', 'L', '7', 'F'}
	goal := [2]Pos{path[1], path[len(path)-1]}
	starting := path[0]

	for _, possible := range possibles {
		next := GetPossibleMoves(starting, possible)
		// Check the two possible combinations
		if (next[0] == goal[0] && next[1] == goal[1]) || (next[1] == goal[0] && next[0] == goal[1]) {
			return possible
		}
	}
	return 'S'
}

func clearMatrix(loopPath []Pos, matrix [][]rune) [][]rune {
	starting := loopPath[0]
	for x, line := range matrix {
		for y := range line {
			if !slices.Contains(loopPath, Pos{X: x, Y: y}) {
				matrix[x][y] = '.'
			}
		}
	}
	replacedStarting := ReplaceStarting(loopPath)
	matrix[starting.X][starting.Y] = replacedStarting

	return matrix
}

func GetNextMove(current, previous Pos, matrix [][]rune) Pos {
	currentValue := matrix[current.X][current.Y]
	possible := GetPossibleMoves(current, currentValue)

	// We cant go back, so get the possible move
	// different than the previous one
	if possible[0] == previous {
		return possible[1]
	} else {
		return possible[0]
	}
}

func GetPossibleMoves(current Pos, currentValue rune) []Pos {
	var possible []Pos

	switch currentValue {
	case '|': // up or down
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X + 1, Y: current.Y}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case '-': // left or right
		pos1 := Pos{X: current.X, Y: current.Y - 1}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'J': // up or left
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y - 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'L': // up or right
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case '7': // down or left
		pos1 := Pos{X: current.X + 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y - 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'F': // down or right
		pos1 := Pos{X: current.X + 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	}
	return possible
}

func FindFirstMove(p Pos, board [][]rune) Pos {
	positions := []Pos{
		{X: -1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: 1, Y: 0},
	}
	height := len(board)
	width := len(board[0])

	for _, position := range positions {
		x := position.X + p.X
		y := position.Y + p.Y

		if x >= height || y >= width || x < 0 || y < 0 {
			continue
		}
		if board[x][y] == '.' {
			continue
		}
		return Pos{X: x, Y: y}
	}
	return Pos{X: 0, Y: 0}
}

func generateMatrix(input string) ([][]rune, Pos) {
	starting := Pos{X: 0, Y: 0}
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))

	for x, line := range lines {
		matrix[x] = make([]rune, len(line))
		for y, column := range line {
			if column == 'S' {
				starting = Pos{X: x, Y: y}
			}
			matrix[x][y] = column
		}
	}
	return matrix, starting
}
